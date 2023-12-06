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
var seedKey = "seed"

type pair struct {
	source string
	target string
}

var mapper = map[string]pair{
	"seed-to-soil map:":            {source: seedKey, target: "soil"},
	"soil-to-fertilizer map:":      {source: "soil", target: "fertilizer"},
	"fertilizer-to-water map:":     {source: "fertilizer", target: "water"},
	"water-to-light map:":          {source: "water", target: "light"},
	"light-to-temperature map:":    {source: "light", target: "temperature"},
	"temperature-to-humidity map:": {source: "temperature", target: "humidity"},
	"humidity-to-location map:":    {source: "humidity", target: "location"},
}

type seed struct {
	values map[string]int
}

func (s seed) fillWithDefault(src, dst string) {
	_, ok := s.values[dst]
	if !ok {
		s.values[dst] = s.values[src]
	}
}

type seeds map[int]*seed

func (s seeds) fillWithDefault(src, dst string) {
	for i := range s {
		s[i].fillWithDefault(src, dst)
	}
}

func main() {
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := make(seeds)
	key := pair{}
	part1 := 0
	part2 := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		newKey, ok := mapper[line]
		if ok {
			key = newKey
			s.fillWithDefault(key.source, key.target)
			continue
		}
		parse(s, key, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minLocation := 9223372036854775807
	for i := range s {
		l := s[i].values["location"]
		if l < minLocation {
			minLocation = l
			part1 = l
		}
	}

	log.Printf("Part1 : %d\n", part1)
	log.Printf("Part2 : %d\n", part2)
}

func parse(s seeds, p pair, str string) {
	if strings.HasPrefix(str, "seeds:") {
		parseSeeds(s, str)
		return
	}

	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(str, -1)
	if len(matches) < 3 {
		return
	}
	desStart, _ := strconv.Atoi(matches[0][1])
	sourceStart, _ := strconv.Atoi(matches[1][1])
	count, _ := strconv.Atoi(matches[2][1])

	for i := range s {
		v := s[i].values[p.source]
		if v >= sourceStart && v < sourceStart+count {
			s[i].values[p.target] = desStart + (v - sourceStart)
		}
	}
}

func parseSeeds(s seeds, str string) {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		sd, _ := strconv.Atoi(match[1])
		s[sd] = &seed{values: map[string]int{seedKey: sd, "soil": sd}}
	}
}
