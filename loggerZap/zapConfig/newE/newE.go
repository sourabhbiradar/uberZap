package main

import "go.uber.org/zap"

func main() {
	config := zap.NewExampleConfig() // no such function
	// only NewDevelopmentConfig() & NewProductionConfig()

	logger := config.Bulid()
	logger.Sync()

}
