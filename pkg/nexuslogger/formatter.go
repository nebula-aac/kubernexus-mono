package nexuslogger

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
)

type Format int

const (
	TextFormatter Format = iota
	JSONFormatter
	SyslogFormatter
)

const (
	colorGray   = "#98c379"
	colorYellow = "#e5c07b"
	colorPurple = "#c678dd"
	colorRed    = "#e06c75"
	colorBlue   = "#61afef"
)

var (
	infoStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(colorBlue)).Bold(true)
	warnStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(colorYellow)).Bold(true)
	debugStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(colorPurple)).Bold(true)
	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(colorRed)).Bold(true)
)

type TerminalFormatter struct {
	TimestampFormat  string
	NoUppercaseLevel bool
	CallerFirst      bool
	NoColors         bool
	ShowFullLevel    bool
	NoFieldsSpace    bool
	NoFieldsColors   bool
	FieldsOrder      []string
	TrimMessages     bool
}

// Format formats the log entry using lipgloss styles based on the log level
func (f *TerminalFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := getColorByLevel(entry.Level)

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.StampMilli
	}

	// output buffer
	b := &bytes.Buffer{}

	// write time
	b.WriteString(entry.Time.Format(timestampFormat))

	// write level
	var level string
	if f.NoUppercaseLevel {
		level = entry.Level.String()
	} else {
		level = strings.ToUpper(entry.Level.String())
	}

	if !f.NoColors {
		level = lipgloss.NewStyle().Foreground(lipgloss.Color(levelColor)).Render(level)
	}

	if f.CallerFirst {
		f.writeCaller(b, entry)
	}

	b.WriteString(" [")
	if f.ShowFullLevel {
		b.WriteString(level)
	} else {
		b.WriteString(level[:4])
	}
	b.WriteString("]")

	if !f.NoFieldsSpace {
		b.WriteString(" ")
	}

	if !f.NoColors && f.NoFieldsColors {
		b.WriteString(lipgloss.Style{}.String())
	}

	// write fields
	if f.FieldsOrder == nil {
		f.writeFields(b, entry)
	} else {
		f.writeOrderedFields(b, entry)
	}

	if f.NoFieldsSpace {
		b.WriteString(" ")
	}

	if !f.NoColors && !f.NoFieldsColors {
		b.WriteString(lipgloss.Style{}.String())
	}

	// write message
	if f.TrimMessages {
		b.WriteString(strings.TrimSpace(entry.Message))
	} else {
		b.WriteString(entry.Message)
	}

	if !f.CallerFirst {
		f.writeCaller(b, entry)
	}

	b.WriteByte('\n')

	return b.Bytes(), nil
}

// getColorByLevel returns the lipgloss color code (as an integer) based on the log level.
func getColorByLevel(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return colorGray
	case logrus.WarnLevel:
		return colorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return colorRed
	default:
		return colorBlue
	}
}

func (f *TerminalFormatter) writeFields(b *bytes.Buffer, entry *logrus.Entry) {
	for key, value := range entry.Data {
		fmt.Fprintf(b, " %s=%v", key, value)
	}
}

func (f *TerminalFormatter) writeOrderedFields(b *bytes.Buffer, entry *logrus.Entry) {
	for _, key := range f.FieldsOrder {
		if value, ok := entry.Data[key]; ok {
			fmt.Fprintf(b, " %s=%v", key, value)
		}
	}
}

func (f *TerminalFormatter) writeCaller(b *bytes.Buffer, entry *logrus.Entry) {
	if entry.HasCaller() {
		// Use Sprint instead of Fprintf to avoid adding a space after the caller.
		b.WriteString(fmt.Sprint(" ", entry.Caller.File, ":", entry.Caller.Line))
	}
}
