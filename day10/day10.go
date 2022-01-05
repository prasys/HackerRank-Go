package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)
	fmt.Printf("%d \n", numOfSetBits(n))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func numOfSetBits(n int) int {
	count, max, prev := 0, 0, 1
	for n != 0 {
		if n&1 == 1 {
			prev = 1
		} else {
			count = 0
			prev = 0
		}
		if prev == 1 && n&1 == 1 {
			count += n & 1
			if count > max {
				max = count
			}
		}

		n >>= 1
	}
	return max

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
