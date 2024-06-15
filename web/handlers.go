package web

import (
	"html/template"
	"net/http"
	"strings"

	"text-to-ascii-art/program"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	if r.Method != http.MethodGet {
		app.MethodNotAllowed(w)
		return
	}

	temp, err := template.ParseFiles("ui/html/home.html")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = temp.Execute(w, Data{})
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) Output(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("input")
	if text == "" {
		app.BadRequest(w)
		return
	}

	if r.Method != http.MethodPost {
		app.MethodNotAllowed(w)
		return
	}

	banner := r.FormValue("font")
	text = strings.ReplaceAll(text, "\r", "")

	output, err := program.TextToASCIIArt(text, banner)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	temp, err := template.ParseFiles("ui/html/home.html")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = temp.Execute(w, Data{text, output})
	if err != nil {
		app.ServerError(w, err)
	}
}
