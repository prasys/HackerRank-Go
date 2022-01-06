package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func computeMaxHourGlass(arr [][]int32) int32 {
	var sum int32 = 0
	var max int32 = -999
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			//fmt.Println(sum)
			if sum > max {
				max = sum

			}
		}
	}
	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)

	}
	fmt.Println(computeMaxHourGlass(arr))
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
