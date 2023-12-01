package problem1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Trebuchet() int {
	var sum int

	values := getValuesFromInputFile()

	for _, value := range values {
		sum += getNum(value)
	}

	return sum
}

func getNum(value string) int {
	var first, last = -1, 0
	var num2 []int
	
	for _, char := range value {
		if !unicode.IsDigit(char) {
			continue
		}

		num, err := strconv.Atoi(string(char))
		if err != nil {
			log.Println(err)
		}

		num2 = append(num2, num)

		if first == -1 {
			first = num2[0]
			continue
		}

		last = num
	}

	if len(num2) == 1 {
		last = first
	}

	numStr := fmt.Sprintf("%d%d", first, last)
	num, _ := strconv.Atoi(numStr)

	return num
}

func getValuesFromInputFile() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
	}

	var values []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	return values
}
