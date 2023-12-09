package errors

import (
	"fmt"
	"runtime"
)

type Severity int

type Error struct {
	Code                 string
	Severity             Severity
	ShortDescription     []string
	LongDescription      []string
	ProbableCause        []string
	SuggestedRemediation []string
	Stack                []string
}

const (
	SeverityLow Severity = iota
	SeverityMedium
	SeverityHigh
)

const (
	Emergency Severity = iota // System unusable
	None                      // None severity
	Alert                     // Immediate action needed
	Critical                  // Critical condition—default level
	Fatal                     // Fatal condition
)

func New(code string, severity Severity, shortDesc, longDesc, probableCause, remediation []string) *Error {
	stack := getCallers(3)
	return &Error{
		Code:                 code,
		Severity:             severity,
		ShortDescription:     shortDesc,
		LongDescription:      longDesc,
		ProbableCause:        probableCause,
		SuggestedRemediation: remediation,
		Stack:                stack,
	}
}

func Errorf(code string, severity Severity, format string, args ...interface{}) *Error {
	stack := getCallers(3)
	return &Error{
		Code:                 code,
		Severity:             severity,
		ShortDescription:     []string{fmt.Sprintf(format, args...)},
		LongDescription:      nil,
		ProbableCause:        nil,
		SuggestedRemediation: nil,
		Stack:                stack,
	}
}

func (e *Error) Error() string {
	if len(e.ShortDescription) > 0 {
		return e.ShortDescription[0]
	}
	return ""
}

func (e *Error) StackTrace() []string {
	return e.Stack
}

func getCallers(skip int) []string {
	callers := make([]string, 0)
	pc := make([]uintptr, 100)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, frame.Function)
	}
	return callers
}
