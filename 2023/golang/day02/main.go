package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputs = "../inputs/day02_02.txt"

var restriction = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type GameResult struct {
	number  int
	results map[string]int
}

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

	part1 := 0
	part2 := 0
	scanner := bufio.NewScanner(file)
	games := make([]GameResult, 0, 5)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, newGameResult(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := range games {
		part1 += checkGame(&games[i])
		part2 += fewestNumber(&games[i])
	}

	log.Printf("Part1 : %d\n", part1)
	log.Printf("Part2 : %d\n", part2)
}

func checkGame(g *GameResult) int {
	if g.isPossible() {
		return g.number
	}
	return 0
}

func fewestNumber(g *GameResult) int {
	result := 1
	for _, v := range g.results {
		result = result * v
	}
	return result
}
