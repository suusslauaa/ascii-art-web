package main

import (
	"html/template"
	"net/http"
	"strings"

	"ascii-art-web/program"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	data := Data{
		Input:  "",
		Output: "",
	}

	temp, err := template.ParseFiles("ui/html/home.html")
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) output(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("input")
	if text == "" {
		app.badRequest(w)
		return
	}

	banner := r.FormValue("font")
	text = strings.ReplaceAll(text, "\r", "")

	output, err := program.TextToASCIIArt(text, banner)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := Data{
		Input:  text,
		Output: output,
	}

	temp, err := template.ParseFiles("ui/html/home.html")
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}
