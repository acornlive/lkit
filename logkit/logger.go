package logkit

import (
	"github.com/acornlive/lkit/strkit"
	"runtime"
	"time"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
)

const (
	Console = "console"
	File    = "file"
)

type LogMgr struct {
	level       int
	storePath   string
	handleFlag  int
	handleChain map[string]LogHandler
	prefix      string
}

type Logger struct {
	level  int
	log    string
	file   string
	line   int
	time   time.Time
	logMgr *LogMgr
}

var defaultLogger = &LogMgr{
	level: INFO,
	handleChain: map[string]LogHandler{
		Console: &consoleHandler{},
	},
}

func Log() *LogMgr {
	return defaultLogger
}

func (log *LogMgr) SetLevel(level int) *LogMgr {
	log.level = level
	return log
}

func (log *LogMgr) StorePath(path string) *LogMgr {
	log.storePath = path
	return log
}

func (log *LogMgr) SetPrefix(prefix string) *LogMgr {
	log.prefix = prefix
	return log
}

func (log *LogMgr) Handle(handleName string) *LogMgr {

	var h LogHandler
	switch handleName {
	case File:
		h = &fileHandler{}
	}

	if _, ok := log.handleChain[handleName]; !ok && h != nil {
		log.handleChain[handleName] = h
	}

	return log
}

func (log *LogMgr) AddHandler(handlers ...LogHandler) *LogMgr {
	if handlers != nil && len(handlers) > 0 {
		for _, h := range handlers {
			if _, ok := log.handleChain[h.Name()]; !ok {
				log.handleChain[h.Name()] = h
			}
		}
	}
	return log
}

func levelStr(level int) string {
	switch level {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return ""
	}

}

func Trace(msg string, args ...any) {
	logf(TRACE, msg, args...)
}

func Debug(msg string, args ...any) {
	logf(DEBUG, msg, args...)
}

func Info(msg string, args ...any) {
	logf(INFO, msg, args...)
}

func Warn(msg string, args ...any) {
	logf(WARN, msg, args...)
}

func Error(msg string, args ...any) {
	logf(ERROR, msg, args...)
}

func logf(level int, msg string, args ...any) {

	if level <= defaultLogger.level {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}

		formatMsg := strkit.Format(msg, "{}", args...)

		logIns := &Logger{
			level:  level,
			file:   file,
			line:   line,
			time:   time.Now(),
			log:    formatMsg,
			logMgr: defaultLogger,
		}

		for _, handler := range defaultLogger.handleChain {
			handler.Handle(logIns)
		}

	}
}
