package problem2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numbersWords = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Trebuchet() int {
	var sum int

	values, err := getValuesFromInputFile()
	if err != nil {
		return 0
	}

	for _, value := range values {
		num := getNum(value)
		if num < 0 {
			log.Fatal("no  numbers")
		}
		sum += num
	}

	return sum
}

func getNum(value string) int {
	firstChan := make(chan int)
	lastChan := make(chan int)

	go func() {
		firstChan <- getFirtNumber(value)
	}()

	go func() {
		lastChan <- getLastNumber(value)
	}()

	first := <-firstChan
	last := <-lastChan

	str := fmt.Sprintf("%d%d", first, last)
	num, _ := strconv.Atoi(str)

	return num
}

func getLastNumber(value string) int {
	var word string
	for i := len(value) - 1; i >= 0; i-- {
		word = string(value[i]) + word
		for w, _ := range numbersWords {
			if strings.Contains(word, w) {
				return numbersWords[w]
			}
		}

		if unicode.IsDigit(rune(value[i])) {
			word = ""
			asciiToInt, _ := strconv.Atoi(string(value[i]))
			return asciiToInt
		}
	}
	return -1
}

func getFirtNumber(value string) int {
	var word string
	for i := 0; i < len(value); i++ {
		word += string(value[i])
		for w, _ := range numbersWords {
			if strings.Contains(word, w) {
				return numbersWords[w]
			}
		}

		if unicode.IsDigit(rune(value[i])) {
			word = ""
			asciiToInt, _ := strconv.Atoi(string(value[i]))
			return asciiToInt
		}
	}
	return -1
}

func getValuesFromInputFile() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var values []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return values, nil
}
