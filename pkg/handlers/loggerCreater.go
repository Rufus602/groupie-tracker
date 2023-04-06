package handlers

import (
	"log"
	"net/http"
	"os"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	HTTP     *http.ServeMux
}

func LoggerCreater() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	return infoLog, errorLog
}
