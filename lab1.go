package main

import "strconv"

func PolybiusSquare(input string) string {
	l := map[int32]int{
		'a': 11,
		'b': 12,
		'c': 13,
		'd': 14,
		'e': 15,
		'f': 21,
		'g': 22,
		'h': 23,
		'i': 24,
		'j': 24,
		'k': 25,
		'l': 31,
		'm': 32,
		'n': 33,
		'o': 34,
		'p': 35,
		'q': 41,
		'r': 42,
		's': 43,
		't': 44,
		'u': 45,
		'v': 51,
		'w': 52,
		'x': 53,
		'y': 54,
		'z': 55,
	}
	var r string
	for _, char := range input {
		r = r + strconv.Itoa(l[char])
	}
	return r
}

func CaesarCipher(input string, offset int) string {
	if offset > 26 {
		offset %= 26
	}
	var r string
	for _, char := range input {
		pos := int(char) + offset
		if pos > int('z') {
			pos = int('a') + pos - int('z') - 1
		}
		r = r + string(rune(pos))
	}
	return r
}

func Trithemius(input string, key string) string {
	var r string
	i := 0
	for _, char := range input {
		if char > 'z' || char < 'a' {
			r = r + string(char)
			continue
		}
		pos := int(char) + int(key[i%len(key)]) - int('a') + 1
		if pos > int('z') {
			pos = int('a') + pos - int('z') - 1
		}
		r = r + string(rune(pos))
		i++
	}
	return r
}

func Vigenere(input string) string {
	offset := 0

	var r string
	for _, char := range input {
		offset++
		if offset > 26 {
			offset %= 26
		}
		pos := int(char) + offset
		if pos > int('z') {
			pos = int('a') + pos - int('z') - 1
		}
		r = r + string(rune(pos))
	}
	return r
}

func VigenereAdvanced(input string, key string) string {
	offsetIdx := 0

	var r string
	for _, char := range input {
		if char > 'z' || char < 'a' {
			r = r + string(char)
			continue
		}

		offset := int(key[offsetIdx%len(key)]) - int('a')
		pos := int(char) + offset
		if pos > int('z') {
			pos = int('a') + pos - int('z') - 1
		}
		r = r + string(rune(pos))
		offsetIdx++
	}
	return r
}
