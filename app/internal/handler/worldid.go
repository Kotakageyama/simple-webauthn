package handler

import (
	"app/internal/repository"
	"app/internal/usecase"
	"context"
	"encoding/json"
	"net/http"
)

type Handler struct {
	loginUsecase    usecase.LoginUsecase
	sessionRepo     repository.SessionRepository
	userRepo        repository.UserRepository
}

func NewHandler(
	loginUsecase usecase.LoginUsecase,
	sessionRepo repository.SessionRepository,
	userRepo repository.UserRepository,
) *Handler {
	return &Handler{
		loginUsecase: loginUsecase,
		sessionRepo:  sessionRepo,
		userRepo:     userRepo,
	}
}

type WorldIDVerifyRequest struct {
	Proof  json.RawMessage `json:"proof"`
	Action string          `json:"action"`
	Signal string          `json:"signal"`
}

type WorldIDVerifyResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	UserID  []byte `json:"user_id,omitempty"`
}

func (h *Handler) HandleWorldIDVerify(w http.ResponseWriter, r *http.Request) {
	var req WorldIDVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Extract nullifier from proof (this will be implemented in verifyWorldIDProof)
	nullifier := "temp-nullifier" // Temporary placeholder

	// Use login usecase for World ID authentication
	user, sessionID, err := h.loginUsecase.LoginWorldID(nullifier, string(req.Proof))
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     sessionID.ToCookieKey(),
		Value:    sessionID.String(),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	// Return success response with user information
	resp := WorldIDVerifyResponse{
		Success: true,
		Message: "World ID verification successful",
		UserID:  user.ID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) verifyWorldIDProof(ctx context.Context, proof json.RawMessage, action, signal string) (bool, error) {
	// TODO: Implement World ID proof verification using the SDK
	// This will be implemented in the next step after setting up the verification package
	return true, nil
}
