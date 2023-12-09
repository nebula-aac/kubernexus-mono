package nexuslogger

import (
	"fmt"

	"xorm.io/core"
)

var _ core.ILogger = &XormLogger{}

type XormLogger struct {
	base    *wrapper
	showSQL bool
	level   core.LogLevel
}

// Debug implements core.ILogger.
func (xl *XormLogger) Debug(v ...interface{}) {
	xl.base.logger.Debug(fmt.Sprint(v...))
}

// Error implements core.ILogger.
func (xl *XormLogger) Error(v ...interface{}) {
	xl.base.logger.Error(fmt.Sprint(v...))
}

// Info implements core.ILogger.
func (xl *XormLogger) Info(v ...interface{}) {
	xl.base.logger.Info(fmt.Sprint(v...))
}

// Warn implements core.ILogger.
func (xl *XormLogger) Warn(v ...interface{}) {
	xl.base.logger.Warn(fmt.Sprint(v...))
}

// IsShowSQL implements core.ILogger.
func (xl *XormLogger) IsShowSQL() bool {
	return xl.showSQL
}

func (xl *XormLogger) Debugf(format string, v ...interface{}) {
	xl.base.logger.Debug(fmt.Sprintf(format, v...))
}

func (xl *XormLogger) Errorf(format string, v ...interface{}) {
	xl.base.logger.Error(fmt.Sprintf(format, v...))
}

func (xl *XormLogger) Infof(format string, v ...interface{}) {
	xl.base.logger.Info(fmt.Sprintf(format, v...))
}

func (xl *XormLogger) Warnf(format string, v ...interface{}) {
	xl.base.logger.Warn(fmt.Sprintf(format, v...))
}

// Level returns the logger level
func (xl *XormLogger) Level() core.LogLevel {
	return xl.level
}

// SetLevel sets the logger level
func (xl *XormLogger) SetLevel(l core.LogLevel) {
	xl.level = l
}

// ShowSQL sets whether SQL should be shown or not
func (xl *XormLogger) ShowSQL(show ...bool) {
	if len(show) > 0 {
		xl.showSQL = show[0]
	}
}

func (w *wrapper) XormLogger() core.ILogger {
	return &XormLogger{
		base:  w,
		level: core.LOG_DEBUG, // Set the default log level to debug
	}
}
