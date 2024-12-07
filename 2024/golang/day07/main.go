package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var inputs = "../inputs/day07.txt"

type test struct {
	result int
	values []int
}

func (t *test) check(acc, iter int) bool {
	if acc > t.result {
		return false
	}

	sum := acc + t.values[iter]
	prod := acc * t.values[iter]
	next := iter + 1
	if len(t.values) == next {
		return sum == t.result || prod == t.result
	}

	return t.check(sum, next) || t.check(prod, next)
}

func (t *test) checkCum(acc, iter int) bool {
	if acc > t.result {
		return false
	}

	sum := acc + t.values[iter]
	prod := acc * t.values[iter]
	cum := concat(acc, t.values[iter])
	next := iter + 1
	if len(t.values) == next {
		return sum == t.result || prod == t.result || cum == t.result
	}

	return t.checkCum(sum, next) || t.checkCum(prod, next) || t.checkCum(cum, next)
}

func count(input []test) (int, int) {
	result := 0
	result2 := 0

	for _, t := range input {
		if t.check(t.values[0], 1) {
			result += t.result
		}

		if t.checkCum(t.values[0], 1) {
			result2 += t.result
		}
	}

	return result, result2
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

	r1, r2 := count(input)
	fmt.Printf("Part 1: %d\n", r1)
	fmt.Printf("Part 2: %d\n", r2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) ([]test, error) {
	m := make([]test, 0, 10)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		t := test{}

		parts := bytes.Split(line, []byte{':'})
		t.result = byteToInt(parts[0])

		values := bytes.Split(parts[1], []byte{' '})
		t.values = make([]int, 0)
		for _, v := range values {
			if len(v) == 0 {
				continue
			}
			t.values = append(t.values, byteToInt(v))
		}

		m = append(m, t)
	}

	if err := scanner.Err(); err != nil {
		return m, err
	}

	return m, nil
}

func byteToInt(b []byte) int {
	var value int
	for _, b := range b {
		value = value*10 + int(b-48)
	}
	return value
}

func concat(left, right int) int {
	shift := 0
	digits := right
	for digits > 0 {
		digits = digits / 10
		shift++
	}
	ret := left
	for i := 0; i < shift; i++ {
		ret *= 10
	}
	ret += right
	return ret
}
