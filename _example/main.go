package main

import (
	errors2 "errors"
	"github.com/GabrielHCataldo/go-errors/errors"
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
	err := simple()
	logger.Info("simple err:", err)
	logger.Info("simple err msg:", errors.Details(err).GetMessage())
	logger.Info("simple err file:", errors.Details(err).GetFile())
	logger.Info("simple err line:", errors.Details(err).GetLine())
	logger.Info("simple err func:", errors.Details(err).GetFuncName())
	errors.Details(err).PrintCause()
	errors.Details(err).PrintStackTrace()
	is()
	nilErr()
	details(err)
}

func simple() error {
	return errors.New("error by message with any value", 2, true)
}

func details(err error) {
	errDetail := errors.Details(err)
	logger.Info("message:", errDetail.GetMessage())
	logger.Info("file:", errDetail.GetFile())
	logger.Info("line:", errDetail.GetLine())
	logger.Info("funcName:", errDetail.GetFuncName())
	logger.Info("debugStack:", errDetail.GetDebugStack())
	logger.Info("cause:", errDetail.GetCause())
	errDetail.PrintStackTrace()
	errDetail.PrintCause()
}

func nilErr() {
	err := errors.NewSkipCaller(2, nil)
	logger.Error("nilErr err:", err == nil, helper.IsNil(err))
}

func is() {
	err2 := errors2.New("test")
	err := errors.New("error by message with any value", 2)
	target := errors.New("test")
	logger.Info("errors is:", errors.Is(err, target))
	logger.Info("errors2 is:", errors.Is(err2, target))
	logger.Info("errors is not:", errors.IsNot(err, target))
	logger.Info("is error detail?", errors.IsErrorDetail(err))
	logger.Info("is error detail?", errors.IsErrorDetail(err2))
}
