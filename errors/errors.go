package errors

import (
	"errors"
	"fmt"
	"github.com/GabrielHCataldo/go-helper/helper"
	"regexp"
)

const regexErrorDetail = `\[(.+?):(\d+)] (.+?): (.+)`

type errorDetail struct {
	file     string
	line     string
	funcName string
	message  string
}

// New error with space separated message values, if the message parameter is empty we return a
// nil value, otherwise returns normal value
func New(message ...any) error {
	msg := printMessage(message...)
	if helper.IsEmpty(msg) {
		return nil
	}
	file, line, funcName := helper.GetCallerInfo(2)
	return &errorDetail{
		file:     file,
		line:     line,
		funcName: funcName,
		message:  msg,
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
	return &errorDetail{
		file:     file,
		line:     line,
		funcName: funcName,
		message:  msg,
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
func (e *errorDetail) Error() string {
	return fmt.Sprint("[", e.file, ":", e.line, "]", " ", e.funcName, ": ", e.message)
}

// IsErrorDetail check if the error is an errorDetail containing the pattern with file name, line, function name
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
