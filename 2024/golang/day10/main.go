package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var inputs = "../inputs/day10.txt"

type trail map[path]struct{}

type path struct {
	start, end point
}

type point struct {
	x, y int
}

type input struct {
	field        [][]int
	final        []point
	sizeX, sizeY int
}

func part1(input input) int {
	result := 0
	cache := make(trail)

	for _, f := range input.final {
		acc := make([]point, 0, 10)
		acc = check(input, 9, f.x-1, f.y, acc)
		acc = check(input, 9, f.x+1, f.y, acc)
		acc = check(input, 9, f.x, f.y-1, acc)
		acc = check(input, 9, f.x, f.y+1, acc)

		for _, p := range acc {
			c := path{start: p, end: f}
			if _, ok := cache[c]; !ok {
				result++
				cache[c] = struct{}{}
			}
		}
	}

	return result
}

func check(input input, prev, x, y int, acc []point) []point {
	if x < 0 || y < 0 || x > input.sizeX || y > input.sizeY {
		return acc
	}

	v := input.field[y][x]

	if v == -1 {
		return acc
	}

	if prev-v != 1 {
		return acc
	}

	if v == 0 {
		return append(acc, point{x: x, y: y})
	}

	acc = check(input, v, x-1, y, acc)
	acc = check(input, v, x+1, y, acc)
	acc = check(input, v, x, y-1, acc)
	acc = check(input, v, x, y+1, acc)
	return acc
}

func part2(input input) int {
	checkSum := 0
	return checkSum
}

func main() {
	timeStart := time.Now()
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(input)
	p2 := part2(input)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) (input, error) {
	input := input{
		field: make([][]int, 0, 10),
		final: make([]point, 0, 10),
	}

	row := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		r := make([]int, 0, 10)
		input.sizeX = len(line)

		for col := range line {
			if line[col] == '.' {
				r = append(r, -1)
				continue
			}

			v := byteToInt([]byte{line[col]})

			r = append(r, v)
			if v == 9 {
				input.final = append(input.final, point{x: col, y: row})
			}
		}

		input.field = append(input.field, r)
		row++
	}

	input.sizeX--
	input.sizeY = row - 1

	if err := scanner.Err(); err != nil {
		return input, err
	}

	return input, nil
}

func byteToInt(b []byte) int {
	var value int
	for _, b := range b {
		value = value*10 + int(b-48)
	}
	return value
}
