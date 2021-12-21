package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func splitEvenOddString(str []string) {
	var even, odd string
	for i := 1; i < len(str); i++ {
		for j := 0; j < len(str[i]); j++ {
			if j%2 == 0 {
				even += string(str[i][j])
			} else {
				odd += string(str[i][j])
			}
		}
		fmt.Println(even + " " + odd)
		even, odd = "", ""
	}

}

func main() {
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()

		if len(text) != 0 {
			arr = append(arr, text)
		} else {
			break
		}

	}
	splitEvenOddString(arr)
	//checkError(err)
	// multi(n)
}

func readLine(reader *bufio.Reader) string {
	str, err := reader.ReadString('\n')
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(str, "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
