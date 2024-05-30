package main

import "go.uber.org/zap"

func main() {
	config := zap.NewDevelopmentConfig() // development config

	config.OutputPaths = []string{
		"stdout",   // on console
		"newD.log", // in newD.log file
	}

	logger, _ := config.Build()
	defer logger.Sync()

	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warn message")
	logger.Error("This is an error message")
}
