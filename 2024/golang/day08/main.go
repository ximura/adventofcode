package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var inputs = "../inputs/day08.txt"

type coord struct {
	x, y int
}

func (c coord) sub(o coord) coord {
	return coord{c.x - o.x, c.y - o.y}
}

func (c coord) add(o coord) coord {
	return coord{c.x + o.x, c.y + o.y}
}

type Map struct {
	antennas     map[byte][]coord
	sizeX, sizeY int
}

func (m Map) inBounds(o coord) bool {
	return o.x >= 0 && o.y >= 0 && o.x < m.sizeX && o.y < m.sizeY
}

func count(input Map) (int, int) {
	antinodes := map[coord]struct{}{}
	antinodes2 := map[coord]struct{}{}

	for _, nodes := range input.antennas {
		for i := range nodes {
			a := nodes[i]
			for j := range nodes {
				if i == j {
					continue
				}

				b := nodes[j]
				delta := a.sub(b)
				f := a.add(delta)

				if input.inBounds(f) {
					antinodes[f] = struct{}{}
				}

				for input.inBounds(f) {
					antinodes2[f] = struct{}{}
					f = f.add(delta)
				}

				delta = b.sub(a)
				f = b.add(delta)
				if input.inBounds(f) {
					antinodes[f] = struct{}{}
				}

				for input.inBounds(f) {
					antinodes2[f] = struct{}{}
					f = f.add(delta)
				}

				antinodes2[a] = struct{}{}
				antinodes2[b] = struct{}{}
			}
		}
	}

	return len(antinodes), len(antinodes2)
}

func part2(input Map) int {
	result := 0

	return result
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

	p1, p2 := count(input)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) (Map, error) {
	m := Map{
		antennas: make(map[byte][]coord),
	}
	row := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		m.sizeX = len(line)

		for col := range line {
			if line[col] == '.' {
				continue
			}

			m.antennas[line[col]] = append(m.antennas[line[col]], coord{x: col, y: row})
		}
		row++
	}

	m.sizeY = row

	if err := scanner.Err(); err != nil {
		return m, err
	}

	return m, nil
}
