package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputs = "../inputs/day04.txt"

type scratchcard struct {
	victory []int
	score   int
	match   int
}

func (s *scratchcard) member(value int) bool {
	for _, v := range s.victory {
		if v == value {
			return true
		}
	}

	return false
}

func (s *scratchcard) increaseScore() {
	s.match++
	if s.score == 0 {
		s.score = 1
		return
	}

	s.score *= 2
}

func main() {
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	copies := make(map[int]int)
	index := 1
	part1 := 0
	part2 := 0
	scratchcards := make([]scratchcard, 0, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		card := parse(line)
		part1 += card.score
		scratchcards = append(scratchcards, card)
		copies[index]++
		part2++

		for i := index + 1; i < index+1+card.match; i++ {
			_, ok := copies[index]
			if !ok {
				copies[index] = 1
			}

			copies[i] += copies[index]
			part2 += copies[index]
		}
		index++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Part1 : %d\n", part1)
	log.Printf("Part2 : %d\n", part2)
}

func parse(str string) scratchcard {
	parts := strings.Split(str, ":")
	parts = strings.Split(parts[1], "|")
	card := scratchcard{
		victory: make([]int, 0, 0),
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

	return card
}
