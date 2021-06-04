package main

import (
	"fmt"
	"unicode/utf8"
)

//DO NOT USE LEN TO GET THE LENGTH OF A STRING
//works ok only for english chars, because it counts bytes, not chars
func main() {
	ro := "șalău"        // 7 bytes (fiecare diacr. are cate 2)
	fmt.Println(len(ro)) // 7

	//this counts nr of chars
	fmt.Println(
		utf8.RuneCountInString(ro)) // 5
}
