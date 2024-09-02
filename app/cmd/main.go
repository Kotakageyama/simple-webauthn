package main

import (
	"app/internal/di"
	"app/internal/handler/oapi"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-webauthn/webauthn/webauthn"
)

var (
	Wc  *webauthn.WebAuthn
	err error
)

func main() {
	Wc, err = webauthn.New(&webauthn.Config{
		RPDisplayName: "Simple WebAuthn",                 // Display Name for your site
		RPID:          "localhost",                       // Generally the FQDN for your site
		RPOrigins:     []string{"http://localhost:3000"}, // The origin URLs allowed for WebAuthn requests
	})

	if err != nil {
		log.Fatal("failed to create WebAuthn from config:", err)
	}

	r := chi.NewMux()

	// ミドルウェアの設定
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// CORSの設定
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // 許可するオリジン
		AllowedMethods:   []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // キャッシュの有効期限（秒）
	}
	r.Use(cors.Handler(corsOptions))

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	handler := di.Wire(nil, Wc)
	h := oapi.HandlerFromMux(handler, r)

	// サーバーを起動
	log.Println("Starting server on :8080...")
	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	if err := s.ListenAndServe(); err != nil {
		log.Println("Server shutdown...")
		log.Fatal(err)
	}
}
