package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

var inputs = "../inputs/day01.txt"

type pair struct {
	l, r int
}

type Data struct {
	left    []int
	leftMap map[int]pair
	right   []int
}

func (d *Data) InsertLeft(v int) {
	p, _ := d.leftMap[v]
	p.l += 1
	d.leftMap[v] = p
	d.left = insertSorted(d.left, v)
}

func (d *Data) InsertRight(v int) {
	d.right = insertSorted(d.right, v)
}

func (d *Data) Distance() int {
	n := len(d.left)
	result := 0
	for i := 0; i < n; i++ {
		result += abs(d.left[i] - d.right[i])
	}

	return result
}

func (d *Data) Simularity() int {
	result := 0
	for _, v := range d.right {
		p, _ := d.leftMap[v]
		p.r += 1
		d.leftMap[v] = p
	}

	for k, v := range d.leftMap {
		result += k * v.l * v.r
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func intCmp(a, b int) int {
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}

	return -1
}

func insertSorted(slice []int, v int) []int {
	n := len(slice)
	i, _ := slices.BinarySearchFunc(slice, v, intCmp) // find slot
	if i == n {
		// add element in the end
		slice = append(slice, v)
		return slice
	}

	c := cap(slice)
	a := slice
	if n == c {
		// we need extend array
		a = append(a, 0)
	}
	a = a[:n+1]
	copy(a[i+1:], a[i:])
	a[i] = v
	return a
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

	fmt.Printf("Part 1: %d\n", d.Distance())
	fmt.Printf("Part 2: %d\n", d.Simularity())
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) (Data, error) {
	d := Data{
		left:    make([]int, 0, 10),
		right:   make([]int, 0, 10),
		leftMap: make(map[int]pair, 10),
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		p := bytes.Split(line, []byte(" "))
		l := len(p)
		if l < 2 {
			return d, fmt.Errorf("unsupported line format %s", string(line))
		}

		left, _ := strconv.Atoi(string(p[0]))
		right, _ := strconv.Atoi(string(p[l-1]))

		d.InsertLeft(left)
		d.InsertRight(right)
	}

	if err := scanner.Err(); err != nil {
		return d, err
	}

	return d, nil
}
