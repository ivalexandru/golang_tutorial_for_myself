package main

import (
	"fmt"
	"strings"
)

//trim the white lines
func main() {
	msg := `

	The weather looks good.
I should go and play.
	`
	//removes all extra whitespaces
	fmt.Println(strings.Trim(msg, "\n,\t"))

	//v2:
	fmt.Println(strings.TrimSpace(msg))

	//removes whitespace at the end
	fmt.Println(strings.TrimRight(msg, " "))
}
