package logger

import "go.uber.org/zap"

func NewLogger() (*zap.SugaredLogger, error) {
	rawLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return rawLogger.Sugar(), nil
}
