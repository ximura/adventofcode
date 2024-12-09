package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var inputs = "../inputs/day09.txt"

type blockInfo struct {
	start, size, value int
}

type input struct {
	disk []int
	gaps []blockInfo
	data []blockInfo
}

func part1(input input) int {
	checkSum := 0
	disk := make([]int, len(input.disk))
	gaps := input.gaps
	copy(disk, input.disk)

	end := len(disk)
	for _, g := range gaps {
		i := 0
		moved := 0

		for k := end - 1; moved < g.size && g.start+i < k; k-- {
			v := disk[k]
			if v == -1 {
				continue
			}
			j := g.start + i
			disk[j] = v
			disk[k] = -1
			i++
			moved++
		}

		end -= g.size
	}

	for i, v := range disk {
		if v == -1 {
			break
		}
		checkSum += i * v
	}

	return checkSum
}

func part2(input input) int {
	checkSum := 0
	disk := make([]int, len(input.disk))
	gaps := input.gaps
	copy(disk, input.disk)

	for i := len(input.data) - 1; i >= 0; i-- {
		b := input.data[i]
		for j, g := range gaps {

			if b.size > g.size {
				continue
			}

			if g.start > b.start {
				break
			}

			for k := 0; k < b.size; k++ {
				j := g.start + k
				disk[j] = b.value
				disk[b.start+b.size-k-1] = -1
			}

			g.size -= b.size
			g.start += b.size
			gaps[j] = g
			break
		}
	}

	for i, v := range disk {
		if v == -1 {
			continue
		}
		checkSum += i * v
	}

	return checkSum
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

	p1 := part1(input)
	p2 := part2(input)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) (input, error) {
	input := input{
		disk: make([]int, 0, 10),
		gaps: make([]blockInfo, 0, 10),
		data: make([]blockInfo, 0, 10),
	}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		counter := 0
		for i, c := range line {
			v := byteToInt([]byte{c})
			value := -1
			if i%2 == 0 {
				value = counter
				counter++
				input.data = append(input.data, blockInfo{start: len(input.disk), size: v, value: value})
			} else {
				input.gaps = append(input.gaps, blockInfo{start: len(input.disk), size: v})
			}

			for j := 0; j < v; j++ {
				input.disk = append(input.disk, value)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return input, err
	}

	return input, nil
}

func byteToInt(b []byte) int {
	var value int
	for _, b := range b {
		value = value*10 + int(b-48)
	}
	return value
}
