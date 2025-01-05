package errors

import (
	"runtime"
	"strconv"
	"strings"
)

type stack []uintptr

func callers() *stack {
	const depth = 32
	const callerStackSkip = 4

	var pcs [depth]uintptr
	n := runtime.Callers(callerStackSkip, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func (s *stack) StackTrace() StackTrace {
	f := make([]Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = Frame((*s)[i])
	}
	return f
}

// StackTrace is stack of Frames from innermost (newest) to outermost (oldest)
type StackTrace []Frame

func (st StackTrace) Format() string {
	var trace string
	for _, f := range st {
		if strings.Contains(f.name(), "runtime") {
			continue
		}
		trace += strings.Join([]string{f.name(), f.file(), strconv.Itoa(f.line()), "\n"}, " ")
	}
	return trace
}
