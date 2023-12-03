package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputs = "../inputs/day02_01.txt"

var restriction = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type GameResult struct {
	number  int
	results map[string]int
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func newGameResult(str string) GameResult {
	parts := strings.Split(str, ":")
	// Define a regular expression to extract the numeric part
	re := regexp.MustCompile(`(\d+)`)

	// Find the match in the input string
	match := re.FindStringSubmatch(parts[0])
	gameNumber, err := strconv.Atoi(match[1])
	if err != nil {
		log.Fatalln(err)
	}

	re = regexp.MustCompile(`(\d+) ([a-zA-Z]+)`)
	colorCounts := make(map[string]int)
	// Extract color and count information using regular expression
	matches := re.FindAllStringSubmatch(parts[1], -1)
	for _, match := range matches {
		count, _ := strconv.Atoi(match[1])
		color := match[2]
		old := colorCounts[color]
		if old < count {
			colorCounts[color] = count
		}
	}

	return GameResult{
		number:  gameNumber,
		results: colorCounts,
	}
}

func (g *GameResult) isPossible() bool {
	for k, r := range restriction {
		v := g.results[k]
		if v > r {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		result += checkGame(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}

func checkGame(str string) int {
	gameResult := newGameResult(str)
	if gameResult.isPossible() {
		return gameResult.number
	}
	return 0
}
