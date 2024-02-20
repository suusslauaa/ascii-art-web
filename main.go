package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Input  string
	Output string
}

func GetFont(fontName string) (map[rune][]string, error) {
	file, err := os.Open(fontName)
	if err != nil {
		fmt.Println(err)
		return map[rune][]string{}, err
	}
	count := 0
	r := ' '
	font := make(map[rune][]string, 8)

	scanner := bufio.NewScanner(file)
	for ; scanner.Scan(); count++ {
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

func Print(text []string, font map[rune][]string) string {
	result := ""
	output := make([][8]string, len(text))
	for k, val := range text {
		if val == "" {
			output[k] = [8]string{"\n"}
			continue
		}
		for _, r := range val {
			for i, g := range font[r] {
				output[k][i] += g
			}
		}
	}

	for _, h := range output {
		for _, k := range h {
			if k == "\n" {
				result += "\n"
				break
			}
			result += k + "\n"
		}
	}

	return result
}

func NewLineBreaker(text string) ([]string, error) {
	if len(text) == 0 {
		return []string{}, fmt.Errorf("Empty input")
	}
	words := []string{}
	word := ""
	count := 0
	for _, r := range text {
		if r == '\n' {
			count++
		}
	}
	if count == len(text) {
		for i := 0; i < len(text); i++ {
			fmt.Println()
		}
		return []string{}, nil
	}
	for i, r := range text {
		if r == '\n' {
			words = append(words, word)
			word = ""
			if i == len(text)-1 {
				words = append(words, word)
			}
			continue
		}
		word += string(r)
	}
	if len(word) != 0 {
		words = append(words, word)
	}

	return words, nil
}

func TextToASCIIArt(text, banner string) (string, error) {
	words, err := NewLineBreaker(text)
	art, err := GetFont("banners/standard.txt")
	if err != nil {
		return "", err
	}
	asciiArt := Print(words, art)
	return asciiArt, nil
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "NotFound", http.StatusNotFound)
		return
	}

	data := Data{
		Input:  "",
		Output: "",
	}

	temp, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
}

func output(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("input")
	bannerFilePath := r.FormValue("font")

	output, err := TextToASCIIArt(text, bannerFilePath)
	if err != nil {
		http.Error(w, "ConvertError", 1437)
		return
	}

	data := Data{
		Input:  text,
		Output: output,
	}

	temp, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/text-to-ascii-art", output)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
