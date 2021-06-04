package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var inputValues = []string{}
	// ask for input until empty line
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		inputValues = append(inputValues, line)
		fmt.Println(line, " was added to inputValues")
	}
	fmt.Println(inputValues[0])
}
