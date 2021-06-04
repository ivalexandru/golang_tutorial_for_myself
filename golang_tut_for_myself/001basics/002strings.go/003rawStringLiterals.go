package main

import "fmt"

func main() {
	//without string literals:
	// path := "c:\\program files\\duper super\\fun.txt\n" +
	// 	"c:\\program files\\really\\funny.png"

	path := `c:\program files\duper super\fun.txt
	c:\program files\really\funny.png`

	fmt.Println(path)
}
