package handlers

import (
	"groupie_tracker/pkg/structure"
	"net/http"
	"strconv"
)

var API []structure.Artist

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
	}
	_, err := structure.JsonReader()
	if err != nil {
		app.serverError(w, err)
	}
	//s, err:=
}

func (app *application) artist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || (id < 1 && id > 52) {
		app.notFound(w)
		return
	}

}
