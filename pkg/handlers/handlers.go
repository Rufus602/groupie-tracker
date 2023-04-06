package handlers

import (
	"fmt"
	"groupie_tracker/pkg/structure"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//var Api func() []structure.Artist

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
	}
	Information, err := structure.JsonReader()
	if err != nil {
		app.serverError(w, err)
	}
	temp, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {

		return
	}
	err = temp.ExecuteTemplate(w, "index", Information)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (app *Application) artist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/artist/" {
		fmt.Println("no here")

		app.notFound(w)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || (id < 1 && id > 52) {

		fmt.Println("here")
		app.notFound(w)
		return
	}
	//Information := Api()
	temp, err := template.ParseFiles("./ui/html/artPage.html")
	if err != nil {
		app.serverError(w, err)
		return
	}
	Information, err := structure.JsonReader()
	if err != nil {
		app.serverError(w, err)
	}
	err = temp.ExecuteTemplate(w, "artPage", Information[id-1])
	if err != nil {
		app.serverError(w, err)
		return
	}
}

//
//func Jsongiver(app *Application, w http.ResponseWriter) func() []structure.Artist {
//	Information, err := structure.JsonReader()
//	if err != nil {
//		app.serverError(w, err)
//	}
//	return func() []structure.Artist {
//		return Information
//	}
//}
