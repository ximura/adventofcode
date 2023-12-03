package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var inputs = "../inputs/day03_1.txt"

type symbol struct {
	x_start int
	x_end   int
	text    string
	value   int
}

type field struct {
	values [][]symbol
	text   []string
	x      int
	y      int
}

func (f *field) hasAdjacent(x, start, end int) bool {
	if f.getSymbol(x, start-1) || f.getSymbol(x, end+1) {
		return true
	}

	for i := start - 1; i <= end+1; i++ {
		if f.getSymbol(x+1, i) || f.getSymbol(x-1, i) {
			return true
		}
	}

	return false
}

func (f *field) getSymbol(x, y int) bool {
	if x >= f.x || y >= f.y || x < 0 || y < 0 {
		return false
	}

	if f.text[x][y] != '.' {
		return true
	}
	return false
}

func main() {
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1 := 0
	part2 := 0
	field := field{}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		field.values = append(field.values, parse(line))
		field.text = append(field.text, line)
		field.x = len(line)
		field.y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := range field.values {
		for j := range field.values[i] {
			v := field.values[i][j]
			if field.hasAdjacent(i, v.x_start, v.x_end) {
				part1 += v.value
			}
		}
	}

	log.Printf("Part1 : %d\n", part1)
	log.Printf("Part2 : %d\n", part2)
}

func parse(str string) []symbol {
	result := make([]symbol, 0, 0)
	sym := symbol{x_start: -1}
	for i, s := range str {
		if isNumber(s) {
			if sym.x_start == -1 {
				sym.x_start = i
			}
			continue
		}

		if sym.x_start != -1 {
			sym.text = str[sym.x_start:i]
			sym.x_end = i - 1
			value, err := strconv.Atoi(sym.text)
			if err != nil {
				log.Fatalln(err)
			}
			sym.value = value
			result = append(result, sym)
			sym = symbol{x_start: -1}
		}

		if s == '.' {
			continue
		}
	}
	if sym.x_start != -1 {
		sym.text = str[sym.x_start:]
		sym.x_end = len(str) - 1
		value, err := strconv.Atoi(sym.text)
		if err != nil {
			log.Fatalln(err)
		}
		sym.value = value
		result = append(result, sym)
	}

	return result
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}
