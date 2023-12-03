package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var digitsMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

type IndexSet struct {
	indexes  map[int]int
	minIndex int
	maxIndex int
}

func NewIndexSet() IndexSet {
	return IndexSet{indexes: make(map[int]int), minIndex: 9223372036854775807, maxIndex: -1}
}

func (s *IndexSet) insert(i, v int) {
	if i < 0 {
		return
	}

	s.indexes[i] = v

	if i > s.maxIndex {
		s.maxIndex = i
	}

	if i < s.minIndex {
		s.minIndex = i
	}
}

func (s *IndexSet) value() int {
	return s.indexes[s.minIndex]*10 + s.indexes[s.maxIndex]
}

func main() {
	file, err := os.Open("../inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		n := filterNum(line)
		log.Printf("%s - %d\n", line, n)
		result += n
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}

func filterNum(str string) int {
	index := NewIndexSet()
	for d, v := range digitsMap {
		i := strings.Index(str, d)
		index.insert(i, v)
		j := strings.LastIndex(str, d)
		index.insert(j, v)
	}

	return index.value()
}
