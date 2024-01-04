package main

import (
	"encoding/json"
	"github.com/GabrielHCataldo/go-error-detail/errors"
	"github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
	err := simple()
	logger.Error("simple result:", err)
}

func simple() error {
	return errors.New("error by message with any value", 2, true)
}

func all() {
	err := errors.New("error by message with any value", 2)
	logger.Error(err)
	err = errors.NewE("/endpoint", "error by message with any value", 2)
	logger.Error(err)
	err = errors.NewSkipCaller(1, "error by message with any value", 2)
	logger.Error(err)
	err = errors.NewESkipCaller(1, "/endpoint", "error by message with any value", 2)
	logger.Error(err)
	_, errJson := json.Marshal(make(chan int))
	logger.Error(errors.NewByErr(errJson))
	logger.Error(errors.NewByErr(errJson).Error())
	logger.Error(errors.NewByErrSkipCaller(1, errJson))
	logger.Error(errors.NewEByErrSkipCaller(1, errJson, "/endpoint"))
	logger.Error(errors.NewEByErrSkipCaller(1, errJson, "/endpoint"))
}
