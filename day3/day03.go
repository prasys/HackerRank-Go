package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func printNumber(num int32) string {

	if num%2 != 0 { // n is odd
		return ("Weird")
	}
	if num%2 == 0 && (num > 1 && num < 6) { // n is even from 2 to 5 (inclusive)
		return ("Not Weird")
	}
	if num%2 == 0 && (num > 5 && num < 21) { // n is even and from 6 to 20 (inclusive)
		return ("Weird")
	}
	if num%2 == 0 && (num > 21) { //n is even and greater than 20
		return ("Not Weird")
	}

	return ""

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)
	fmt.Println(printNumber(N))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
