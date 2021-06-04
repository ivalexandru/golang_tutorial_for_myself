package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var i int = 3
	var d float64 = 3.3
	var s string = "ana"

	scanner := bufio.NewScanner(os.Stdin)
	inputvals := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}
		inputvals = append(inputvals, line)
	}

	fmt.Println(inputvals)

	iInput, _ := strconv.Atoi(inputvals[0])
	fmt.Println(i + iInput)

	dInput, _ := strconv.ParseFloat(inputvals[1], 64)
	fmt.Println(d + dInput)

	fmt.Println(s + inputvals[2])
}
