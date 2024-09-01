// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package oapi

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// LoginChallengePasskeyRequest Initialize Assertion Response
type LoginChallengePasskeyRequest = map[string]interface{}

// LoginChallengePasskeyResponse Initialize Assertion Response
type LoginChallengePasskeyResponse = map[string]interface{}

// LoginPasskeyRequest Finalize Assertion Request
type LoginPasskeyRequest = map[string]interface{}

// LoginPasskeyResponse Finalize Assertion Response
type LoginPasskeyResponse = map[string]interface{}

// RegisterChallengePasskeyResponse Initialize Attestation Response
type RegisterChallengePasskeyResponse = map[string]interface{}

// RegisterPasskeyRequest Finalize Attestation Request
type RegisterPasskeyRequest = map[string]interface{}

// LoginPasskeyParams defines parameters for LoginPasskey.
type LoginPasskeyParams struct {
	// Assertion session
	Assertion string `form:"__assertion__" json:"__assertion__"`
}

// RegisterPasskeyParams defines parameters for RegisterPasskey.
type RegisterPasskeyParams struct {
	// Attestation session
	Attestation string `form:"__attestation__" json:"__attestation__"`
}

// RegisterChallengePasskeyJSONBody defines parameters for RegisterChallengePasskey.
type RegisterChallengePasskeyJSONBody struct {
	Email *openapi_types.Email `json:"email,omitempty"`
}

// LoginPasskeyJSONRequestBody defines body for LoginPasskey for application/json ContentType.
type LoginPasskeyJSONRequestBody = LoginPasskeyRequest

// LoginChallengePasskeyJSONRequestBody defines body for LoginChallengePasskey for application/json ContentType.
type LoginChallengePasskeyJSONRequestBody = LoginChallengePasskeyRequest

// RegisterPasskeyJSONRequestBody defines body for RegisterPasskey for application/json ContentType.
type RegisterPasskeyJSONRequestBody = RegisterPasskeyRequest

// RegisterChallengePasskeyJSONRequestBody defines body for RegisterChallengePasskey for application/json ContentType.
type RegisterChallengePasskeyJSONRequestBody RegisterChallengePasskeyJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Login with a passkey
	// (POST /passkey/login)
	LoginPasskey(w http.ResponseWriter, r *http.Request, params LoginPasskeyParams)
	// Generate a login challenge
	// (POST /passkey/login-challenge)
	LoginChallengePasskey(w http.ResponseWriter, r *http.Request)
	// Register a passkey
	// (POST /passkey/register)
	RegisterPasskey(w http.ResponseWriter, r *http.Request, params RegisterPasskeyParams)
	// Register a new passkey
	// (POST /passkey/register-challenge)
	RegisterChallengePasskey(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Login with a passkey
// (POST /passkey/login)
func (_ Unimplemented) LoginPasskey(w http.ResponseWriter, r *http.Request, params LoginPasskeyParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Generate a login challenge
// (POST /passkey/login-challenge)
func (_ Unimplemented) LoginChallengePasskey(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Register a passkey
// (POST /passkey/register)
func (_ Unimplemented) RegisterPasskey(w http.ResponseWriter, r *http.Request, params RegisterPasskeyParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Register a new passkey
// (POST /passkey/register-challenge)
func (_ Unimplemented) RegisterChallengePasskey(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// LoginPasskey operation middleware
func (siw *ServerInterfaceWrapper) LoginPasskey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params LoginPasskeyParams

	var cookie *http.Cookie

	if cookie, err = r.Cookie("__assertion__"); err == nil {
		var value string
		err = runtime.BindStyledParameterWithOptions("simple", "__assertion__", cookie.Value, &value, runtime.BindStyledParameterOptions{Explode: true, Required: true})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "__assertion__", Err: err})
			return
		}
		params.Assertion = value

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "__assertion__"})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LoginPasskey(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LoginChallengePasskey operation middleware
func (siw *ServerInterfaceWrapper) LoginChallengePasskey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LoginChallengePasskey(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// RegisterPasskey operation middleware
func (siw *ServerInterfaceWrapper) RegisterPasskey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params RegisterPasskeyParams

	var cookie *http.Cookie

	if cookie, err = r.Cookie("__attestation__"); err == nil {
		var value string
		err = runtime.BindStyledParameterWithOptions("simple", "__attestation__", cookie.Value, &value, runtime.BindStyledParameterOptions{Explode: true, Required: true})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "__attestation__", Err: err})
			return
		}
		params.Attestation = value

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "__attestation__"})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RegisterPasskey(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// RegisterChallengePasskey operation middleware
func (siw *ServerInterfaceWrapper) RegisterChallengePasskey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RegisterChallengePasskey(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/passkey/login", wrapper.LoginPasskey)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/passkey/login-challenge", wrapper.LoginChallengePasskey)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/passkey/register", wrapper.RegisterPasskey)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/passkey/register-challenge", wrapper.RegisterChallengePasskey)
	})

	return r
}
