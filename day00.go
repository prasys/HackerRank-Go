package main

import (
	"bufio" // using bufffer IO to handle this
	"fmt"
	"os"
)

func main() {
	var input string
	fmt.Println("Hello, World.")
	// https://stackoverflow.com/questions/20895552/how-can-i-read-from-standard-input-in-the-console - Adapted instead of using Scan
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
	}
	fmt.Println(input)
}
