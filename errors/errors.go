package errors

import (
	"errors"
	"fmt"
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"regexp"
	"runtime/debug"
)

const regexErrorDetail = `\[(.+?):(\d+)] (.+?): (.+)`

type ErrorDetail struct {
	file       string
	line       string
	funcName   string
	message    string
	debugStack string
}

// New error with space separated message values, if the message parameter is empty we return a
// nil value, otherwise returns normal value
func New(message ...any) *ErrorDetail {
	msg := printMessage(message...)
	if helper.IsEmpty(msg) {
		return nil
	}
	file, line, funcName := helper.GetCallerInfo(2)
	debugStack := debug.Stack()
	return &ErrorDetail{
		file:       file,
		line:       line,
		funcName:   funcName,
		message:    msg,
		debugStack: string(debugStack),
	}
}

// NewSkipCaller error with message values separate per space and skipCaller, if the message parameter is empty we return a
// nil value, otherwise returns normal value
func NewSkipCaller(skipCaller int, message ...any) *ErrorDetail {
	msg := printMessage(message...)
	if helper.IsEmpty(msg) {
		return nil
	}
	file, line, funcName := helper.GetCallerInfo(skipCaller + 1)
	debugStack := debug.Stack()
	return &ErrorDetail{
		file:       file,
		line:       line,
		funcName:   funcName,
		message:    msg,
		debugStack: string(debugStack),
	}
}

// Is validate equal errors
func Is(err, target error) bool {
	if IsErrorDetail(err) {
		_, _, _, message := GetErrorDetails(err)
		err = errors.New(message)
	}
	if IsErrorDetail(target) {
		_, _, _, message := GetErrorDetails(target)
		target = errors.New(message)
	}
	return helper.IsNotNil(err) && helper.Equals(err, target)
}

// IsNot validate not equal errors
func IsNot(err, target error) bool {
	return !Is(err, target)
}

// Error print the error as a string, genetic implementation of error in go
func (e *ErrorDetail) Error() string {
	return fmt.Sprint("[", e.file, ":", e.line, "]", " ", e.funcName, ": ", e.message)
}

// PrintStack print red message with detail error and debug stack
func (e *ErrorDetail) PrintStack() {
	logger.ErrorSkipCaller(2, "Error:", e.Error(), "Stack:", e.debugStack)
}

// GetMessage returns the value of the error message field
func (e *ErrorDetail) GetMessage() string {
	return e.message
}

// GetFile returns the value of the error file field
func (e *ErrorDetail) GetFile() string {
	return e.file
}

// GetLine returns the value of the error line field
func (e *ErrorDetail) GetLine() int {
	return helper.SimpleConvertToInt(e.line)
}

// GetFuncName returns the value of the error funcName field
func (e *ErrorDetail) GetFuncName() string {
	return e.funcName
}

// GetDebugStack returns the value of the error debugStack field
func (e *ErrorDetail) GetDebugStack() string {
	return e.debugStack
}

// IsErrorDetail check if the error is an ErrorDetail containing the pattern with file name, line, function name
// and message
func IsErrorDetail(err error) bool {
	regex := regexp.MustCompile(regexErrorDetail)
	return helper.IsNotNil(err) && regex.MatchString(err.Error())
}

// GetErrorDetails we obtain the values of an error detail separately, if the parameter is nil we return all empty
// values, and if the passed error parameter is not in the desired pattern, we return only the filled message and the
// rest empty.
func GetErrorDetails(err error) (file, line, funcName, message string) {
	if helper.IsNil(err) {
		return
	}
	message = err.Error()
	regex := regexp.MustCompile(regexErrorDetail)
	matches := regex.FindStringSubmatch(message)
	if helper.IsNotEmpty(matches) {
		file = matches[1]
		line = matches[2]
		funcName = matches[3]
		message = matches[4]
	}
	return
}

func printMessage(v ...any) string {
	return helper.Sprintln(v...)
}
