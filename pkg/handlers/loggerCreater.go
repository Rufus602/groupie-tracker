package handlers

import (
	"log"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func LoggerCreater() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	return infoLog, errorLog
}
