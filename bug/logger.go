package bug

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger(out io.Writer, prefix string) *Logger {
	return &Logger{Logger: log.New(out, prefix, log.Ldate|log.Ltime|log.Lshortfile)}
}

func NewStdErr(prefix string) *Logger {
	return NewLogger(os.Stderr, prefix)
}

func (l *Logger) Ensure(err error) {
	if err != nil {
		l.Println(err)
	}
}
