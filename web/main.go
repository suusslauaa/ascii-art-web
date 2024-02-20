package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Input  string
	Output string
}

type Error struct {
	Message string
	Code    int
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/text-to-ascii-art", app.output)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Запуск сервера на %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

func Errors(w http.ResponseWriter, msgs string, status int) {
	error := Error{
		Message: msgs,
		Code:    status,
	}

	temp, err := template.ParseFiles("ui/html/error.html")
	if err != nil {
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, error)
	if err != nil {
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
}
