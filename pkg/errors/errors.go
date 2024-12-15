package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("code=%d, message=%s, error=%v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

// 预定义错误
var (
	ErrInvalidParams = &Error{
		Code:    40001,
		Message: "Invalid parameters",
	}
	
	ErrUnauthorized = &Error{
		Code:    40100,
		Message: "Unauthorized",
	}
	
	ErrForbidden = &Error{
		Code:    40300,
		Message: "Forbidden",
	}
	
	ErrNotFound = &Error{
		Code:    40400,
		Message: "Resource not found",
	}
	
	ErrInternalServer = &Error{
		Code:    50000,
		Message: "Internal server error",
	}
)

// 错误包装
func Wrap(err error, message string) *Error {
	if err == nil {
		return nil
	}
	
	if e, ok := err.(*Error); ok {
		return &Error{
			Code:    e.Code,
			Message: message,
			Err:     e,
		}
	}
	
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: message,
		Err:     err,
	}
} 