package errors

import (
	"errors"
	"fmt"
	"net/http"
)

// Sentinel domain errors for consistent HTTP mapping.
var (
	ErrInvalidInput = &DomainError{Code: http.StatusBadRequest, Msg: "invalid input"}
	ErrNotFound     = &DomainError{Code: http.StatusNotFound, Msg: "not found"}
	ErrUnauthorized = &DomainError{Code: http.StatusUnauthorized, Msg: "unauthorized"}
	ErrForbidden    = &DomainError{Code: http.StatusForbidden, Msg: "forbidden"}
	ErrUpstream     = &DomainError{Code: http.StatusBadGateway, Msg: "upstream error"}
)

// DomainError carries an HTTP status code and optional message override.
type DomainError struct {
	Code int
	Msg  string
	Err  error
}

func (e *DomainError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Msg, e.Err)
	}
	return e.Msg
}

func (e *DomainError) Unwrap() error { return e.Err }

// InvalidInput returns an error that maps to 400.
func InvalidInput(msg string) error {
	return &DomainError{Code: http.StatusBadRequest, Msg: msg}
}

// NotFound returns an error that maps to 404.
func NotFound(msg string) error {
	return &DomainError{Code: http.StatusNotFound, Msg: msg}
}

// Upstream wraps a cause as 502 upstream error.
func Upstream(cause error) error {
	if cause == nil {
		return nil
	}
	return &DomainError{Code: http.StatusBadGateway, Msg: "upstream error", Err: cause}
}

// StatusCode returns the HTTP status for err, or 500 for unknown errors.
func StatusCode(err error) int {
	var domain *DomainError
	if errors.As(err, &domain) {
		return domain.Code
	}
	return http.StatusInternalServerError
}

// Message returns a safe user-facing message for err (domain message or generic).
func Message(err error) string {
	if err == nil {
		return ""
	}
	var domain *DomainError
	if errors.As(err, &domain) {
		if domain.Msg != "" {
			return domain.Msg
		}
	}
	return err.Error()
}
