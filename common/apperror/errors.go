package apperror

import (
	"fmt"
	"strings"
)




type ErrorCode int



const (
	UnauthorizedError ErrorCode = 4001
    ForbiddenError ErrorCode = 4003
    BadRequestError ErrorCode = 4003
    InternalError ErrorCode = 5000
    NotFoundError ErrorCode = 4004
)

func Unauthorized(message string) error {
    message = strings.ToLower(message)
    return fmt.Errorf("%d: %s", UnauthorizedError, message)
}
func Forbidden(message string) error {
    message = strings.ToLower(message)
    return fmt.Errorf("%d: %s", ForbiddenError, message)
}

func NotFound(message string) error {
    message = strings.ToLower(message)
    return fmt.Errorf("%d: %s", NotFoundError, message)
}

func BadRequest(message string) error {
    message = strings.ToLower(message)
    return fmt.Errorf("%d: %s", BadRequestError, message)
}

func Internal(message string) error {
    message = strings.ToLower(message)
    return fmt.Errorf("%d: %s", InternalError, message)
}