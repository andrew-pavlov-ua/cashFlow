package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

// Init initializes the global logger
func Init(serviceName, env string) error {
	var zapLogger *zap.Logger
	var err error

	config := zap.NewProductionConfig()

	if env == "development" || env == "dev" {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.TimeKey = "time"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	// Add service name to all logs
	config.InitialFields = map[string]interface{}{
		"service": serviceName,
	}

	zapLogger, err = config.Build()
	if err != nil {
		return err
	}

	Log = zapLogger.Sugar()
	return nil
}

// Sync flushes any buffered log entries
func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

// Info logs an info message
func Info(args ...interface{}) {
	Log.Info(args...)
}

// Infof logs a formatted info message
func Infof(template string, args ...interface{}) {
	Log.Infof(template, args...)
}

// Infow logs an info message with structured fields
func Infow(msg string, keysAndValues ...interface{}) {
	Log.Infow(msg, keysAndValues...)
}

// Error logs an error message
func Error(args ...interface{}) {
	Log.Error(args...)
}

// Errorf logs a formatted error message
func Errorf(template string, args ...interface{}) {
	Log.Errorf(template, args...)
}

// Errorw logs an error message with structured fields
func Errorw(msg string, keysAndValues ...interface{}) {
	Log.Errorw(msg, keysAndValues...)
}

// Warn logs a warning message
func Warn(args ...interface{}) {
	Log.Warn(args...)
}

// Warnf logs a formatted warning message
func Warnf(template string, args ...interface{}) {
	Log.Warnf(template, args...)
}

// Warnw logs a warning message with structured fields
func Warnw(msg string, keysAndValues ...interface{}) {
	Log.Warnw(msg, keysAndValues...)
}

// Debug logs a debug message
func Debug(args ...interface{}) {
	Log.Debug(args...)
}

// Debugf logs a formatted debug message
func Debugf(template string, args ...interface{}) {
	Log.Debugf(template, args...)
}

// Debugw logs a debug message with structured fields
func Debugw(msg string, keysAndValues ...interface{}) {
	Log.Debugw(msg, keysAndValues...)
}

// Fatal logs a fatal message and exits
func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}

// Fatalf logs a formatted fatal message and exits
func Fatalf(template string, args ...interface{}) {
	Log.Fatalf(template, args...)
}

// Fatalw logs a fatal message with structured fields and exits
func Fatalw(msg string, keysAndValues ...interface{}) {
	Log.Fatalw(msg, keysAndValues...)
}

// Fatalw logs a fatal message with structured fields and exits
func Panic(msg string, err error) {
	Log.Panic(msg, err)
}
