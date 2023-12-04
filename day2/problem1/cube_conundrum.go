package problem1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var gameData = make(map[string][]map[string]int)

func CubeConundrum() int {
	convertToMap()

	var sum int
	for gameNumber, turns := range gameData {
		if isValidGame(turns) {
			if gameNum, err := extractGameNumber(gameNumber); err == nil {
				sum += gameNum
			}
		}
	}
	return sum
}

func isValidGame(turns []map[string]int) bool {
	for _, turn := range turns {
		if !isValidTurn(turn) {
			return false
		}
	}
	return true
}

func isValidTurn(turn map[string]int) bool {
	for color, count := range turn {
		if count > cubes[color] {
			return false
		}
	}
	return true
}

func extractGameNumber(gameNumber string) (int, error) {
	parts := strings.Split(gameNumber, " ")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid game number format")
	}
	return strconv.Atoi(parts[1])
}

func convertToMap() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	colorPattern := regexp.MustCompile(`(\d+)\s(\w+)`)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		gameNumber, colorData := parts[0], parts[1]

		var segmentsData []map[string]int
		segments := strings.Split(colorData, ";")

		for _, segment := range segments {
			segmentData := map[string]int{"blue": 0, "green": 0, "red": 0}
			matches := colorPattern.FindAllStringSubmatch(strings.TrimSpace(segment), -1)

			for _, match := range matches {
				count, _ := strconv.Atoi(match[1])
				color := match[2]
				segmentData[color] = count
			}
			segmentsData = append(segmentsData, segmentData)
		}
		gameData[gameNumber] = segmentsData
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
