package problem1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Trebuchet() (int, error) {
	var sum int

	values, err := getValuesFromInputFile()
	if err != nil {
		return 0, err
	}

	for _, value := range values {
		num, _ := getNum(value)
		sum += num
	}

	return sum, nil
}

func getNum(value string) (int, error) {
	var digits []int
	for _, char := range value {
		if unicode.IsDigit(char) {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return 0, err
			}
			digits = append(digits, num)
		}
	}
	if len(digits) == 0 {
		return 0, fmt.Errorf("no digits found in string")
	}
	if len(digits) == 1 {
		digits = append(digits, digits[0])
	}

	numStr := fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1])
	return strconv.Atoi(numStr)
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
