package main

import "go.uber.org/zap"

func main() {
	config := zap.NewProductionConfig()

	config.OutputPaths = []string{
		"stdout",   // on console
		"newP.log", // in newP.log file
	}

	logger, _ := config.Build()
	defer logger.Sync()

	logger.Info("This is an info message")
	logger.Warn("This is a warn message")
	logger.Error("This is an error message")
}
