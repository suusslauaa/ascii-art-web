package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	app.Errors(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	app.Errors(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) badRequest(w http.ResponseWriter) {
	app.clientError(w, http.StatusBadRequest)
}

func (app *application) Errors(w http.ResponseWriter, msgs string, status int) {
	error := Error{
		Message: msgs,
		Code:    status,
	}

	app.errorLog.Printf("%s: %d", msgs, status)

	temp, err := template.ParseFiles("ui/html/error.html")
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = temp.Execute(w, error)
	if err != nil {
		app.serverError(w, err)
	}
}
