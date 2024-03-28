package errors

import (
	"errors"
	"fmt"
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"regexp"
	"runtime/debug"
	"strings"
)

const regexErrorDetail = `\[CAUSE]: \(([^:]+):(\d+)\) ([^:]+): (.+?) \[STACK]:\s*([\s\S]+)`

type ErrorDetail struct {
	file       string
	line       string
	funcName   string
	message    string
	debugStack string
}

// New error with space separated message values, if the message parameter is empty we return a
// nil value, otherwise returns normal value
func New(message ...any) error {
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
func NewSkipCaller(skipCaller int, message ...any) error {
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
		errDetails := Details(err)
		err = errors.New(errDetails.GetMessage())
	}
	if IsErrorDetail(target) {
		errDetails := Details(target)
		target = errors.New(errDetails.GetMessage())
	}
	return helper.IsNotNil(err) && helper.Equals(err, target)
}

// IsNot validate not equal errors
func IsNot(err, target error) bool {
	return !Is(err, target)
}

// Contains validate message target contains on message err
func Contains(err, target error) bool {
	if IsErrorDetail(err) {
		errDetails := Details(err)
		err = errors.New(errDetails.GetMessage())
	}
	if IsErrorDetail(target) {
		errDetails := Details(target)
		target = errors.New(errDetails.GetMessage())
	}
	return helper.IsNotNil(err) && helper.IsNotNil(target) && strings.Contains(err.Error(), target.Error())
}

// NotContains validate message target not contains on message err
func NotContains(err, target error) bool {
	return !Contains(err, target)
}

// Error print the error as a string, genetic implementation of error in go
func (e *ErrorDetail) Error() string {
	return fmt.Sprint("[CAUSE]: ", e.GetCause(), " [STACK]: ", e.debugStack)
}

// PrintStackTrace print red message with detail error and debug stack
func (e *ErrorDetail) PrintStackTrace() {
	logger.ErrorSkipCaller(2, e.debugStack)
}

// PrintCause print red message with cause error
func (e *ErrorDetail) PrintCause() {
	logger.ErrorSkipCaller(2, e.GetCause())
}

// GetCause returns formatted error cause
func (e *ErrorDetail) GetCause() string {
	return fmt.Sprint("(", e.file, ":", e.line, ")", " ", e.funcName, ": ", e.message)
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

// IsErrorDetail check if the error is an ErrorDetail containing the pattern with file name, line, function name,
// message and debugStack
func IsErrorDetail(err error) bool {
	regex := regexp.MustCompile(regexErrorDetail)
	return helper.IsNotNil(err) && regex.MatchString(err.Error())
}

// Details we obtain the values of an error passed in the parameter, and transform them into an ErrorDetail object, if
// the passed parameter is null, the return will also be null and if it is not in the desired errorDetail pattern, we
// return a new ErrDetail with the message being the err passed in the parameter.
func Details(err error) *ErrorDetail {
	if helper.IsNil(err) {
		return nil
	}
	var file string
	var line string
	var funcName string
	var message string
	var debugStack string
	regex := regexp.MustCompile(regexErrorDetail)
	matches := regex.FindStringSubmatch(err.Error())
	if helper.IsNotEmpty(matches) {
		file = matches[1]
		line = matches[2]
		funcName = matches[3]
		message = matches[4]
		debugStack = matches[5]
	} else {
		file, line, funcName = helper.GetCallerInfo(2)
		debugStack = string(debug.Stack())
		message = printMessage(err.Error())
	}
	return &ErrorDetail{
		file:       file,
		line:       line,
		funcName:   funcName,
		message:    message,
		debugStack: debugStack,
	}
}

func printMessage(v ...any) string {
	for i, iv := range v {
		ivError, ok := iv.(error)
		if helper.IsErrorType(iv) && ok {
			errDetail := Details(ivError)
			v[i] = errDetail.message
		}
	}
	msg := helper.Sprintln(v...)
	msg = strings.ReplaceAll(msg, "[STACK]", "")
	msg = strings.ReplaceAll(msg, "[CAUSE]", "")
	re := regexp.MustCompile(`\r?\n`)
	msg = re.ReplaceAllString(msg, " ")
	return msg
}
