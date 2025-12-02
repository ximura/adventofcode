package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var inputs = "../inputs/day02.txt"

type Data = []IDRange

type IDRange struct {
	Start int64
	End   int64
}

func main() {
	timeStart := time.Now()
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	p1, p2 := calcualte(d)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) (Data, error) {
	d := Data{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		ranges := bytes.Split(line, []byte(","))
		for i := range ranges {
			if len(ranges[i]) == 0 {
				continue
			}
			parts := bytes.Split(ranges[i], []byte("-"))

			start, _ := strconv.ParseInt(string(parts[0]), 10, 64)
			end, _ := strconv.ParseInt(string(parts[1]), 10, 64)

			d = append(d, IDRange{
				Start: start,
				End:   end,
			})
		}
	}

	return d, nil
}

func calcualte(d Data) (int64, int64) {
	part1 := int64(0)
	part2 := int64(0)

	for _, r := range d {
		for i := r.Start; i <= r.End; i++ {
			if isInvalidRule1(i) {
				part1 += i
			}
			if isInvalidRule2(i) {
				part2 += i
			}
		}
	}

	return part1, part2
}

// returns true if number is of the form XX where X repeated twice
func isInvalidRule1(n int64) bool {
	s := strconv.FormatInt(n, 10)
	l := len(s)

	// must have even number of digits
	if l%2 != 0 {
		return false
	}

	half := l / 2
	first := s[:half]
	second := s[half:]

	return first == second
}

// returns true if string s consists of a repeating substring repeated >= 2 times
func isInvalidRule2(n int64) bool {
	s := strconv.FormatInt(n, 10)
	length := len(s)

	// Try every possible substring length from 1 to half of the number length
	for subLen := 1; subLen*2 <= length; subLen++ {

		// Only lengths that evenly divide total length
		if length%subLen != 0 {
			continue
		}

		sub := s[:subLen]
		repeats := length / subLen

		// Reconstruct pattern
		ok := true
		for i := 1; i < repeats; i++ {
			if s[i*subLen:(i+1)*subLen] != sub {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}
