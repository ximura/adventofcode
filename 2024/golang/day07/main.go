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

type OperationType int8

const (
	OperationNone OperationType = -1
	OperationSum  OperationType = 0
	OperationMul  OperationType = 1
)

func (o OperationType) String() string {
	if o == OperationSum {
		return "+"
	}

	if o == OperationMul {
		return "*"
	}

	return ""
}

type test struct {
	result int
	values []int
	valid  bool
}

func part1(input []test) int {
	result := 0

	for _, t := range input {
		bCount := len(t.values) - 1
		combinationCount := bCount * 2
		tmp := make([]OperationType, bCount)
		for i := 0; i <= combinationCount; i++ {
			value := t.values[0]
			for j := 0; j < bCount; j++ {
				if hasBit(i, uint(j)) {
					value *= t.values[j+1]
					tmp[j] = OperationMul
				} else {
					value += t.values[j+1]
					tmp[j] = OperationSum
				}

				if value > t.result {
					break
				}
			}

			if value == t.result {
				result += value
				t.valid = true

				// for i, v := range t.values {
				// 	fmt.Printf("%d ", v)
				// 	if i < bCount {
				// 		fmt.Printf("%s ", tmp[i])
				// 	}
				// }
				// fmt.Printf(" = %d\n", t.result)

				fmt.Printf("%d : %+v\n", t.result, t.values)

				break
			}
		}
	}

	return result
}

func part2(input []test) int {
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

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
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

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}
