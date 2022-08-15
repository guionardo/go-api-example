package infra

import (
	"io"
	"log"
	"os"
)

var (
	logFile *os.File
)

func SetupLog() {
	var err error
	logFile, err = os.OpenFile("api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
