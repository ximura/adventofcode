package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputs = "../inputs/day03.txt"

type symbol struct {
	x_start int
	x_end   int
	text    string
	value   int
}

type field struct {
	values [][]symbol
	text   []string
	stars  map[string][]int
	x      int
	y      int
}

func (f *field) addStar(x, y, value int) {
	index := fmt.Sprintf("%d,%d", x, y)
	s, ok := f.stars[index]
	if !ok {
		f.stars[index] = []int{value}
	} else {
		f.stars[index] = append(s, value)
	}
}

func (f *field) hasAdjacent(x, start, end, value int) bool {
	t := f.getSymbol(x, start-1)
	if t == 2 {
		f.addStar(x, start-1, value)
	}
	t1 := f.getSymbol(x, end+1)
	if t1 == 2 {
		f.addStar(x, end+1, value)
	}
	if t > 0 || t1 > 0 {
		return true
	}

	for i := start - 1; i <= end+1; i++ {
		t := f.getSymbol(x+1, i)
		if t == 2 {
			f.addStar(x+1, i, value)
		}
		t1 := f.getSymbol(x-1, i)
		if t1 == 2 {
			f.addStar(x-1, i, value)
		}
		if t > 0 || t1 > 0 {
			return true
		}
	}

	return false
}

// getSymbol
// 0 - not found
// 1 - found
// 2 - found star
func (f *field) getSymbol(x, y int) int {
	if x >= f.x || y >= f.y || x < 0 || y < 0 {
		return 0
	}

	if f.text[x][y] != '.' {
		if f.text[x][y] == '*' {
			return 2
		}

		return 1
	}
	return 0
}

func main() {
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1 := 0
	part2 := 0
	field := field{stars: make(map[string][]int)}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		symbols, _ := parse(line)
		field.values = append(field.values, symbols)
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
			if field.hasAdjacent(i, v.x_start, v.x_end, v.value) {
				part1 += v.value
			}
		}
	}

	for _, s := range field.stars {
		if len(s) <= 1 {
			continue
		}

		part2 += s[0] * s[1]
	}

	log.Printf("Part1 : %d\n", part1)
	log.Printf("Part2 : %d\n", part2)
}

func parse(str string) ([]symbol, []int) {
	result := make([]symbol, 0, 0)
	stars := make([]int, 0, 0)
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

		if s == '*' {
			stars = append(stars, i)
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

	return result, stars
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}
