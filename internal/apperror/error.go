package apperror

import (
	"errors"
	"fmt"
)

// swagger:parameters apperror.AppError
type AppError struct {
	Msg string `json:"message"`
	Err error  `json:"-"`
}

var ErrIncorrectNumber = NewError("the number should start with 7 and with a length of 11", errors.New("incorrect_number"))

var ErrIncorrectTime = NewError("the requested time is less than the present", errors.New("incorrect_time"))

var ErrClientAttribute = NewError("the client attribute is empty", errors.New("not_found_id"))

var ErrEmptyNotification = NewError("so far there are no mailings in the database", errors.New("not_found_notification"))

func (a *AppError) Error() string {
	return fmt.Sprintf("%s: %v", a.Msg, a.Err)
}

func NewError(msg string, err error) *AppError {
	return &AppError{
		Msg: msg,
		Err: err,
	}
}
