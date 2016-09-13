package logsip

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger struct {
	Level Level
	mu    sync.Mutex
	*log.Logger
}

// New returns a new logger
func New() *Logger {
	return &Logger{
		Level:  InfoLevel,
		Logger: log.New(os.Stderr, "", 0),
	}
}

// Panic works just like log.Panic, but with a Red "==> PANIC:" prefix.
func (l *Logger) Panic(v ...interface{}) {
	if l.Level >= PanicLevel {
		l.Logger.Panic(PanicLevel.String() + fmt.Sprint(v...))
	}
}

// Fatal works just like log.Fatal, but with a Red "==> FATAL:" prefix.
func (l *Logger) Fatal(v ...interface{}) {
	if l.Level >= FatalLevel {
		l.Logger.Fatal(FatalLevel.String() + fmt.Sprint(v...))
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.Level >= ErrorLevel {
		l.Logger.Print(ErrorLevel.String() + fmt.Sprint(v...))
	}
}

// Warn works just like log.Print, but with a Yellow "==> WARN:" prefix.
func (l *Logger) Warn(v ...interface{}) {
	if l.Level >= WarnLevel {
		l.Logger.Print(WarnLevel.String() + fmt.Sprint(v...))
	}
}

// Info works just like log.Print, but with a Green "==> INFO:" prefix.
func (l *Logger) Info(v ...interface{}) {
	if l.Level >= InfoLevel {
		l.Logger.Print(InfoLevel.String() + fmt.Sprint(v...))
	}
}

// Debug works just like log.Print, but with a Cyan "==> DEBUG:" prefix.
func (l *Logger) Debug(v ...interface{}) {
	if l.Level >= DebugLevel {
		l.Logger.Print(DebugLevel.String() + fmt.Sprint(v...))
	}
}

// Panicf works just like log.Panicf, but with a Red "==> PANIC:" prefix.
func (l *Logger) Panicf(format string, v ...interface{}) {
	if l.Level >= PanicLevel {
		l.Logger.Panic(PanicLevel.String() + fmt.Sprintf(format, v...))
	}
}

// Fatalf works just like log.Fatalf, but with a Red "==> FATAL:" prefix.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.Level >= FatalLevel {
		l.Logger.Fatal(FatalLevel.String() + fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.Level >= ErrorLevel {
		l.Logger.Print(ErrorLevel.String() + fmt.Sprintf(format, v...))
	}
}

// Warnf works just like log.Printf, but with a Yellow "==> WARN:" prefix.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.Level >= WarnLevel {
		l.Logger.Print(WarnLevel.String() + fmt.Sprintf(format, v...))
	}
}

// Infof works just like log.Printf, but with a Green "==> INFO:" prefix.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.Level >= InfoLevel {
		l.Logger.Print(InfoLevel.String() + fmt.Sprintf(format, v...))
	}
}

// Debugf works just like log.Printf, but with a Cyan "==> DEBUG:" prefix.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.Level >= DebugLevel {
		l.Logger.Print(DebugLevel.String() + fmt.Sprintf(format, v...))
	}
}

// Panicln works just like log.Panicln, but with a Red "==> PANIC:" prefix.
func (l *Logger) Panicln(v ...interface{}) {
	if l.Level >= PanicLevel {
		l.Logger.Panicln(PanicLevel.String() + fmt.Sprint(v...))
	}
}

// Fatalln works just like log.Fatalln, but with a Red "==> FATAL:" prefix.
func (l *Logger) Fatalln(v ...interface{}) {
	if l.Level >= FatalLevel {
		l.Logger.Fatalln(FatalLevel.String() + fmt.Sprint(v...))
	}
}

func (l *Logger) Errorln(v ...interface{}) {
	if l.Level >= ErrorLevel {
		l.Logger.Println(ErrorLevel.String() + fmt.Sprint(v...))
	}
}

// Warnln works just like log.Println, but with a Yellow "==> WARN:" prefix.
func (l *Logger) Warnln(v ...interface{}) {
	if l.Level >= WarnLevel {
		l.Logger.Println(WarnLevel.String() + fmt.Sprint(v...))
	}
}

// Infoln works just like log.Println, but with a Green "==> INFO:" prefix.
func (l *Logger) Infoln(v ...interface{}) {
	if l.Level >= InfoLevel {
		l.Logger.Println(InfoLevel.String() + fmt.Sprint(v...))
	}
}

// Debugln works just like log.Println, but with a Cyan "==> DEBUG:" prefix.
func (l *Logger) Debugln(v ...interface{}) {
	if l.Level >= DebugLevel {
		l.Logger.Println(DebugLevel.String() + fmt.Sprint(v...))
	}
}
