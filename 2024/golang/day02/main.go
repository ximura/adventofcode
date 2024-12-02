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

type DirectionType int8

const (
	DirectionNone DirectionType = 0
	DirectionUp   DirectionType = 1
	DirectionDown DirectionType = 2
)

func NewReport() Report {
	return Report{
		data:      make([]int, 0, 10),
		direction: DirectionNone,
		safe:      true,
	}
}

type Report struct {
	data      []int
	direction DirectionType
	safe      bool
}

func (r *Report) insert(v, index int) {
	if index == -1 {
		r.data = append(r.data, v)
		return
	}

	dif := v - r.data[index]
	newDirection := DirectionNone
	if dif > 0 {
		newDirection = DirectionUp
	}

	if dif < 0 {
		newDirection = DirectionDown
	}
	if r.direction == DirectionNone {
		r.direction = newDirection
	}

	r.safe = r.safe &&
		newDirection != DirectionNone &&
		r.direction == newDirection &&
		abs(dif) < 4
	r.data = append(r.data, v)
}

type Levels struct {
	reports []Report
	count   int
}

func main() {
	timeStart := time.Now()
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lvl, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", lvl.count)
	fmt.Printf("Part 2: %d\n", 0)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) (Levels, error) {
	lvl := Levels{
		reports: make([]Report, 0, 10),
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		p := bytes.Split(line, []byte(" "))
		l := len(p)
		if l < 5 {
			return lvl, fmt.Errorf("unsupported line format %s", string(line))
		}

		r := NewReport()
		for i, b := range p {
			data, _ := strconv.Atoi(string(b))
			r.insert(data, i-1)
		}

		lvl.reports = append(lvl.reports, r)
		if r.safe {
			lvl.count++
		}
	}

	if err := scanner.Err(); err != nil {
		return lvl, err
	}

	return lvl, nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
