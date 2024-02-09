package apperror

import (
	"errors"
	"fmt"
)




type ErrorCode int



const (
	UnauthorizedError ErrorCode = 4001
    ForbiddenError ErrorCode = 4003
    BadRequestError ErrorCode = 4003
    InternalError ErrorCode = 5000
)

func Unauthorized(message string) error {
    return errors.New(fmt.Sprintf("%d: %s", UnauthorizedError, message))
}
func Forbidden(message string) error {
    return errors.New(fmt.Sprintf("%d: %s", ForbiddenError, message))
}

func BadRequest(message string) error {
    return errors.New(fmt.Sprintf("%d: %s", BadRequestError, message))
}

func Internal(message string) error {
    return errors.New(fmt.Sprintf("%d: %s", InternalError, message))
}