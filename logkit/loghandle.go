package logkit

import (
	"fmt"
	"github.com/acornlive/lkit/strkit"
	"io/fs"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

const (
	textNone   = 0
	textWhite  = 37
	textYellow = 33
	textRed    = 31
)

type LogHandler interface {
	Name() string
	Handle(log *Logger)
}

type consoleHandler struct {
	lock sync.Mutex
}

func (c *consoleHandler) Name() string {
	return "console"
}

func (c *consoleHandler) Handle(log *Logger) {
	c.lock.Lock()
	defer c.lock.Unlock()

	color := 0
	switch log.level {
	case ERROR:
		color = textRed
	case WARN:
		color = textYellow
	default:
		color = textNone
	}

	consoleOut(color, formatLog(log))
}

func consoleOut(color int, msg string) {
	fmt.Printf("\x1b[%dm%s\x1b[0m\n", color, msg)
}

func formatLog(log *Logger) string {
	mb := strings.Builder{}
	if strkit.IsNotBlank(log.logMgr.prefix) {
		mb.WriteString(log.logMgr.prefix)
		mb.WriteString(" ")
	}
	mb.WriteString("[" + levelStr(log.level) + "]")
	mb.WriteString(" ")
	mb.WriteString(strkit.FormatTime(log.time))
	mb.WriteString(" ")
	mb.WriteString(log.file)
	mb.WriteString(":" + strconv.Itoa(log.line))
	mb.WriteString(" : ")
	mb.WriteString(log.log)

	if ERROR == log.level {
		for i := 3; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			mb.WriteString("\n")
			mb.WriteString("   " + file + ":" + strconv.Itoa(line))
		}

	}
	return mb.String()

}

type fileHandler struct {
	lock sync.RWMutex
}

func (f *fileHandler) Name() string {
	return "file"
}

func (f *fileHandler) Handle(log *Logger) {
	f.lock.Lock()
	defer f.lock.Unlock()

	logPath := log.logMgr.storePath
	if strkit.IsBlank(logPath) {
		consoleOut(textYellow, "log store path is empty!")
		return
	}

	_, err := os.Stat(logPath)
	if err != nil {
		consoleOut(textYellow, "log store path "+logPath+" is not exist!")
		return
	}

	logFileName := "logkit.log"
	file, err := os.OpenFile(path.Join(logPath, logFileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, fs.ModePerm)
	defer file.Close()
	if err != nil {
		consoleOut(textYellow, "open log store path "+logPath+" error :"+err.Error())
		return
	}

	file.WriteString("\n" + formatLog(log))
}
