package main

import (
	"flag"
	"groupie_tracker/pkg/handlers"
	"log"
	"net/http"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-страницы")
	flag.Parse()
	infoLog, errorLog:= handlers.LoggerCreater()
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app
	}
	infoLog.Printf("Запуск сервера на %s", *addr)
	err:=srv.ListenandServe()
	errorLog.Fatal(err)
}
