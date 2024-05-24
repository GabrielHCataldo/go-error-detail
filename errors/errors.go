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

// New is a function that creates a new error with additional error details.
// It takes in variadic arguments `args` of any type and builds a message using `buildMessage`.
// It then obtains the caller information using `helper.GetCallerInfo` and the current stack trace using `debug.Stack()`.
// It returns an instance of `ErrorDetail` which contains the file, line number, function name, message, and debug stack.
// The caller information and debug stack are used for printing the stack trace.
//
// Example usage:
//
//	err := New("test error detail")
//	fmt.Println(err.Error()) // Output: [CAUSE]: (filename:line) function: test error detail [STACK]: stack trace
//
// Example usage:
//
//	err := Newf("%s", "test error detail")
//	fmt.Println(err.Error()) // Output: [CAUSE]: (filename:line) function: test error detail [STACK]: stack trace
func New(args ...any) error {
	msg := buildMessage(args...)
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

// Newf is a function that creates a new error with additional error details.
// It takes in a format string and variadic arguments `args` of any type and builds a message using `buildMessageByFormat`.
// It then obtains the caller information using `helper.GetCallerInfo` and the current stack trace using `debug.Stack()`.
// It returns an instance of `ErrorDetail` which contains the file, line number, function name, message, and debug stack.
// The caller information and debug stack are used for printing the stack trace.
//
// Usage:
//
//	err := Newf("%s", "test error detail")
//	fmt.Println(err.Error()) // Output: [CAUSE]: (filename:line) function: test error detail [STACK]: stack trace
func Newf(format string, args ...any) error {
	msg := buildMessageByFormat(format, args...)
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

// NewSkipCaller is a function that creates a new error with additional error details, skipping a certain number of callers.
// It takes in an integer argument `skipCaller` to specify the number of callers to skip.
// It also takes in variadic arguments `args` of any type and builds a message using `buildMessage`.
// It then obtains the caller information using `helper.GetCallerInfo` and the current stack trace using `debug.Stack()`.
// It returns an instance of `ErrorDetail` which contains the file, line number, function name, message, and debug stack.
// The caller information and debug stack are used for printing the stack trace.
//
// Usage:
//
//	err := NewSkipCaller(1, "test error detail")
//	logger.Info("err:", err)
//	err = NewSkipCaller(1, nil)
//	logger.Info("err:", err)
//
// Example:
//
//	// Output: [CAUSE]: (filename:line) function: test error detail [STACK]: stack trace
func NewSkipCaller(skipCaller int, args ...any) error {
	msg := buildMessage(args...)
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

// NewSkipCallerf is a function that creates a new error with additional error details,
// skipping a certain number of callers and formatting the message using a format string.
// It takes in an integer argument `skipCaller` to specify the number of callers to skip.
// It also takes in a format string and variadic arguments `args` of any type to build the formatted message
// using `buildMessageByFormat`.
// It then obtains the caller information using `helper.GetCallerInfo` and the current stack trace using `debug.Stack()`.
// It returns an instance of `ErrorDetail` which contains the file, line number, function name, formatted message, and debug stack.
// The caller information and debug stack are used for printing the stack trace.
//
// Usage:
//
//	err := NewSkipCallerf(1, "%s", "test error detail")
//	logger.Info("err:", err)
//	err = NewSkipCallerf(1, "%s")
//	logger.Info("err:", err)
//
// Example:
//
//	// Output: [CAUSE]: (filename:line) function: test error detail [STACK]: stack trace
func NewSkipCallerf(skipCaller int, format string, args ...any) error {
	msg := buildMessageByFormat(format, args...)
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

// Error is a method of the ErrorDetail struct that returns a formatted string representation of the error.
// It returns a string in the format "[CAUSE]: (filename:line) function: message [STACK]: stack trace",
// where filename represents the name of the file where the error occurred,
// line represents the line number in the file where the error occurred,
// function represents the name of the function where the error occurred,
// message represents the specific error message,
// and stack trace represents the stack trace at the time the error occurred.
// This method is used for printing the error message along with the stack trace.
func (e *ErrorDetail) Error() string {
	return fmt.Sprint("[CAUSE]: ", e.GetCause(), " [STACK]: ", e.debugStack)
}

// PrintStackTrace is a method of the ErrorDetail struct that logs the debugStack using logger.ErrorSkipCaller.
// It takes no arguments and does not return anything.
// This method is used for printing the stack trace.
// Example usage:
//
//	err := New("test error detail")
//	Details(err).PrintStackTrace()
func (e *ErrorDetail) PrintStackTrace() {
	logger.ErrorSkipCaller(2, e.debugStack)
}

// PrintCause is a method of the ErrorDetail struct that logs the cause of the error using logger.ErrorSkipCaller.
// It takes no arguments and does not return anything.
// This method is used for logging the cause of the error.
func (e *ErrorDetail) PrintCause() {
	logger.ErrorSkipCaller(2, e.GetCause())
}

// GetCause is a method of the ErrorDetail struct that returns a formatted string representation of the cause of the error.
// It returns a string in the format "(filename:line) function: message",
// where filename represents the name of the file where the error occurred,
// line represents the line number in the file where the error occurred,
// function represents the name of the function where the error occurred,
// and message represents the specific error message.
// This method is used for getting the cause of the error.
func (e *ErrorDetail) GetCause() string {
	return fmt.Sprint("(", e.file, ":", e.line, ")", " ", e.funcName, ": ", e.message)
}

// GetMessage is a method of the ErrorDetail struct that returns the error message.
// It returns a string representing the error message stored in the ErrorDetail instance.
// This method is used for retrieving the error message.
// Example usage:
//
//	err := New("test error detail")
//	message := Details(err).GetMessage()
//	fmt.Println(message) // Output: test error detail
func (e *ErrorDetail) GetMessage() string {
	return e.message
}

// GetFile is a method of the ErrorDetail struct that returns the name of the file where the error occurred.
// It returns a string representing the name of the file.
// This method is used for retrieving the file name.
// Example usage:
//
//	err := New("test error detail")
//	file := Details(err).GetFile()
//	fmt.Println(file) // Output: test_error.go
func (e *ErrorDetail) GetFile() string {
	return e.file
}

// GetLine is a method of the ErrorDetail struct that returns the line number
// where the error occurred as an integer. It uses the SimpleConvertToInt function
// from the helper package to convert the line number from a string to an integer.
// This method is used for retrieving the line number of the error.
// Example usage:
//
//	err := New("test error detail")
//	line := Details(err).GetLine()
//	fmt.Println(line) // Output: 42
func (e *ErrorDetail) GetLine() int {
	return helper.SimpleConvertToInt(e.line)
}

// GetFuncName is a method of the ErrorDetail struct that returns the name of the function where the error occurred.
func (e *ErrorDetail) GetFuncName() string {
	return e.funcName
}

// GetDebugStack is a method of the ErrorDetail struct that returns the debug stack trace.
// It returns a string representing the debug stack trace stored in the ErrorDetail instance.
// This method is used for retrieving the debug stack trace.
// Example usage:
//
//	err := New("test error detail")
//	stack := Details(err).GetDebugStack()
//	fmt.Println(stack) // Output: goroutine 1 [running]:
//	stack_trace_test.(*ErrorDetail).GetDebugStack(0x1140d81e0, 0xc00003e1b0, 0x104126d, 0xc00000c080)
//		/path/to/file.go:42 +0x136
//	stack_trace_test.TestErrorGetDebugStack(0xc00003e1b0)
//		/path/to/file_test.go:11 +0x38
//	testing.tRunner(0xc00003e1b0, 0x118b556)
//		/path/to/testing/testing.go:1233 +0x162
//	created by testing.(*T).Run
//		/path/to/testing/testing.go:1274 +0x298
//
// Note: The actual stack trace content may vary depending on the environment and program execution.
func (e *ErrorDetail) GetDebugStack() string {
	return e.debugStack
}

// Is a function that checks if the given `err` matches the given `target` error.
// If both `err` and `target` are instances of ErrorDetail, it extracts the error message from each
// and creates new errors with the extracted messages. This is to ensure that the error messages are comparable.
// It then compares the modified `err` and `target` using the helper functions `helper.IsNotNil` and `helper.Equals`.
// Returns true if `err` is not nil and is equal to `target`, false otherwise.
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

// IsNot is a function that checks if the given `err` is not equal to the given `target` error.
// It negates the result of the function `Is(err, target)`. Returns true if `err` is not equal to `target`,
// false otherwise.
func IsNot(err, target error) bool {
	return !Is(err, target)
}

// Contains is a function that checks if an error instance or its string representation contains the target error.
// If the input error is an instance of ErrorDetail, it extracts the error message and updates the input error to a
// new instance of errors.New().
// Finally, it checks if both the input error and target error are not nil and if the string representation of the
// input error contains the string representation of the target error.
// It returns a boolean value indicating whether the input error contains the target error.
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

// NotContains is a function that checks if an error instance or its string representation does not contain the
// target error.
// It calls the Contains function to check if the error does contain the target error and negates the result.
// It returns a boolean value indicating whether the error does not contain the target error.
func NotContains(err, target error) bool {
	return !Contains(err, target)
}

// IsErrorDetail is a function that checks if the given `err` is an instance of ErrorDetail.
// It does this by using a regular expression pattern to match the string representation of `err`.
// Returns true if `err` is not nil and matches the pattern, false otherwise.
func IsErrorDetail(err error) bool {
	regex := regexp.MustCompile(regexErrorDetail)
	return helper.IsNotNil(err) && regex.MatchString(err.Error())
}

// Details is a function that takes in an error and returns an instance of *ErrorDetail.
// If the input error is nil, it returns nil.
// It initializes variables file, line, funcName, message, and debugStack to empty strings.
// It uses a regular expression to match the error message of the input error against the regexErrorDetail pattern.
// If there is a match, it extracts the file, line, funcName, message, and debugStack from the error message.
// Otherwise, it obtains the caller information using helper.GetCallerInfo(2) and the current stack trace using debug.Stack().
// It builds the message using buildMessage(err.Error()).
// It returns a pointer to a newly created ErrorDetail struct, with the extracted/obtained information as its field values.
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
		message = buildMessage(err.Error())
	}
	return &ErrorDetail{
		file:       file,
		line:       line,
		funcName:   funcName,
		message:    message,
		debugStack: debugStack,
	}
}

// buildMessage is a function that takes in variadic arguments `v` of any type and builds a message by
// using the helper.Sprintln, filterMsg, and cleanMessage functions.
// It returns the cleaned message string.
func buildMessage(v ...any) string {
	return cleanMessage(helper.Sprintln(filterMsg(v...)...))
}

// buildMessageByFormat is a function that takes in a format string and variadic arguments `v` of any type.
// It uses fmt.Sprintf to format the message string using the format and the filtered arguments.
// It then passes the formatted message to cleanMessage to remove the "[STACK]" and "[CAUSE]" tags,
// and replace all newline characters with a space.
// It returns the cleaned message string.
func buildMessageByFormat(format string, v ...any) string {
	return cleanMessage(fmt.Sprintf(format, filterMsg(v...)...))
}

// cleanMessage is a function that takes in a message string and removes "[STACK]" and "[CAUSE]" tags.
// It uses strings.ReplaceAll to replace all occurrences of "[STACK]" and "[CAUSE]" with an empty string.
// It uses the regular expression \r?\n to match all newline characters and replace them with a space.
// It returns the cleaned message string.
func cleanMessage(msg string) string {
	msg = strings.ReplaceAll(msg, "[STACK]", "")
	msg = strings.ReplaceAll(msg, "[CAUSE]", "")
	re := regexp.MustCompile(`\r?\n`)
	return re.ReplaceAllString(msg, " ")
}

// filterMsg iterates over variadic arguments and extracts error messages if the arguments are of error type.
// It utilizes the Details function to extract the error message from error types.
// It returns the modified arguments with extracted error messages.
func filterMsg(v ...any) []any {
	for i, iv := range v {
		ivError, ok := iv.(error)
		if helper.IsErrorType(iv) && ok {
			errDetail := Details(ivError)
			v[i] = errDetail.message
		}
	}
	return v
}
