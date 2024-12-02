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
)

var inputs = "../inputs/day01.txt"

type Data struct {
	left  []int
	right []int
}

func (d *Data) InsertLeft(v int) {
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
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(d.Distance())
}

func parseFile(r io.Reader) (Data, error) {
	d := Data{
		left:  make([]int, 0, 2),
		right: make([]int, 0, 2),
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
