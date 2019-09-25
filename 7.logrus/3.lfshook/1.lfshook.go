package main

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "/tmp/info.log",
		logrus.ErrorLevel: "/tmp/error.log",
	}

	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return Log
}

func main() {
	NewLogger()

	Log.WithFields(logrus.Fields{
		"animal": "11111111111111111111111111111111111111111111",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
