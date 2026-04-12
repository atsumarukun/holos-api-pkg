package errors

import "fmt"

type wrap struct {
	cause   error
	code    ErrorCode
	message string
	pcs     programCounters
}

func Wrap(err error, code ErrorCode, message string) error {
	// NOTE: programCountersを保持している場合は引き継ぐ.
	if e, ok := err.(interface{ programCounters() programCounters }); ok {
		return &wrap{err, code, message, e.programCounters()}
	}
	return &wrap{err, code, message, callers()}
}

func (w *wrap) Error() string {
	return fmt.Sprintf("%s: %s: %s", w.code, w.message, w.cause.Error())
}

func (w *wrap) Unwrap() error {
	return w.cause
}

func (w *wrap) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v', 's':
		switch {
		case f.Flag('+'):
			fmt.Fprintf(f, "%s\n\n%v", w.Error(), w.pcs)
		default:
			fmt.Fprint(f, w.Error())
		}
	default:
		fmt.Fprintf(f, "%%!%c(%T=%v)", verb, w, w.Error())
	}
}

func (w *wrap) Code() ErrorCode {
	return w.code
}

func (w *wrap) Message() string {
	return w.message
}

func (w *wrap) programCounters() programCounters {
	return w.pcs
}
