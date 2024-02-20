package program

import (
	"bufio"
	"os"
)

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
