package errors

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

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
	Error := &Error{
		Op:  Op,
		Msg: Msg,
		Err: Err,
	}

	log.Error().Msg(Error.Error())

	return Error
}
