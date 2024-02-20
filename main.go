package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Data struct {
	Input  string
	Output string
}

type Error struct {
	Message string
	Code    int
}

func Split(text string) (words []string) {
	count := 0

	for _, r := range text {
		if r == '\n' {
			count++
		}
	}

	if count == len(text) {
		return []string{}
	}

	word := ""

	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			words = append(words, word)
			word = ""

			if i == len(word)-1 {
				words = append(words, word)
			}

			continue
		}

		word += string(text[i])
	}

	if len(word) != 0 {
		words = append(words, word)
	}

	return
}

func GetFont(banner string) (map[rune][]string, error) {
	file, err := os.Open(banner)
	if err != nil {
		return map[rune][]string{}, err
	}

	font, count, r := make(map[rune][]string, 8), 0, ' '

	for scanner := bufio.NewScanner(file); scanner.Scan(); count++ {
		text := scanner.Text()
		if count != 0 && count != 9 {
			font[r] = append(font[r], text)
		} else if count == 9 {
			count = 0
			r++
		}
	}

	return font, nil
}

func TextToASCIIArt(text, banner string) (output string, err error) {
	font, err := GetFont(banner)
	if err != nil {
		return
	}

	words, art := Split(text), [][8]string{}

	for j, word := range words {
		art = append(art, [8]string{})

		if word == "" {
			art[j] = [8]string{"\n", ""}
			continue
		}

		for _, letter := range word {
			for i, line := range font[letter] {
				art[j][i] += line
			}
		}
	}

	for i := 0; i < len(art); i++ {
		for j := 0; j < len(art[i]); j++ {
			if art[i][j] == "\n" {
				output += "\n"
				break
			}

			output += art[i][j] + "\n"
		}
	}

	return
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, "NotFound", http.StatusNotFound)
		return
	}

	data := Data{
		Input:  "",
		Output: "",
	}

	temp, err := template.ParseFiles("home.html")
	if err != nil {
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
}

func output(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("input")
	if text == "" {
		Errors(w, "BadRequest", http.StatusBadRequest)
		return
	}

	banner := r.FormValue("font")
	text = strings.ReplaceAll(text, "\r", "")

	output, err := TextToASCIIArt(text, banner)
	if err != nil {
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	data := Data{
		Input:  text,
		Output: output,
	}

	temp, err := template.ParseFiles("home.html")
	if err != nil {
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		Errors(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/text-to-ascii-art", output)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Errors(w http.ResponseWriter, msgs string, status int) {
	error := Error{
		Message: msgs,
		Code:    status,
	}

	temp, err := template.ParseFiles("error.html")
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
