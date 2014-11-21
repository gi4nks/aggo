package utils

import (
	"io"
	"log"
	"os"
	"io/ioutil"
)

/* Information type*/
type Logger struct {
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func NewLogger() *Logger {
	return &Logger{}
}

func (logger *Logger) init(traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	logger.Debug = log.New(traceHandle,
		"[DEBUG]: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Info = log.New(infoHandle,
		"[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Warning = log.New(warningHandle,
		"[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Error = log.New(errorHandle,
		"[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
}


func (logger *Logger) InitDebug() {
	logger.init(os.Stdout, os.Stdout, os.Stdout, os.Stdout)
}

func (logger *Logger) InitInfo() {
	logger.init(ioutil.Discard, os.Stdout, os.Stdout, os.Stdout)
}

func (logger *Logger) InitWarn() {
	logger.init(ioutil.Discard, ioutil.Discard, os.Stdout, os.Stdout)
}

func (logger *Logger) InitError() {
	logger.init(ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stdout)
}

/*
func main() {
	//Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	Debug.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
*/
