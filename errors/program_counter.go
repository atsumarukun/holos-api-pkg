package errors

import (
	"fmt"
	"runtime"
)

type programCounters []uintptr

func (pcs programCounters) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v', 's':
		for _, pc := range pcs {
			fn := runtime.FuncForPC(pc - 1)
			file, line := fn.FileLine(pc - 1)
			fmt.Fprintf(f, "%s:%d\n", file, line)
		}
	default:
		fmt.Fprintf(f, "%%!%c(%T=%v)\n", verb, pcs, pcs)
	}
}

func callers() programCounters {
	const depth = 32
	pcs := make([]uintptr, depth)
	n := runtime.Callers(3, pcs)
	return pcs[:n]
}
