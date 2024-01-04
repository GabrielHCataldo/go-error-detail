package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GabrielHCataldo/go-helper/helper"
	"runtime"
	"strings"
)

type ErrorDetail struct {
	// File name file from caller new error
	File string `json:"file"`
	// Line from caller new error
	Line int `json:"line"`
	// Message error
	Message string `json:"message"`
	// Endpoint from error
	Endpoint string `json:"endpoint,omitempty"`
}

// New error detail with message values separate per space
func New(message ...any) *ErrorDetail {
	_, file, line, _ := runtime.Caller(1)
	return &ErrorDetail{
		Message: printMessage(message...),
		Line:    line,
		File:    file,
	}
}

// NewByErr error detail by error, if error is nil return nil
func NewByErr(err error) *ErrorDetail {
	var e *ErrorDetail
	if helper.IsNotNil(err) {
		_, file, line, _ := runtime.Caller(1)
		e = &ErrorDetail{
			Message: err.Error(),
			Line:    line,
			File:    file,
		}
	}
	return e
}

// NewE error detail with endpoint and message
func NewE(endpoint string, message ...any) *ErrorDetail {
	_, file, line, _ := runtime.Caller(1)
	return &ErrorDetail{
		Message:  printMessage(message...),
		Line:     line,
		File:     file,
		Endpoint: endpoint,
	}
}

// NewSkipCaller error detail with message values separate per space and skipCaller
func NewSkipCaller(skipCaller int, message ...any) *ErrorDetail {
	_, file, line, _ := runtime.Caller(skipCaller)
	return &ErrorDetail{
		Message: printMessage(message...),
		Line:    line,
		File:    file,
	}
}

// NewESkipCaller error detail with message values separate per space with skipCaller and endpoint
func NewESkipCaller(skipCaller int, endpoint string, message ...any) *ErrorDetail {
	_, file, line, _ := runtime.Caller(skipCaller)
	return &ErrorDetail{
		Message:  printMessage(message...),
		Line:     line,
		File:     file,
		Endpoint: endpoint,
	}
}

// NewByErrSkipCaller error detail by error with skipCaller, if error is nil return nil
func NewByErrSkipCaller(skipCaller int, err error) *ErrorDetail {
	var e *ErrorDetail
	if helper.IsNotNil(err) {
		_, file, line, _ := runtime.Caller(skipCaller)
		e = &ErrorDetail{
			Message: err.Error(),
			Line:    line,
			File:    file,
		}
	}
	return e
}

// NewEByErrSkipCaller error detail by error with skipCaller, if error is nil return nil
func NewEByErrSkipCaller(skipCaller int, err error, endpoint string) *ErrorDetail {
	var e *ErrorDetail
	if helper.IsNotNil(err) {
		_, file, line, _ := runtime.Caller(skipCaller)
		e = &ErrorDetail{
			Message:  err.Error(),
			Line:     line,
			File:     file,
			Endpoint: endpoint,
		}
	}
	return e
}

func NewError(message ...any) error {
	return errors.New(printMessage(message...))
}

func printMessage(v ...any) string {
	return strings.Replace(fmt.Sprintln(v...), "\n", "", -1)
}

func (e *ErrorDetail) Error() string {
	r := ""
	if helper.IsNotEmpty(e) {
		b, _ := json.Marshal(e)
		r = string(b)
	}
	return r
}
