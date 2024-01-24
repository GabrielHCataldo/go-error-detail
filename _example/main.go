package main

import (
	"github.com/GabrielHCataldo/go-errors/errors"
	"github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
	err := simple()
	logger.Error("simple err:", err)
	all()
}

func simple() error {
	return errors.New("error by message with any value", 2, true)
}

func all() {
	err := errors.New("error by message with any value", 2)
	logger.Error(err)
	err = errors.NewSkipCaller(1, "error by message with any value", 2)
	logger.Error(err)
	err = errors.New("test")
	target := errors.New("test")
	logger.Info("errors is:", errors.Is(err, target))
	logger.Info("errors is not:", errors.IsNot(err, target))
	logger.Info("is error detail?", errors.IsErrorDetail(err))
	file, line, funcName, message := errors.GetErrorDetails(err)
	logger.Info("error details:", file, line, funcName, message)
}
