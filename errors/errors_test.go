package errors

import (
	"errors"
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestNew(t *testing.T) {
	logger.Info("err:", New("test error detail"))
	logger.Info("err:", New(""))
}

func TestNewSkipCaller(t *testing.T) {
	logger.Info("err:", NewSkipCaller(1, "test error detail"))
	logger.Info("err:", NewSkipCaller(1, ""))
}

func TestIs(t *testing.T) {
	err := errors.New("test")
	target := New("test")
	logger.Info("errors is:", Is(err, target))

	errDetail := New("test")
	targetDetail := New("test")
	logger.Info("errors is:", Is(errDetail, targetDetail))
	logger.Info("errors is:", Is(nil, nil))
}

func TestIsNot(t *testing.T) {
	err := errors.New("test")
	target := New("test2")
	logger.Info("errors is not:", IsNot(err, target))

	errDetail := New("test")
	targetDetail := New("test2")
	logger.Info("errors is not:", IsNot(errDetail, targetDetail))
}

func TestError(t *testing.T) {
	err := New("test error detail")
	logger.Info("err:", err.Error())
}

func TestErrorPrintStack(t *testing.T) {
	err := New("test error detail")
	err.PrintStack()
}

func TestErrorGetMessage(t *testing.T) {
	err := New("test error detail")
	err.GetMessage()
	logger.Info("err message:", err.GetMessage())
}

func TestErrorGetFile(t *testing.T) {
	err := New("test error detail")
	err.GetMessage()
	logger.Info("err message:", err.GetFile())
}

func TestErrorGetLine(t *testing.T) {
	err := New("test error detail")
	err.GetMessage()
	logger.Info("err message:", err.GetLine())
}

func TestErrorGetFuncName(t *testing.T) {
	err := New("test error detail")
	err.GetMessage()
	logger.Info("err message:", err.GetFuncName())
}

func TestErrorGetDebugStack(t *testing.T) {
	err := New("test error detail")
	err.GetMessage()
	logger.Info("err message:", err.GetDebugStack())
}

func TestIsErrorDetail(t *testing.T) {
	err := New("test error detail")
	logger.Info("err:", IsErrorDetail(err))
}

func TestGetErrorDetails(t *testing.T) {
	err := New("test error detail").Error()
	file, line, funcName, message := GetErrorDetails(errors.New(err))
	logger.Info("details:", file, line, funcName, message)
	v := "test"
	file, line, funcName, message = GetErrorDetails(errors.New(v))
	logger.Info("details:", file, line, funcName, message)
	file, line, funcName, message = GetErrorDetails(nil)
	logger.Info("details:", file, line, funcName, message)
}
