package errors_test

import (
	"testing"

	"github.com/atsumarukun/holos-api-pkg/errors"
)

func TestErrorCode_String(t *testing.T) {
	tests := []struct {
		name   string
		input  errors.ErrorCode
		expect string
	}{
		{name: "BadRequest", input: errors.CodeBadRequest, expect: "BAD_REQUEST"},
		{name: "Unauthenticated", input: errors.CodeUnauthenticated, expect: "UNAUTHENTICATED"},
		{name: "Unauthorized", input: errors.CodeUnauthorized, expect: "UNAUTHORIZED"},
		{name: "NotFound", input: errors.CodeNotFound, expect: "NOT_FOUND"},
		{name: "Duplicate", input: errors.CodeDuplicate, expect: "DUPLICATE"},
		{name: "ConstraintViolation", input: errors.CodeConstraintViolation, expect: "CONSTRAINT_VIOLATION"},
		{name: "InvalidInput", input: errors.CodeInvalidInput, expect: "INVALID_INPUT"},
		{name: "InternalServerError", input: errors.CodeInternalServerError, expect: "INTERNAL_SERVER_ERROR"},
		{name: "Unknown", input: errors.CodeUnknown, expect: "UNKNOWN"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.String()
			if result != tt.expect {
				t.Errorf("expect %s, but got %s", tt.expect, result)
			}
		})
	}
}
