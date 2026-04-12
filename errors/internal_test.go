package errors

import (
	stderr "errors"
	"reflect"
	"testing"
)

func TestWrap_Wrap_ProgramCounters_Inherited(t *testing.T) {
	errWithStack := New(CodeUnknown, "cause error")
	errWithoutStack := stderr.New("cause error")

	tests := []struct {
		name     string
		inputErr error
	}{
		{name: "with stack", inputErr: errWithStack},
		{name: "without stack", inputErr: errWithoutStack},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.inputErr, CodeUnknown, "test error")

			base, ok1 := err.(interface{ programCounters() programCounters })
			if !ok1 {
				t.Error("programCounters not implemented")
			}
			cause, ok2 := stderr.Unwrap(err).(interface{ programCounters() programCounters })

			if ok2 {
				if !reflect.DeepEqual(base.programCounters(), cause.programCounters()) {
					t.Error("programCounters should be inherited but not equal")
				}
			} else {
				if base.programCounters() == nil {
					t.Error("programCounters has not been initialized")
				}
			}
		})
	}
}
