package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

//创建两个日志实例分别用于打印 Info Error 日志
var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

//log methods

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

//支持设置日志的层级(InfoLevel,ErrorLevel,Disabled)

const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel 改变日志层级
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}
	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
