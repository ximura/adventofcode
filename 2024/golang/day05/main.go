package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"time"
)

var inputs = "../inputs/day05.txt"

type RuleSet map[int][]int

func (r RuleSet) insert(k, v int) {
	r[k] = append(r[k], v)
}

type update struct {
	data           []int       // raw data, to get middle value
	structuredData map[int]int // value -> position
	correct        bool
}

func (u *update) adjustment(r RuleSet) bool {
	l := len(u.data)
	for i := 0; i < l; i++ {
		v := u.data[i]
		rule := r[v]

		for _, r := range rule {
			pos, ok := u.structuredData[r]
			if ok && pos < i {
				u.data[i] = r
				u.data[pos] = v
				u.structuredData[r] = i
				u.structuredData[v] = pos
				i = 0
				break
			}
		}
	}

	return true
}

func (u *update) isValid(r RuleSet) bool {
	for i, v := range u.data {
		rule := r[v]

		for _, r := range rule {
			pos, ok := u.structuredData[r]
			if ok && pos < i {
				//fmt.Printf("%+v broke rule %d|%d\n", u.data, v, r)
				return false
			}
		}
	}

	return true
}

func value(data []int) int {
	i := int(math.Ceil(float64(len(data)) / 2.0))
	return data[i-1]
}

type input struct {
	r RuleSet
	p []update
}

func part1(in input) int {
	result := 0
	pages := in.p
	for i := range pages {
		isVal := pages[i].isValid(in.r)
		pages[i].correct = isVal
		if isVal {
			v := value(pages[i].data)
			//fmt.Printf("%+v %d\n", page.data, v)
			result += v
		}
		//fmt.Printf("%+v %t\n", page.data, v)
	}

	return result
}

func part2(in input) int {
	result := 0

	pages := in.p
	for i := range pages {
		if pages[i].correct {
			continue
		}
		isVal := pages[i].adjustment(in.r)
		if isVal {
			result += value(pages[i].data)
		}
		//fmt.Printf("%+v %t\n", pages[i].data, isVal)
	}

	return result
}

func main() {
	timeStart := time.Now()
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	in, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(in))
	fmt.Printf("Part 2: %d\n", part2(in))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) (input, error) {
	in := input{}

	scanner := bufio.NewScanner(r)
	in.r = parseRuleSet(scanner)
	in.p = parsePages(scanner)

	if err := scanner.Err(); err != nil {
		return in, err
	}

	return in, nil
}

func parseRuleSet(scanner *bufio.Scanner) RuleSet {
	r := make(RuleSet)

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}

		parts := bytes.Split(line, []byte{'|'})
		r.insert(byteToInt(parts[0]), byteToInt(parts[1]))
	}

	return r
}

func parsePages(scanner *bufio.Scanner) []update {
	r := make([]update, 0, 10)

	for scanner.Scan() {
		line := scanner.Bytes()
		parts := bytes.Split(line, []byte{','})
		page := update{data: make([]int, len(parts)),
			structuredData: make(map[int]int, len(parts))}
		for i := range parts {
			v := byteToInt(parts[i])
			page.data[i] = v
			page.structuredData[v] = i
		}
		r = append(r, page)
	}

	return r
}

func byteToInt(b []byte) int {
	var value int
	for _, b := range b {
		value = value*10 + int(b-48)
	}
	return value
}
