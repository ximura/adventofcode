package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var inputs = "../inputs/day06.txt"

type Input = []Data

type Data struct {
	numbers    []int
	opperation string
}

func (d Data) calculate() int {
	res := 0
	if d.opperation == "+" {
		for i := range d.numbers {
			res += d.numbers[i]
		}
	}

	if d.opperation == "*" {
		res = 1
		for i := range d.numbers {
			res *= d.numbers[i]
		}
	}

	return res
}

func main() {
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d1, d2 := parseFile(file)

	timeStart := time.Now()
	fmt.Printf("Part 1: %d\n", calculate(d1))
	fmt.Printf("Part 2: %d\n", calculate(d2))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) (Input, Input) {
	lines := []string{}
	maxLen := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		l := len(line)
		if l > maxLen {
			maxLen = l
		}
	}

	return createInput1(lines, maxLen), createInput2(lines, maxLen)
}

func createInput1(lines []string, maxLen int) Input {
	d := Input{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(d) == 0 {
			d = make([]Data, len(parts))
		}

		i := 0
		for j := range parts {
			s := parts[j]
			if s == "" {
				continue
			}

			n, err := strconv.Atoi(s)
			if err == nil {
				d[i].numbers = append(d[i].numbers, n)
			} else {
				d[i].opperation = s
			}
			i++
		}
	}

	return d
}

func createInput2(lines []string, maxLen int) Input {
	input := Input{}

	// Transpose (rows → columns)
	columns := make([]string, maxLen)
	for col := 0; col < maxLen; col++ {
		b := make([]byte, len(lines))
		for row := 0; row < len(lines); row++ {
			b[row] = lines[row][col]
		}
		columns[col] = string(b)
	}

	var d Data
	for i := range columns {
		val := columns[i]
		if strings.TrimSpace(val) == "" {
			input = append(input, d)
			d = Data{}
			continue
		}

		l := len(val)
		op := val[l-1]
		if op == 42 || op == 43 { // * or +
			d.opperation = string(op)
			val = val[:l-1]
		}
		n, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			panic(err)
		}
		d.numbers = append(d.numbers, n)
	}
	input = append(input, d)

	return input
}

func calculate(data Input) int {
	part1 := 0

	for _, d := range data {
		part1 += d.calculate()
	}

	return part1
}
