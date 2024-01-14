package errors

import (
	"encoding/json"
	"errors"
	"github.com/GabrielHCataldo/go-helper/helper"
	"runtime"
)

type Error struct {
	// File name file from caller new error
	File string `json:"file,omitempty"`
	// Line from caller new error
	Line int `json:"line,omitempty"`
	// Message error
	Message string `json:"message,omitempty"`
	// Endpoint from error
	Endpoint string `json:"endpoint,omitempty"`
}

// New error detail with message values separate per space
func New(message ...any) *Error {
	_, file, line, _ := runtime.Caller(1)
	return &Error{
		Message: printMessage(message...),
		Line:    line,
		File:    file,
	}
}

// NewEndpoint error detail with endpoint and message
func NewEndpoint(endpoint string, message ...any) *Error {
	_, file, line, _ := runtime.Caller(1)
	return &Error{
		Message:  printMessage(message...),
		Line:     line,
		File:     file,
		Endpoint: endpoint,
	}
}

// NewSkipCaller error detail with message values separate per space and skipCaller
func NewSkipCaller(skipCaller int, message ...any) *Error {
	_, file, line, _ := runtime.Caller(skipCaller)
	return &Error{
		Message: printMessage(message...),
		Line:    line,
		File:    file,
	}
}

// NewEndpointSkipCaller error detail with message values separate per space with skipCaller and endpoint
func NewEndpointSkipCaller(skipCaller int, endpoint string, message ...any) *Error {
	_, file, line, _ := runtime.Caller(skipCaller)
	return &Error{
		Message:  printMessage(message...),
		Line:     line,
		File:     file,
		Endpoint: endpoint,
	}
}

// Is validate equal errors, if errorDetail we only consider the errorDetail.Message field
func Is(err, target error) bool {
	errDetail, _ := parseToError(err)
	targetDetail, _ := parseToError(target)
	if helper.IsNotNil(errDetail) && helper.IsNotNil(targetDetail) {
		return equal(*errDetail, *targetDetail)
	}
	return errors.Is(err, target)
}

// IsNot validate not equal errors, if errorDetail we only consider the errorDetail.Message field
func IsNot(err, target error) bool {
	return !Is(err, target)
}

// IsErrorDetail check if error interface is errorDetail
func IsErrorDetail(err error) bool {
	_, errParse := parseToError(err)
	return errParse == nil
}

// Error print the error as a string, genetic implementation of error in go
func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// ParseToError parse value to error struct, if value an is not an Error, we create a new one from it as Error.Message.
func ParseToError(a any) *Error {
	result, err := parseToError(a)
	if helper.IsNotNil(err) {
		return NewSkipCaller(2, a)
	}
	return result
}

func equal(a, b Error) bool {
	return a.Message == b.Message
}

func printMessage(v ...any) string {
	return helper.Sprintln(v...)
}

func parseToError(a any) (*Error, error) {
	var dest Error
	errConvert := helper.ConvertToDest(a, &dest)
	if helper.IsNotNil(errConvert) {
		return nil, errConvert
	}
	return &dest, errConvert
}
