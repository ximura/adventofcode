package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var inputs = "../inputs/day07.txt"

type Input struct {
	Width, Height int
	Start         Point
	Passed        []Point
	Splitter      []Point
}

type Point struct {
	X, Y int
}

func main() {
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d := parseFile(file)

	timeStart := time.Now()
	p1, p2 := calculate(d)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) *Input {
	p := &Input{
		Start:    Point{-1, -1},
		Splitter: []Point{},
	}

	y := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r")
		if len(line) > p.Width {
			p.Width = len(line)
		}

		for x, ch := range line {
			pos := Point{x, y}
			switch ch {
			case '^':
				p.Splitter = append(p.Splitter, pos)
			case 'S':
				p.Start = pos
			}
		}
		y++
	}
	p.Height = y

	return p
}

func calculate(data *Input) (int, int) {
	part1 := 0
	part2 := 0

	// ----- part 2 -------------
	// beamCounts: map column -> count of beams at that column in current row
	beamCounts := map[Point]int{data.Start: 1}
	for range data.Height {
		next := map[Point]int{}
		for point, cnt := range beamCounts {
			point.Y++
			if contains(data.Splitter, point) {
				part1++
				r := Point{point.X + 1, point.Y}
				l := Point{point.X - 1, point.Y}
				next[r] += cnt
				next[l] += cnt
			} else {
				next[point] += cnt
			}
		}
		beamCounts = next
	}
	for _, cnt := range beamCounts {
		part2 += cnt
	}

	return part1, part2
}

func contains(gs []Point, p Point) bool {
	for _, g := range gs {
		if g == p {
			return true
		}
	}
	return false
}

func RenderParsed(p *Input) {
	for y := 0; y < p.Height; y++ {
		for x := 0; x < p.Width; x++ {
			pos := Point{x, y}

			switch {
			case pos == p.Start:
				fmt.Print("S")
			case contains(p.Splitter, pos):
				fmt.Print("^")
			case contains(p.Passed, pos):
				fmt.Print("|")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
