package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/nebula-aac/kubernexus-mono/pkg/errors"
	"github.com/sirupsen/logrus"
	gormlogger "gorm.io/gorm/logger"
	"xorm.io/core"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

type Logger interface {
	SetLogLevel(level LogLevel)
	GetLogLevel() LogLevel
	SetLogFormatter(formatter TerminalFormatter)

	Info(msg string, fields ...interface{})
	Debug(msg string, fields ...interface{})
	Warn(err error) error
	Error(err error) error

	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})

	AddHook(hook Hook)
	DatabaseLogger() gormlogger.Interface
	XormLogger() core.ILogger
}

type wrapper struct {
	logger        *logrus.Entry
	logLevel      LogLevel
	InfoStyle     lipgloss.Style
	DebugStyle    lipgloss.Style
	WarnStyle     lipgloss.Style
	ErrorStyle    lipgloss.Style
	Timestamp     bool
	ShowLogColors bool
	logFormatter  TerminalFormatter
}

func (l *wrapper) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

func (l *wrapper) GetLogLevel() LogLevel {
	return l.logLevel
}

func (l *wrapper) SetLogFormatter(formatter TerminalFormatter) {
	l.logFormatter = formatter
}

func (l *wrapper) Info(msg string, fields ...interface{}) {
	logFields := l.extractFields(fields...)
	l.log(logrus.InfoLevel, l.InfoStyle.Bold(true).Foreground(lipgloss.Color("45")).Render(msg), logFields)
}

func (l *wrapper) Infof(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	styledMsg := l.InfoStyle.Bold(true).Foreground(lipgloss.Color("45")).Render(message)
	l.logger.Info(styledMsg)
}

func (l *wrapper) Debug(msg string, fields ...interface{}) {
	logFields := l.extractFields(fields...)
	if l.logLevel >= DebugLevel {
		l.logger.WithFields(logrus.Fields(logFields)).Debug(l.DebugStyle.Bold(true).Foreground(lipgloss.Color("27")).Render(msg))
	}
}

func (l *wrapper) Debugf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	styledMsg := l.DebugStyle.Bold(true).Foreground(lipgloss.Color("27")).Render(message)
	l.Debug(styledMsg)
}

func (l *wrapper) Warn(err error) error {
	if meshErr, ok := err.(*errors.Error); ok {
		logFields := logrus.Fields{
			"Code":          meshErr.Code,
			"ShortDesc":     meshErr.ShortDescription,
			"ProbableCause": meshErr.ProbableCause,
		}
		l.log(logrus.WarnLevel, l.WarnStyle.Bold(true).Foreground(lipgloss.Color("220")).Render(err.Error()), logFields)
	} else {
		l.logger.Warn(l.WarnStyle.Bold(true).Foreground(lipgloss.Color("220")).Render(err.Error()))
	}
	return err
}

func (l *wrapper) Warnf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	styledMsg := l.WarnStyle.Bold(true).Foreground(lipgloss.Color("220")).Render(message)
	l.logger.Warn(styledMsg)
}

func (l *wrapper) Error(err error) error {
	if meshErr, ok := err.(*errors.Error); ok {
		logFields := logrus.Fields{
			"Code":                  meshErr.Code,
			"Severity":              meshErr.Severity,
			"Short Description":     meshErr.ShortDescription,
			"Long Description":      meshErr.LongDescription,
			"Probable Cause":        meshErr.ProbableCause,
			"Suggested Remediation": meshErr.SuggestedRemediation,
		}
		l.log(logrus.ErrorLevel, l.ErrorStyle.Render(err.Error()), logFields)
	} else {
		l.logger.Error(err)
	}
	return err
}

func (l *wrapper) Errorf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	styledMsg := l.ErrorStyle.Render(message)

	l.logger.Error(styledMsg)
}

func (l *wrapper) AddHook(hook Hook) {
	l.logger.Logger.AddHook(hook)
}

func NewLogger(appname string, opts Options) (Logger, error) {
	logrusLogger := logrus.New()

	var formatter TerminalFormatter

	switch opts.Format {
	case JSONFormatter:
		logrusLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	case SyslogFormatter:
		logrusLogger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
			PadLevelText:    true,
			ForceColors:     true,
		})
	case TextFormatter: // Use the custom TerminalFormatter
		formatter := &TerminalFormatter{
			TimestampFormat:  time.StampMilli,
			NoUppercaseLevel: false,
			CallerFirst:      false,
			NoColors:         false,
			ShowFullLevel:    true,
			NoFieldsSpace:    false,
			NoFieldsColors:   false,
			FieldsOrder:      nil,
			TrimMessages:     false,
		}
		logrusLogger.SetFormatter(formatter)
	default:
		return nil, fmt.Errorf("invalid log format: %v", opts.Format)
	}

	logrusLogger.SetOutput(os.Stdout)
	if opts.Output != nil {
		logrusLogger.SetOutput(opts.Output)
	}

	logrusLogger.SetLevel(logrus.InfoLevel)
	if opts.DebugLevel {
		logrusLogger.SetLevel(logrus.DebugLevel)
	}

	entry := logrusLogger.WithFields(logrus.Fields{"app": appname + " "})

	logger := &wrapper{
		logger:        entry,
		logLevel:      InfoLevel, // Set default log level to InfoLevel
		Timestamp:     true,
		ShowLogColors: true,
		logFormatter:  formatter, // Assign the formatter based on the format in the next function call.
	}

	return logger, nil
}

func (l *wrapper) log(level logrus.Level, msg string, fields logrus.Fields) {
	var styledMsg string
	switch level {
	case logrus.DebugLevel:
		styledMsg = debugStyle.Render(msg)
	case logrus.WarnLevel:
		styledMsg = warnStyle.Render(msg)
	case logrus.ErrorLevel:
		styledMsg = errorStyle.Render(msg)
	default:
		// For InfoLevel, use the default info style
		styledMsg = infoStyle.Render(msg)
	}

	message := l.logger.WithFields(fields).Message
	l.logger.Log(level, styledMsg, message)
}

func (l *wrapper) extractFields(fields ...interface{}) logrus.Fields {
	logFields := make(logrus.Fields)
	if len(fields) > 0 {
		for i := 0; i < len(fields); i += 2 {
			if key, ok := fields[i].(string); ok && i+1 < len(fields) {
				logFields[key] = fields[i+1]
			}
		}
	}
	return logFields
}
