package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var inputs = "../inputs/day04.txt"

type pair struct {
	x, y int
}

var xmas = []byte{'X', 'M', 'A', 'S'}

func part1(lines []string) int {
	result := 0

	direction := []pair{pair{x: -1, y: -1}, pair{x: -1, y: 0}, pair{x: -1, y: 1},
		pair{x: 0, y: -1}, pair{x: 0, y: 1},
		pair{x: 1, y: -1}, pair{x: 1, y: 0}, pair{x: 1, y: 1}}

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] != 'X' {
				continue
			}

			// we found start coordinates
			for _, d := range direction {
				if findNextMatchByDirection(lines, i, j, 1, d) {
					result++
				}
			}
		}
	}

	return result
}

func findNextMatchByDirection(lines []string, i, j, k int, direction pair) bool {
	if k == 4 {
		return true
	}

	columntCount := len(lines)
	rowCount := len(lines[0])
	c := xmas[k]

	newX := i + direction.x
	newY := j + direction.y

	if newX < 0 || newX >= columntCount {
		return false
	}

	if newY < 0 || newY >= rowCount {
		return false
	}

	if lines[newX][newY] == c {
		return findNextMatchByDirection(lines, newX, newY, k+1, direction)
	}

	return false
}

func part2(lines []string) int {
	result := 0

	columntCount := len(lines)
	rowCount := len(lines[0])
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] != 'A' {
				continue
			}

			if i-1 < 0 || j-1 < 0 ||
				i+1 >= columntCount || j+1 >= rowCount {
				continue
			}

			// we found start coordinates
			if ((lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S') ||
				(lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M')) &&
				((lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S') ||
					(lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M')) {
				result++
			}
		}
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

	lines, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) ([]string, error) {
	lines := make([]string, 0, 10)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, nil
}
