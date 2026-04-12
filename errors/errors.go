package errors

import "fmt"

type err struct {
	code    ErrorCode
	message string
	pcs     programCounters
}

func New(code ErrorCode, message string) error {
	return &err{code, message, callers()}
}

func (e *err) Error() string {
	return fmt.Sprintf("%s: %s", e.code, e.message)
}

func (e *err) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v', 's':
		switch {
		case f.Flag('+'):
			fmt.Fprintf(f, "%s\n\n%v", e.Error(), e.pcs)
		default:
			fmt.Fprint(f, e.Error())
		}
	default:
		fmt.Fprintf(f, "%%!%c(%T=%v)", verb, e, e.Error())
	}
}

func (e *err) Code() ErrorCode {
	return e.code
}

func (e *err) Message() string {
	return e.message
}

func (e *err) programCounters() programCounters {
	return e.pcs
}
