package main

import (
	"flag"
	"groupie_tracker/pkg/handlers"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-страницы")
	flag.Parse()
	infoLog, errorLog := handlers.LoggerCreater()
	app := &handlers.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		HTTP:     handlers.New(),
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}
	infoLog.Printf("Запуск сервера на %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
