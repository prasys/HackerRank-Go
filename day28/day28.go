package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func bubbleSort(array *[30]string, size *int) {
	for i := 0; i < *size-1; i++ {
		for j := 0; j < *size-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

}

func sortandPrint(array *[30]string, size *int) {
	bubbleSort(array, size)
	for i := 0; i < *size; i++ {
		fmt.Println(array[i])
	}

}

func orderedEmaillist(firstName string, emailID string) (response string, err bool) {

	//^(?=.{1,20}$)[a-zA-Z]+(?:[a-zA-Z]+)*$ - first name regex
	// ^[a-z](\.?[a-z]){5,}@gmail\.com$
	match, _ := regexp.MatchString("^[a-z](\\.?[a-z]){2,}@gmail\\.com$", emailID)
	if match {
		if len(firstName) <= 20 {
			return firstName, true
		}
	}

	return "", false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)
	var myarr [30]string
	var valid int = 0

	for NItr := 0; NItr < int(N); NItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		firstName := firstMultipleInput[0]

		emailID := firstMultipleInput[1]
		name, ok := orderedEmaillist(firstName, emailID)
		if ok {
			myarr[valid] = name
			valid++
		}

	}
	sortandPrint(&myarr, &valid)

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
