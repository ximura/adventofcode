package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var inputs = "../inputs/day03.txt"

type Data = []string

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
		line := scanner.Text()
		d = append(d, line)
	}

	return d, nil
}

func calcualte(d Data) (int64, int64) {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	//fmt.Printf("Using %d CPU threads\n", numCPU)

	var wg sync.WaitGroup
	var part1 atomic.Int64
	var part2 atomic.Int64

	for i := range d {
		wg.Add(1)

		go func(in string) {
			defer wg.Done()
			p1 := bestPair(d[i])
			p2 := bestN(in, 12)
			//fmt.Printf("In %s, you can make the largest joltage possible, %d\n", d[i], p2)
			part1.Add(p1)
			part2.Add(p2)
		}(d[i])
	}
	wg.Wait()

	return part1.Load(), part2.Load()
}

func bestPair(line string) int64 {
	n := len(line)
	if n < 2 {
		return -1
	}

	best := -1
	maxRight := -1 // the largest digit seen so far from the right side

	for i := n - 1; i >= 0; i-- {
		cur := int(line[i] - '0')

		if maxRight != -1 { // if something exists to the right
			val := cur*10 + maxRight
			if val > best {
				best = val
			}
		}

		// update maxRight (suffix best digit)
		if cur > maxRight {
			maxRight = cur
		}
	}

	return int64(best)
}

func bestN(line string, k int) int64 {
	n := len(line)
	toRemove := n - k
	stack := make([]byte, 0, k)

	for i := 0; i < n; i++ {
		d := line[i]
		// While we can remove digits and the last digit in stack is smaller
		for toRemove > 0 && len(stack) > 0 && stack[len(stack)-1] < d {
			stack = stack[:len(stack)-1]
			toRemove--
			// continue popping while beneficial
		}
		stack = append(stack, d)
	}

	// If still need to remove digits, remove from the end
	if len(stack) > k {
		stack = stack[:k]
	}

	var result int64 = 0
	for _, b := range stack {
		result = result*10 + int64(b-'0')
	}
	return result
}
