package handlers

import (
	"net/http"
)

func New() *http.ServeMux {
	return http.NewServeMux()
}

func (app *Application) Routes() *http.ServeMux {

	app.HTTP.HandleFunc("/", app.home)
	app.HTTP.HandleFunc("/artist", app.artist)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	app.HTTP.Handle("/static/", http.StripPrefix("/static", fileServer))
	return app.HTTP
}
