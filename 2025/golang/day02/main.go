package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
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
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Using %d CPU threads\n", numCPU)

	var wg sync.WaitGroup
	type result struct {
		rule1 int64
		rule2 int64
	}
	resultCh := make(chan result, len(d))

	for _, r := range d {
		wg.Add(1)

		go func(r IDRange) {
			defer wg.Done()

			buf := make([]byte, 0, 32)

			localRule1 := int64(0)
			localRule2 := int64(0)

			for i := r.Start; i <= r.End; i++ {
				buf = strconv.AppendInt(buf[:0], i, 10)
				s := string(buf)

				if isInvalidRule1Fast(s) {
					localRule1 += i
					localRule2 += i // all Rule1 are also Rule2
					continue
				}

				if isInvalidRule2Fast(s) {
					localRule2 += i
				}
			}

			resultCh <- result{rule1: localRule1, rule2: localRule2}
		}(r)
	}

	wg.Wait()
	close(resultCh)

	part1 := int64(0)
	part2 := int64(0)

	for res := range resultCh {
		part1 += res.rule1
		part2 += res.rule2
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

// RULE 1 — exactly 2 repetitions
func isInvalidRule1Fast(s string) bool {
	l := len(s)
	if l%2 != 0 {
		return false
	}
	h := l / 2
	// fast check: compare first with second
	return s[:h] == s[h:]
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

// RULE 2 — repeated >= 2 times
func isInvalidRule2Fast(s string) bool {
	length := len(s)

	// Try every divisor of the length
	for subLen := 1; subLen*2 <= length; subLen++ {
		if length%subLen != 0 {
			continue
		}

		repeats := length / subLen
		firstSub := s[:subLen]
		ok := true

		// compare blocks
		for i := 1; i < repeats; i++ {
			if s[i*subLen:(i+1)*subLen] != firstSub {
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
