package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init(dev bool) error {
	var err error

	if dev {
		Log, err = zap.NewDevelopment()
	} else {
		Log, err = zap.NewProduction()

	}

	return err

}
