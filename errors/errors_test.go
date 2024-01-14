package errors

import (
	"errors"
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestNew(t *testing.T) {
	logger.Info("err:", New("test error detail"))
}

func TestNewE(t *testing.T) {
	logger.Info("err:", NewEndpoint("/test", errors.New("test error detail")))
}

func TestNewESkipCaller(t *testing.T) {
	logger.Info("err:", NewEndpointSkipCaller(1, "/test", errors.New("test error detail")))
}

func TestNewSkipCaller(t *testing.T) {
	logger.Info("err:", NewSkipCaller(1, "test error detail"))
}

func TestIs(t *testing.T) {
	err := errors.New("test")
	target := New("test")
	logger.Info("errors is:", Is(err, target))

	errDetail := New("test")
	targetDetail := New("test")
	logger.Info("errors is:", Is(errDetail, targetDetail))
}

func TestIsNot(t *testing.T) {
	err := errors.New("test")
	target := New("test")
	logger.Info("errors is not:", IsNot(err, target))

	errDetail := New("test")
	targetDetail := New("test")
	logger.Info("errors is not:", IsNot(errDetail, targetDetail))
}

func TestError(t *testing.T) {
	err := New("test error detail")
	logger.Info("err:", err.Error())
}

func TestIsErrorDetail(t *testing.T) {
	err := New("test error detail")
	logger.Info("err:", IsErrorDetail(err))
}

func TestParseToError(t *testing.T) {
	err := New("test error detail").Error()
	logger.Info("err:", ParseToError(err))
	v := "test"
	logger.Info("err:", ParseToError(v))
}
