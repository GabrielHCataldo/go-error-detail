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
	File string `json:"file,omitempty"`
	// Line from caller new error
	Line int `json:"line,omitempty"`
	// Message error
	Message string `json:"message,omitempty"`
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

// NewErr new error interface
func NewErr(message ...any) error {
	return errors.New(printMessage(message...))
}

// Is validate equal errors, if ErrorDetail we only consider the ErrorDetail.Message field
func Is(err, target error) bool {
	if IsErrorDetail(err) && IsErrorDetail(target) {
		errDetail, _ := parseErrorToDetail(err)
		targetDetail, _ := parseErrorToDetail(target)
		return equal(*errDetail, *targetDetail)
	}
	return errors.Is(err, target)
}

// IsErrorDetail check if error interface is ErrorDetail
func IsErrorDetail(err error) bool {
	_, errParse := parseErrorToDetail(err)
	return errParse == nil
}

// Error print the error as a string, genetic implementation of error in go
func (e *ErrorDetail) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func equal(a, b ErrorDetail) bool {
	return a.Message == b.Message
}

func printMessage(v ...any) string {
	return strings.Replace(fmt.Sprintln(v...), "\n", "", -1)
}

func parseErrorToDetail(err error) (*ErrorDetail, error) {
	var dest ErrorDetail
	errConvert := helper.ConvertToDest(err, &dest)
	return &dest, errConvert
}
