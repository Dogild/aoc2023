package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const _inputFileName = "input.txt"

func main() {

	file, err := os.Open(_inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalCount := 0
	scanner := bufio.NewScanner(file)
	// red := 12
	// green := 13
	// blue := 14
	index := 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		index++
		var word string
		word = string(scanner.Bytes())
		word = word[strings.Index(word, ":")+2:]
		words := strings.Split(word, ";")
		tred := 0
		tgreen := 0
		tblue := 0

		fmt.Println(word)

		for _, w := range words {
			values := strings.Split(w, ",")

			for _, v := range values {

				v = strings.TrimSpace(v)
				if strings.Contains(v, "red") {
					t, _ := strconv.Atoi(v[:strings.Index(v, " ")])
					if tred < t {
						tred = t
					}
				}

				if strings.Contains(v, "green") {
					t, _ := strconv.Atoi(v[:strings.Index(v, " ")])
					if tgreen < t {
						tgreen = t
					}
				}

				if strings.Contains(v, "blue") {
					t, _ := strconv.Atoi(v[:strings.Index(v, " ")])
					if tblue < t {
						tblue = t
					}
				}
			}
		}

		totalCount += tblue * tred * tgreen

	}

	fmt.Println(totalCount)
}
