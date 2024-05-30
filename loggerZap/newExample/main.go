package main

import (
	"errors"

	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync() // lets any buffer logs to be entered no matter how program exits.

	logger.Debug("First zap log", zap.String("zap", "NewExapmle"))

	logger.Info("Second zap log",
		zap.Int("user ID:", 3210),
		zap.String("name", "Abc"),
		zap.Error(errors.New("dummy error")), // automatically add `error :` field
	)
}
