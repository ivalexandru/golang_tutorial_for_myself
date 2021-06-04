package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	msg := os.Args[1]
	length := utf8.RuneCountInString(msg)
	fmt.Println(length)

	//count the nr of chars in the second argument you type when running the program
	//and return that many exclamation signs
	// ex: go run fileName.go abcd
	res := msg + strings.Repeat("!", length)
	res = strings.ToUpper(res)
	fmt.Println(res) // ABCD!!!!
}
