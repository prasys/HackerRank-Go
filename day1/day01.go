package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var newI int
	var newD float64
	var newS string

	// Read and save an integer, double, and String to your variables.
	var counter int = 0 // counter to keep track
	for scanner.Scan() {
		switch counter {
		case 0: // int needs to be summed by i
			newI, _ = strconv.Atoi(scanner.Text())
			i += uint64(newI)
			counter++
		case 1: // double needs to be summed by d
			newD, _ = strconv.ParseFloat(scanner.Text(), 64)
			newD += d
			counter++
		case 2:
			newS = scanner.Text()
			newS = s + newS
			counter++
		default:
			break
		}
	}
	// Print the sum of both integer variables on a new line.

	fmt.Println(i)

	// Print the sum of the double variables on a new line.

	fmt.Println(fmt.Sprintf("%.1f", newD)) // format it to 1decimal place

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

	fmt.Println(newS)

}
