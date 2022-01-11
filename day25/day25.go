package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkPrimeNumber(number int32) string {
	if int(number) == 1 {
		return "Not prime"
	}

	for i := 2; i*i <= int(number); i++ {
		if int(number)%i == 0 {
			return "Not prime"
		}
	}

	return "Prime"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		fmt.Printf("%s \n", checkPrimeNumber(aItem))
	}

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
