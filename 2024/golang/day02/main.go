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
		unsafePos: -1,
		safe:      true,
	}
}

type Report struct {
	data      []int
	direction DirectionType
	unsafePos int
	safe      bool
}

func (r *Report) insert(v int) {
	l := len(r.data)
	if l == 0 {
		r.data = append(r.data, v)
		return
	}

	dif := v - r.data[l-1]
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

	if !r.safe && r.unsafePos == -1 {
		r.unsafePos = l - 1
	}

	r.data = append(r.data, v)
}

func (r *Report) isSafe(data []int, skip int) bool {
	l := len(data) - 1

	d := make([]int, 0, l)
	d = append(d, data[:skip]...)
	d = append(d, data[skip+1:]...)
	direction := DirectionNone
	i := 1
	for {
		dif := d[i] - d[i-1]
		newDirection := DirectionNone
		if dif > 0 {
			newDirection = DirectionUp
		}

		if dif < 0 {
			newDirection = DirectionDown
		}
		if direction == DirectionNone {
			direction = newDirection
		}

		safe := newDirection != DirectionNone &&
			direction == newDirection &&
			abs(dif) < 4

		if !safe {
			return false
		}

		i++
		if i >= l {
			break
		}
	}

	return true
}

func (r *Report) applyDamper() bool {
	return r.safe ||
		r.isSafe(r.data, 0) ||
		r.isSafe(r.data, r.unsafePos) ||
		r.isSafe(r.data, r.unsafePos+1) ||
		r.isSafe(r.data, r.unsafePos+2)
}

type Levels struct {
	reports []Report
	count   int
}

func (l *Levels) applyDamper() int {
	count := 0
	for _, r := range l.reports {
		if r.applyDamper() {
			count++
		}
	}

	return count
}

func (l *Levels) print() {
	for _, r := range l.reports {
		fmt.Printf("%+v\n", r)
	}
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
	fmt.Printf("Part 2: %d\n", lvl.applyDamper())
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
		for _, b := range p {
			data, _ := strconv.Atoi(string(b))
			r.insert(data)
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
