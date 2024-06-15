package web

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
)

type ApplicationError struct {
	Message string
	Code    int
}

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	app.Errors(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	app.Errors(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}

func (app *Application) BadRequest(w http.ResponseWriter) {
	app.ClientError(w, http.StatusBadRequest)
}

func (app *Application) MethodNotAllowed(w http.ResponseWriter) {
	app.ClientError(w, http.StatusMethodNotAllowed)
}

func (app *Application) Errors(w http.ResponseWriter, errorMessage string, errorCode int) {
	w.WriteHeader(errorCode)

	temp, err := template.ParseFiles("ui/html/error.html")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	if err := temp.Execute(w, ApplicationError{Message: errorMessage, Code: errorCode}); err != nil {
		http.Error(w, "err", http.StatusInternalServerError)
	}
}
