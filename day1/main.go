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
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		word := scanner.Bytes()
		value := 0

		for _, c := range word {
			v, err := strconv.Atoi(string(c))
			if err == nil {
				value = v * 10
				break
			}
		}

		for i := len(word) - 1; i >= 0; i-- {
			v, err := strconv.Atoi(string(word[i]))
			if err == nil {
				value += v
				break
			}
		}
		totalCount += value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1")
	fmt.Println(totalCount)

	totalCount = 0
	file2, err := os.Open(_inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file2)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println("string(word)")
		word := scanner.Bytes()
		fmt.Println(string(word))
		word = []byte(massageData(string(word)))
		value := 0

		fmt.Println(string(word))
		for _, c := range word {
			v, err := strconv.Atoi(string(c))
			if err == nil {
				value = v * 10
				break
			}
		}

		for i := len(word) - 1; i >= 0; i-- {
			v, err := strconv.Atoi(string(word[i]))
			if err == nil {
				value += v
				break
			}
		}
		totalCount += value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 2")
	fmt.Println(totalCount)

}

func massageData(s string) string {
	s = strings.ReplaceAll(s, "one", "one1one")
	s = strings.ReplaceAll(s, "two", "two2two")
	s = strings.ReplaceAll(s, "three", "three3three")
	s = strings.ReplaceAll(s, "four", "four4four")
	s = strings.ReplaceAll(s, "five", "five5five")
	s = strings.ReplaceAll(s, "six", "six6six")
	s = strings.ReplaceAll(s, "seven", "seven7seven")
	s = strings.ReplaceAll(s, "eight", "eight8eight")
	s = strings.ReplaceAll(s, "nine", "nine9nine")
	return s
}
