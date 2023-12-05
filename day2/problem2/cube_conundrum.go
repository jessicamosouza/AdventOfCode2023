package problem1

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cubes = map[string]int{
	"red":   0,
	"green": 0,
	"blue":  0,
}

var gameData = make(map[string][]map[string]int)
var powerSum int

func CubeConundrumTwo() int {
	convertToMap()

	for _, turns := range gameData {
		for _, turn := range turns {
			getNumCubes(turn)
		}
		calculateFinal()

		clearMap()
	}

	return powerSum
}

func calculateFinal() {
	power := cubes["red"] * cubes["blue"] * cubes["green"]
	powerSum += power
}

func clearMap() {
	cubes["red"] = 0
	cubes["blue"] = 0
	cubes["green"] = 0
}

func getNumCubes(turn map[string]int) {
	for color, value := range turn {
		if value > cubes[color] {
			cubes[color] = value
		}
	}
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
