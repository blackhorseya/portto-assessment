package loggerx

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapLogger initializes the logging instance.
func NewZapLogger(options Options) (*zap.Logger, error) {
	level := zap.NewAtomicLevel()
	err := level.UnmarshalText([]byte(options.Level))
	if err != nil {
		return nil, err
	}

	cw := zapcore.Lock(os.Stdout)
	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	enc := zapcore.NewConsoleEncoder(zapConfig)
	if options.Format == "json" {
		zapConfig = zap.NewProductionEncoderConfig()
		enc = zapcore.NewJSONEncoder(zapConfig)
	}

	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(enc, cw, level))

	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	zap.ReplaceGlobals(logger)

	return logger, nil
}
