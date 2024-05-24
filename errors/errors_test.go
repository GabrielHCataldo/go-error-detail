package errors

import (
	"errors"
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestNew(t *testing.T) {
	logger.Info("err:", New("test error detail"))
	logger.Info("err:", New("test error detail", New("sub error message\ntest\ttes2")))
	logger.Info("err:", New(""))
}

func TestNewf(t *testing.T) {
	logger.Info("err:", Newf("%s", "test error detail"))
	logger.Info("err:", Newf("%s %s", "test error detail", New("sub error message\ntest\ttes2")))
	logger.Info("err:", Newf("%s", ""))
}

func TestNewSkipCaller(t *testing.T) {
	err := NewSkipCaller(1, "test error detail")
	logger.Info("err:", err)
	err = NewSkipCaller(1, nil)
	logger.Info("err:", err)
}

func TestNewSkipCallerf(t *testing.T) {
	err := NewSkipCallerf(1, "%s", "test error detail")
	logger.Info("err:", err)
	err = NewSkipCaller(1, nil)
	logger.Info("err:", err)
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

func TestContains(t *testing.T) {
	err := errors.New("test")
	target := New("test")
	logger.Info("errors contains:", Contains(err, target))

	errDetail := New("test")
	targetDetail := New("test2")
	logger.Info("errors contains:", Contains(errDetail, targetDetail))
}

func TestNotContains(t *testing.T) {
	err := errors.New("test")
	target := New("test")
	logger.Info("errors not contains:", NotContains(err, target))

	errDetail := New("test")
	targetDetail := New("test2")
	logger.Info("errors not contains:", NotContains(errDetail, targetDetail))
}

func TestError(t *testing.T) {
	err := New("test error detail")
	logger.Info("err:", err.Error())
}

func TestErrorPrintStack(t *testing.T) {
	err := New("test error detail")
	Details(err).PrintStackTrace()
}

func TestErrorPrintCause(t *testing.T) {
	err := New("test error detail")
	Details(err).PrintCause()
}

func TestErrorGetMessage(t *testing.T) {
	err := New("test error detail")
	logger.Info("err message:", Details(err).GetMessage())
}

func TestErrorGetFile(t *testing.T) {
	err := New("test error detail")
	logger.Info("err file:", Details(err).GetFile())
}

func TestErrorGetLine(t *testing.T) {
	err := New("test error detail")
	logger.Info("err line:", Details(err).GetLine())
}

func TestErrorGetFuncName(t *testing.T) {
	err := New("test error detail")
	logger.Info("err message:", Details(err).GetFuncName())
}

func TestErrorGetDebugStack(t *testing.T) {
	err := New("test error detail")
	logger.Info("err message:", Details(err).GetDebugStack())
}

func TestIsErrorDetail(t *testing.T) {
	err := New("test error detail:", 1, "test empty:", "another test", true)
	logger.Info("err:", IsErrorDetail(err))
}

func TestDetail(t *testing.T) {
	err := New("test error detail:", 1, "test empty:", "another test", true)
	logger.Info("err:", IsErrorDetail(err))
}

func TestDetails(t *testing.T) {
	err := New("test error detail:", nil, "test empty:", "another test: - STACK", true, "empty:")
	logger.Info("err details:", Details(err))
	logger.Info("err details:", Details(nil))
	logger.Info("err details:", Details(errors.New("test")))
	logger.Info("err details:", Details(New("test")))
}
