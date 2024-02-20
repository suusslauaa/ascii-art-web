package main

import (
	"html/template"
	"net/http"
	"strings"

	"ascii-art-web/program"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.errorLog.Println("NotFound", http.StatusNotFound)
		Errors(w, "NotFound", http.StatusNotFound)
		return
	}

	data := Data{
		Input:  "",
		Output: "",
	}

	temp, err := template.ParseFiles("ui/html/home.html")
	if err != nil {
		app.errorLog.Println(err.Error())
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		app.errorLog.Println(err.Error())
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
}

func (app *application) output(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("input")
	if text == "" {
		app.errorLog.Println("BadRequest", http.StatusBadRequest)
		Errors(w, "BadRequest", http.StatusBadRequest)
		return
	}

	banner := r.FormValue("font")
	text = strings.ReplaceAll(text, "\r", "")

	output, err := program.TextToASCIIArt(text, banner)
	if err != nil {
		app.errorLog.Println(err.Error())
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	data := Data{
		Input:  text,
		Output: output,
	}

	temp, err := template.ParseFiles("ui/html/home.html")
	if err != nil {
		app.errorLog.Println(err.Error())
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		app.errorLog.Println(err.Error())
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
}
