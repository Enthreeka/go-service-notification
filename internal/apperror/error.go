package apperror

import (
	"errors"
	"fmt"
)

type appError struct {
	Msg string `json:"message"`
	Err error  `json:"-"`
}

var ErrIncorrectNumber = NewError("the number should start with 7 and with a length of 11", errors.New("incorrect_number"))

func (a *appError) Error() string {
	return fmt.Sprintf("%s: %v", a.Msg, a.Err)
}

func NewError(msg string, err error) *appError {
	return &appError{
		Msg: msg,
		Err: err,
	}
}
