package errors

import (
	"errors"
	"fmt"
)

type Err struct {
	err error
	msg string
	stk *stack
}

func Wrap(err error) error {
	return Wrapf(err, "")
}

// New is function of "errors" package
func New(text string) error {
	return &Err{
		err: nil,
		msg: text,
		stk: callers(),
	}
}

func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	return &Err{
		err: err,
		msg: fmt.Sprintf(format, args...),
		stk: callers(),
	}
}

func (e *Err) Error() string {
	if e.err == nil {
		return e.msg
	}
	if e.msg == "" {
		return e.err.Error()
	}
	return e.msg + ": " + e.err.Error()
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
