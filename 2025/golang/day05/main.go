package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Part1
// simple in range search (inAnyRange): 0.59ms
// binary search: 0.13ms
// merge + simple search: 0.25ms
// merge + binary search: 0.12ms

var inputs = "../inputs/day05.txt"

type Data struct {
	ranges      []Interval
	ingridients []int
}

type Interval struct {
	lo, hi int
}

func parseRange(line string) (Interval, error) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return Interval{}, fmt.Errorf("invalid range: %s", line)
	}

	lo, err := strconv.Atoi(parts[0])
	if err != nil {
		return Interval{}, err
	}
	hi, err := strconv.Atoi(parts[1])
	if err != nil {
		return Interval{}, err
	}

	return Interval{lo, hi}, nil
}

func main() {
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	timeStart := time.Now()
	p1, p2 := calculate(d)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) (Data, error) {
	d := Data{
		ranges:      make([]Interval, 0, 100),
		ingridients: make([]int, 0, 100),
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		rang, err := parseRange(line)
		if err != nil {
			ing, _ := strconv.Atoi(line)
			d.ingridients = append(d.ingridients, ing)
		} else {
			d.ranges = append(d.ranges, rang)
		}

	}

	return d, nil
}

func calculate(grid Data) (int, int) {
	part1 := 0
	part2 := 0

	sort.Slice(grid.ranges, func(i, j int) bool {
		return grid.ranges[i].lo < grid.ranges[j].lo
	})

	grid.ranges = mergeRanges(grid.ranges)

	for _, ing := range grid.ingridients {
		if inRangeBinary(ing, grid.ranges) {
			part1++
		}
	}

	for _, r := range grid.ranges {
		part2 += r.hi - r.lo + 1
	}

	return part1, part2
}

// for part 1 gives 0.59ms
func inAnyRange(id int, ranges []Interval) bool {
	for _, r := range ranges {
		if id >= r.lo && id <= r.hi {
			return true
		}
	}
	return false
}

// inRangeBinary checks using binary search over intervals.
func inRangeBinary(id int, ranges []Interval) bool {
	// Binary search for the rightmost range with start <= id
	i := sort.Search(len(ranges), func(i int) bool {
		return ranges[i].lo > id
	})
	// i is index of first range where lo > id
	// so the candidate interval is i-1
	if i == 0 {
		return false // all ranges start > id
	}
	r := ranges[i-1]
	return id >= r.lo && id <= r.hi
}

// mergeRanges merges all overlapping or contiguous intervals.
func mergeRanges(ranges []Interval) []Interval {
	if len(ranges) == 0 {
		return nil
	}

	merged := []Interval{ranges[0]}

	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]

		// Overlap or touching (hi >= next.lo - 1)
		if r.lo <= last.hi+1 {
			if r.hi > last.hi {
				last.hi = r.hi
			}
		} else {
			// Disjoint interval
			merged = append(merged, r)
		}
	}

	return merged
}
