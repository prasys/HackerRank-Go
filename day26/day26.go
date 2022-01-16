package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	SAMEMONTH = 15
	SAMEYEAR  = 500
	DIFFYEAR  = 10000
)

type book struct {
	return_day   int
	return_month int
	return_year  int
	due_day      int
	due_month    int
	due_year     int
}

func calculateDate(booky *book) int {
	if booky.return_year > booky.due_year {
		return DIFFYEAR
	}
	if booky.return_year <= booky.due_year {
		if booky.return_month <= booky.due_month {
			if booky.return_day <= booky.due_day {
				return 0 // the book is returned on or before the expected return date, no fine will be charged
			} else { // the book is returned after the expected return day but still within the same
				return SAMEMONTH * int(math.Abs(float64(booky.return_day-booky.due_day)))
			}
		} else { //the book is returned after the expected return month but still within the same calendar year
			if (booky.return_year - booky.due_year) < 0 { // handle scenarios whereby if it's a new year and among others
				return 0
			} else {
				return SAMEYEAR * int(math.Abs(float64(booky.return_month-booky.due_month)))
			}
		}
	}
	return -1 // if there are any error
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp := strings.TrimSpace(readLine(reader)) // read for the first time
	returnInfo := strings.Split(nTemp, " ")
	//n := int32(nTemp)
	r_d, _ := strconv.Atoi(returnInfo[0])
	r_m, _ := strconv.Atoi(returnInfo[1])
	r_y, _ := strconv.Atoi(returnInfo[2])
	nTemp = strings.TrimSpace(readLine(reader))
	dueInfo := strings.Split(nTemp, " ")
	d_d, _ := strconv.Atoi(dueInfo[0])
	d_m, _ := strconv.Atoi(dueInfo[1])
	d_y, _ := strconv.Atoi(dueInfo[2])
	booky := &book{return_day: r_d, return_month: r_m, return_year: r_y, due_day: d_d, due_month: d_m, due_year: d_y}
	fmt.Println(calculateDate(booky))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
