package errors_test

import (
	stderr "errors"
	"fmt"
	"strings"
	"testing"

	"github.com/atsumarukun/holos-api-pkg/errors"
)

func TestWrap_Wrap(t *testing.T) {
	err := stderr.New("cause error")

	tests := []struct {
		name         string
		inputErr     error
		inputCode    errors.ErrorCode
		inputMessage string
		expectErr    string
		expectCause  error
	}{
		{name: "basic", inputErr: err, inputCode: errors.CodeUnknown, inputMessage: "test error", expectErr: "UNKNOWN: test error: cause error", expectCause: err},
		{name: "empty message", inputErr: err, inputCode: errors.CodeUnknown, inputMessage: "", expectErr: "UNKNOWN: : cause error", expectCause: err},
		{name: "nil", inputErr: nil, inputCode: errors.CodeUnknown, inputMessage: "", expectErr: "", expectCause: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errors.Wrap(tt.inputErr, tt.inputCode, tt.inputMessage)

			if err == nil {
				if tt.expectCause == nil {
					return
				}
				t.Error("expected error, but got nil")
			}

			if err.Error() != tt.expectErr {
				t.Errorf("expected %s, but got %s", tt.expectErr, err.Error())
			}

			cause := stderr.Unwrap(err)
			if !stderr.Is(cause, tt.expectCause) {
				t.Errorf("expected %v, but got %v", tt.expectCause, cause)
			}
		})
	}
}

func TestWrap_Format(t *testing.T) {
	tests := []struct {
		name        string
		inputFormat string
		checkFunc   func(t *testing.T, result string)
	}{
		{
			name:        "%s",
			inputFormat: "%s",
			checkFunc: func(t *testing.T, result string) {
				expect := "UNKNOWN: test error: cause error"
				if result != expect {
					t.Errorf("expected %s, but got %s", expect, result)
				}
			},
		},
		{
			name:        "%v",
			inputFormat: "%v",
			checkFunc: func(t *testing.T, result string) {
				expect := "UNKNOWN: test error: cause error"
				if result != expect {
					t.Errorf("expected %s, but got %s", expect, result)
				}
			},
		},
		{
			name:        "%+v",
			inputFormat: "%+v",
			checkFunc: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "UNKNOWN: test error: cause error") {
					t.Error("no message")
				}
				if !strings.Contains(result, ".go:") {
					t.Error("no stack trace")
				}
			},
		},
		{
			name:        "invalid verb",
			inputFormat: "%d",
			checkFunc: func(t *testing.T, result string) {
				if result != "%!d(*errors.wrap=UNKNOWN: test error: cause error)" {
					t.Errorf("no format error message")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errors.Wrap(stderr.New("cause error"), errors.CodeUnknown, "test error")
			tt.checkFunc(t, fmt.Sprintf(tt.inputFormat, err))
		})
	}
}

func TestWrap_Code(t *testing.T) {
	tests := []struct {
		name       string
		inputCode  errors.ErrorCode
		expectCode errors.ErrorCode
	}{
		{name: "BadRequest", inputCode: errors.CodeBadRequest, expectCode: errors.CodeBadRequest},
		{name: "Unauthenticated", inputCode: errors.CodeUnauthenticated, expectCode: errors.CodeUnauthenticated},
		{name: "Unauthorized", inputCode: errors.CodeUnauthorized, expectCode: errors.CodeUnauthorized},
		{name: "NotFound", inputCode: errors.CodeNotFound, expectCode: errors.CodeNotFound},
		{name: "Duplicate", inputCode: errors.CodeDuplicate, expectCode: errors.CodeDuplicate},
		{name: "ConstraintViolation", inputCode: errors.CodeConstraintViolation, expectCode: errors.CodeConstraintViolation},
		{name: "InvalidInput", inputCode: errors.CodeInvalidInput, expectCode: errors.CodeInvalidInput},
		{name: "InternalServerError", inputCode: errors.CodeInternalServerError, expectCode: errors.CodeInternalServerError},
		{name: "Unknown", inputCode: errors.CodeUnknown, expectCode: errors.CodeUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errors.Wrap(stderr.New("cause error"), tt.inputCode, "test error")

			e, ok := err.(interface {
				Code() errors.ErrorCode
			})
			if !ok {
				t.Error("does not implement Code()")
			}

			if e.Code() != tt.expectCode {
				t.Errorf("expected %s, but got %s", tt.expectCode, e.Code())
			}
		})
	}
}

func TestWrap_Message(t *testing.T) {
	tests := []struct {
		name          string
		inputMessage  string
		expectMessage string
	}{
		{name: "basic", inputMessage: "test error", expectMessage: "test error"},
		{name: "empty", inputMessage: "", expectMessage: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errors.Wrap(stderr.New("cause error"), errors.CodeUnknown, tt.inputMessage)

			e, ok := err.(interface {
				Message() string
			})
			if !ok {
				t.Error("does not implement Message()")
			}

			if e.Message() != tt.expectMessage {
				t.Errorf("expected %s, but got %s", tt.expectMessage, e.Message())
			}
		})
	}
}
