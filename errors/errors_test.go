package errors_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/atsumarukun/holos-api-pkg/errors"
)

func TestErr_New(t *testing.T) {
	tests := []struct {
		name         string
		inputCode    errors.ErrorCode
		inputMessage string
		expect       string
	}{
		{name: "basic", inputCode: errors.CodeUnknown, inputMessage: "test error", expect: "UNKNOWN: test error"},
		{name: "empty message", inputCode: errors.CodeUnknown, inputMessage: "", expect: "UNKNOWN: "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errors.New(tt.inputCode, tt.inputMessage)

			if err == nil {
				t.Error("expected error, but got nil")
			}

			if err.Error() != tt.expect {
				t.Errorf("expected %s, but got %s", tt.expect, err.Error())
			}
		})
	}
}

func TestErr_Code(t *testing.T) {
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
			err := errors.New(tt.inputCode, "test error")

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

func TestErr_Message(t *testing.T) {
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
			err := errors.New(errors.CodeUnknown, tt.inputMessage)

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

func TestErr_Format(t *testing.T) {
	tests := []struct {
		name        string
		inputFormat string
		checkFunc   func(t *testing.T, result string)
	}{
		{
			name:        "%s",
			inputFormat: "%s",
			checkFunc: func(t *testing.T, result string) {
				expect := "UNKNOWN: test error"
				if result != expect {
					t.Errorf("expected %s, but got %s", expect, result)
				}
			},
		},
		{
			name:        "%v",
			inputFormat: "%v",
			checkFunc: func(t *testing.T, result string) {
				expect := "UNKNOWN: test error"
				if result != expect {
					t.Errorf("expected %s, but got %s", expect, result)
				}
			},
		},
		{
			name:        "%+v",
			inputFormat: "%+v",
			checkFunc: func(t *testing.T, result string) {
				if !strings.HasPrefix(result, "UNKNOWN: test error") {
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
				if result != "%!d(*errors.err=UNKNOWN: test error)" {
					t.Errorf("no format error message")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errors.New(errors.CodeUnknown, "test error")
			tt.checkFunc(t, fmt.Sprintf(tt.inputFormat, err))
		})
	}
}
