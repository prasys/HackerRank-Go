package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addEntry(book map[string]int, entry string) {
	entries := strings.Split(entry, " ")
	book[entries[0]], _ = strconv.Atoi(entries[1])

}

func checkEntry(book map[string]int, entry string) {
	_, exists := book[entry]
	if exists {
		fmt.Printf("%s=%d \n", entry, book[entry])

	} else {
		fmt.Println("Not found")
	}

}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	addressBook := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	entryNum := true
	entries := 0
	entryCount := 0
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			break
		}

		if entryNum {
			entries, _ = strconv.Atoi(text)
			entryNum = false
			continue
		}
		if entryCount < entries {
			addEntry(addressBook, text)
			entryCount++
		} else {
			checkEntry(addressBook, text)
		}

	}

}
