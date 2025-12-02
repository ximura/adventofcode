package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var inputs = "../inputs/day01.txt"

type Data = []entity

type entity struct {
	direction bool // R true; L false
	amount    int
	original  string
}

func main() {
	timeStart := time.Now()
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	p1, p2 := calcualte(d)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) (Data, error) {
	d := Data{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		direction := string(line[0])
		amount, _ := strconv.Atoi(string(line[1:]))

		d = append(d, entity{
			direction: direction == "R",
			amount:    amount,
			original:  string(line),
		})
	}

	return d, nil
}

func calcualte(d Data) (int, int) {
	start := 50
	p1, p2 := 0, 0

	//fmt.Println("The dial starts by pointing at 50.")
	for i := range d {
		step := d[i]
		amount := step.amount % 100
		fullCircle := step.amount / 100
		startWithZero := start == 0
		passZero := fullCircle

		if step.direction {
			start += amount
		} else {
			start -= amount
		}

		if start == 0 || start == 100 {
			p1++
			passZero++
			start = 0
		}

		if start > 100 {
			start -= 100
			passZero++
		} else if start < 0 {
			start = 100 + start
			if !startWithZero {
				passZero++
			}
		}

		p2 += passZero

		// fmt.Printf("The dial is rotated %s to point at %d.\n", step.original, start)
		// if passZero > 0 {
		// 	fmt.Printf("during this rotation, it points at 0: %d.\n", passZero)
		// }
	}

	return p1, p2
}

func part1(d Data) int {
	start := 50
	count := 0

	//fmt.Println("The dial starts by pointing at 50.")
	for i := range d {
		step := d[i]
		if step.direction {
			start += step.amount % 100
		} else {
			start -= step.amount % 100
		}

		if start >= 100 {
			start -= 100
		} else if start < 0 {
			start = 100 + start
		}

		if start == 0 {
			count++
		}

		//fmt.Printf("The dial is rotated %s to point at %d.\n", step.original, start)
	}

	return count
}

func part2(d Data) int {
	start := 50
	count := 0

	//fmt.Println("The dial starts by pointing at 50.")
	for i := range d {
		step := d[i]
		amount := step.amount % 100
		fullCircle := step.amount / 100
		startWithZero := start == 0
		passZero := fullCircle

		if step.direction {
			start += amount % 100
		} else {
			start -= amount % 100
		}

		if start >= 100 {
			start -= 100
			passZero++
		} else if start < 0 {
			start = 100 + start
			if !startWithZero {
				passZero++
			}
		} else if start == 0 {
			passZero++
		}

		count += passZero

		// fmt.Printf("The dial is rotated %s to point at %d.\n", step.original, start)
		// if passZero > 0 {
		// 	fmt.Printf("during this rotation, it points at 0: %d.\n", passZero)
		// }
	}

	return count
}
