package errors

import (
	"testing"
)

func TestStatusCode_Unknown(t *testing.T) {
	if got := StatusCode(nil); got != 500 {
		t.Errorf("StatusCode(nil) = %d, want 500", got)
	}
	if got := StatusCode(errPlain("x")); got != 500 {
		t.Errorf("StatusCode(plain) = %d, want 500", got)
	}
}

func TestStatusCode_Domain(t *testing.T) {
	if got := StatusCode(ErrNotFound); got != 404 {
		t.Errorf("StatusCode(ErrNotFound) = %d, want 404", got)
	}
	if got := StatusCode(ErrInvalidInput); got != 400 {
		t.Errorf("StatusCode(ErrInvalidInput) = %d, want 400", got)
	}
	if got := StatusCode(ErrForbidden); got != 403 {
		t.Errorf("StatusCode(ErrForbidden) = %d, want 403", got)
	}
	if got := StatusCode(ErrUpstream); got != 502 {
		t.Errorf("StatusCode(ErrUpstream) = %d, want 502", got)
	}
}

func TestMessage_Domain(t *testing.T) {
	if got := Message(ErrNotFound); got != "not found" {
		t.Errorf("Message(ErrNotFound) = %q, want \"not found\"", got)
	}
	if got := Message(NotFound("custom")); got != "custom" {
		t.Errorf("Message(NotFound(\"custom\")) = %q, want \"custom\"", got)
	}
}

func TestMessage_Plain(t *testing.T) {
	e := errPlain("something broke")
	if got := Message(e); got != "something broke" {
		t.Errorf("Message(plain) = %q, want \"something broke\"", got)
	}
}

type errPlain string

func (e errPlain) Error() string { return string(e) }
