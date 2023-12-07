package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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

type rule struct {
	sourceStart int
	destStart   int
	count       int
}

type ruleSet struct {
	source string
	target string

	r []rule
}

type seed map[string]int

func (s seed) applyRule(rs ruleSet) {
	v := s[rs.source]
	for _, r := range rs.r {
		if v >= r.sourceStart && v < r.sourceStart+r.count {
			s[rs.target] = r.destStart + (v - r.sourceStart)
			break
		}
	}

	_, ok := s[rs.target]
	if !ok {
		s[rs.target] = s[rs.source]
	}
}

func (s seed) reverse(rs ruleSet) {
	v := s[rs.target]
	for _, r := range rs.r {
		source := v + r.sourceStart - r.destStart
		if source >= r.sourceStart && source < r.sourceStart+r.count {
			s[rs.source] = source
			return
		}
	}

	s[rs.source] = v
}

func main() {
	timeStart := time.Now()
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules := make([]ruleSet, 0)
	seeds := make([]int, 0)
	rs := ruleSet{}
	key := pair{}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seeds = parseSeeds(line)
			continue
		}

		newKey, ok := mapper[line]
		if ok {
			if len(rs.r) > 0 {
				rules = append(rules, rs)
			}
			key = newKey
			rs = ruleSet{source: key.source, target: key.target}
			continue
		}
		rule := parse(line)
		rs.r = append(rs.r, rule)
	}
	rules = append(rules, rs)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	p1 := part1(seeds, rules)
	p2 := part2(rules, seeds)

	log.Printf("Part1 : %d\n", p1)
	log.Printf("Part2 : %d\n", p2)
	log.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func part1(seeds []int, rules []ruleSet) int {
	min := 9223372036854775807
	for _, s := range seeds {
		l := applyRules(seed{seedKey: s, "soil": s}, rules)
		if l < min {
			min = l
		}
	}
	return min
}

func part2(rules []ruleSet, seeds []int) int {
	l := 0
	for {
		s := seed{seedKey: -1, "location": l}
		for i := len(rules); i > 0; i-- {
			rs := rules[i-1]
			s.reverse(rs)
		}

		seedValue := s[seedKey]
		for i := 0; i < len(seeds); i = i + 2 {
			start := seeds[i]
			count := seeds[i+1]
			if seedValue > start && seedValue < start+count {
				return l
			}
		}

		l++
	}
}

func parse(str string) rule {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(str, -1)

	desStart, _ := strconv.Atoi(matches[0][1])
	sourceStart, _ := strconv.Atoi(matches[1][1])
	count, _ := strconv.Atoi(matches[2][1])

	return rule{sourceStart: sourceStart, destStart: desStart, count: count}
}

func parseSeeds(str string) []int {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(str, -1)
	result := make([]int, len(matches))
	for i, match := range matches {
		sd, _ := strconv.Atoi(match[1])
		result[i] = sd
	}
	return result
}

func applyRules(s seed, rs []ruleSet) int {
	for _, r := range rs {
		s.applyRule(r)
	}

	return s["location"]
}
