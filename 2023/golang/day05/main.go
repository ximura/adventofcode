package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputs = "../inputs/day05.txt"

type seed struct {
	values map[string]int
}

func main() {
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seeds := make(map[int]seed)

	part1 := 0
	part2 := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		parse(seeds, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Part1 : %d\n", part1)
	log.Printf("Part2 : %d\n", part2)
}

func parse(s map[int]seed, str string) {
	if strings.HasPrefix(str, "seeds:") {
		parseSeeds(s, str)
		return
	}

	m := ""
	if strings.HasPrefix(str, "seed-to-soil map:") {
		m = "soil"
	}
	if strings.HasPrefix(str, "soil-to-fertilizer map:") {
		m = "fertilizer"
	}

	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(parts[0], -1)
	for _, match := range matches {
		score, _ := strconv.Atoi(match[1])
		card.victory = append(card.victory, score)
	}

	matches = re.FindAllStringSubmatch(parts[1], -1)
	for _, match := range matches {
		score, _ := strconv.Atoi(match[1])
		if card.member(score) {
			card.increaseScore()
		}
	}
}

func parseSeeds(s map[int]seed, str string) {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		sd, _ := strconv.Atoi(match[1])
		s[sd] = seed{}
	}
}
