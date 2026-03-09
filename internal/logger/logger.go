package logger

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func Init(env string) error {
	var cfg zap.Config

	if env == "production" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	logger, err := cfg.Build()
	if err != nil {
		return err
	}

	Log = logger.Sugar()

	return nil
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
