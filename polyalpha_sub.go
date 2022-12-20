package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode/utf16"
)

const PATH = "book.txt"
const ALPHABET_SIZE = 30

var nonAlphanumericRegex = regexp.MustCompile(`[^а-яА-Я]+`)

func PolyalphabeticSubstitution(text string) string {
	let, letMap := GenerateAlphabet()
	book := ReadBook(PATH)
	var result string
	for i, c := range strings.ToUpper(text) {
		wordLetNum := letMap[string(c)]
		bookLetNum := letMap[string(book[i/2])]

		letNum := (wordLetNum + bookLetNum) % ALPHABET_SIZE
		if letNum < 1 {
			letNum = 1
		}
		result += let[letNum-1]
		fmt.Printf("Letter '%c' +  Book '%s' = C '%s' \n", c, string(book[i/2]), result)
	}

	return result
}

func ReadBook(path string) []uint16 {
	content, err := os.ReadFile(PATH)
	if err != nil {
		log.Fatal(err)
	}
	book := nonAlphanumericRegex.ReplaceAllString(string(content), "")
	book = strings.ToUpper(book)
	return utf16.Encode([]rune(book))
}

func GenerateAlphabet() ([]string, map[string]int) {
	var l []string
	for i := 'А'; i <= 'Я'; i++ {
		if i == 'Ы' || i == 'Э' {
			continue
		}
		l = append(l, string(i))
	}
	letMap := make(map[string]int, 30)
	for i := 0; i < len(l); i++ {
		letMap[l[i]] = i + 1
	}

	return l, letMap
}
