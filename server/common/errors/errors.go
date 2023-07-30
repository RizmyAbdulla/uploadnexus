package errors

import "fmt"

type Op string
type Msg string
type Err error

type Error struct {
	Op  Op
	Msg Msg
	Err Err
}

func (e *Error) Error() string {
	return fmt.Sprintf("Operation: %v, Message: %v, Error: %v", e.Op, e.Msg, e.Err)
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, args := range args {
		switch args.(type) {
		case Op:
			e.Op = args.(Op)
		case Msg:
			e.Msg = args.(Msg)
		case Err:
			e.Err = args.(Err)
		}
	}
	return e
}

func Ops(e *Error) []Op {
	res := []Op{e.Op}

	subErr, ok := e.Err.(*Error)
	if !ok {
		return res
	}

	res = append(res, Ops(subErr)...)

	return res
}
