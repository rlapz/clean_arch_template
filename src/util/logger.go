package util

import (
	"io"
	"log"
)

type Logger struct {
	stderr *log.Logger
	stdout *log.Logger
}

func NewLogger(outWriter io.Writer, errWriter io.Writer) Logger {
	return Logger{
		stdout: log.New(outWriter, "", log.Ldate|log.Lmicroseconds),
		stderr: log.New(errWriter, "", log.Ldate|log.Lmicroseconds),
	}
}

func (l *Logger) Infof(format string, v ...any) {
	l.stdout.Printf("[INFO]:  "+format+"\n", v...)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.stderr.Printf("[ERROR]: "+format+"\n", v...)
}
