package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// region word
	fmt.Println("Please enter a word.")
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); ok {
		input += strings.ToLower(scanner.Text())
	}
	// endregion

	// region int key
	//fmt.Println("Please enter key length.")
	//var n int
	//_, err := fmt.Scanf("%d", &n)
	//if err != nil {
	//	return
	//}
	// endregion

	// region string key
	//fmt.Println("Please enter key.")
	//var key string
	//if ok := scanner.Scan(); ok {
	//	key += scanner.Text()
	//}
	// endregion

	//r := PolybiusSquare(input)
	//r := CaesarCipher(input, n)
	//r := Trithemius(input, key)
	//r := Vigenere(input)
	//r := VigenereAdvanced(input, key)
	r := PolyalphabeticSubstitutionDecrypt(input)
	fmt.Println("Decrypted: %s", r)
	c := PolyalphabeticSubstitution(r)
	fmt.Println("Encrypted: %s", c)

}
