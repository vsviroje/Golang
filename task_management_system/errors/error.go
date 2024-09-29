package errors

import (
	"fmt"
	"strings"
)

//IError interface
type IError interface {
	ErrorID() int
	ErrorCode() string
	Error() string
	ErrorMsg() string
	ErrorDescription() string
}

//BaseError struct
type BaseError struct {
	id          int
	code        string
	message     string
	description string
	msg         string //to be used in Error()
}

//ErrorID function
func (e BaseError) ErrorID() int {
	return e.id
}

//ErrorCode function
func (e BaseError) ErrorCode() string {
	return e.code
}

//Error function
func (e BaseError) Error() string {
	if len(e.msg) == 0 {
		return e.message
	}
	return e.msg
}

//ErrorMsg for response writer
func (e BaseError) ErrorMsg() string {
	return e.message
}

//ErrorDescription for response writer
func (e BaseError) ErrorDescription() string {
	return e.description
}

//New error function
func New(errorID int, errorCode string, errorMsg string, errorDescription ...string) *BaseError {
	desc := errorMsg
	msg := errorMsg
	if len(errorDescription) > 0 {
		desc = strings.Join(errorDescription, " ")
		msg = fmt.Sprintf("%s %s", errorMsg, desc)
	}

	return newBaseError(errorID, errorCode, errorMsg, desc, msg)
}

func newBaseError(errorID int, errorCode string, errorMsg string, errorDescription string, msg string) *BaseError {
	return &BaseError{
		id:          errorID,
		code:        errorCode,
		message:     errorMsg,
		description: errorDescription,
		msg:         msg,
	}
}
