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

func NewError(Op Op, Msg Msg, Err Err) error {
	return &Error{
		Op:  Op,
		Msg: Msg,
		Err: Err,
	}
}
