package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(a []int32, n int) int {
	numOfSwaps := 0
	for i := 0; i < n; i++ {
		var temp int32 = 0
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				temp = a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				numOfSwaps++
			}
		}

	}
	return numOfSwaps
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	totalSwaps := bubbleSort(a, int(n))
	fmt.Printf("Array is sorted in %d swaps. \n", totalSwaps)
	fmt.Printf("First Element: %d \n", a[0])
	fmt.Printf("Last Element: %d \n", a[int(n)-1])
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
