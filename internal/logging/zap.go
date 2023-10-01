package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Zap struct {
    logger *zap.SugaredLogger
}

func NewZap() Zap {
    // create and configure logger core for stdout without JSON formatting
    stdout := zapcore.AddSync(os.Stdout)
    level := zap.NewAtomicLevelAt(zap.InfoLevel)

    devCfg := zap.NewDevelopmentEncoderConfig()
    devCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

    consoleEncoder := zapcore.NewConsoleEncoder(devCfg)

    core := zapcore.NewCore(consoleEncoder, stdout, level)

    // create logger with configured core
    logger := zap.New(core)
    sugared := logger.Sugar()

    return Zap{logger: sugared}
}

func (z Zap) Flush() error {
    return z.logger.Sync()
}

func (z Zap) Info(m string) {
    z.logger.Info(m)
}

func (z Zap) Infow(m string, args ...interface{}) {
    z.logger.Infow(m, args...)
}

func (z Zap) Warning(m string) {
    z.logger.Warn(m)
}

func (z Zap) Warningw(m string, args ...interface{}) {
    z.logger.Warnw(m, args...)
}

func (z Zap) Error(m string) {
    z.logger.Error(m)
}

func (z Zap) Errorw(m string, args ...interface{}) {
    z.logger.Errorw(m, args...)
}

func (z Zap) Fatal(m string) {
    z.logger.Fatal(m)
}

func (z Zap) Fatalw(m string, args ...interface{}) {
    z.logger.Fatalw(m, args...)
}
