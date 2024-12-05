package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

var inputs = "../inputs/day03.txt"

type Command struct {
	cmd          string
	a, b, result int
}

func part1(cmd []Command) int {
	result := 0
	for i := range cmd {
		result += cmd[i].result
	}

	return result
}

func part2(cmd []Command) int {
	result := 0
	enabled := true
	for i := range cmd {
		c := cmd[i]

		if !enabled {
			enabled = c.cmd == "do"
		} else {
			enabled = c.cmd != "don't"
		}

		if enabled {
			result += c.result
		}
	}

	return result
}

func main() {
	timeStart := time.Now()
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cmd, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(cmd))
	fmt.Printf("Part 2: %d\n", part2(cmd))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) ([]Command, error) {
	commands := make([]Command, 0, 10)

	mulRegexp, err := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)|don't|do`)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		matches := mulRegexp.FindAllSubmatch(line, -1)
		for _, group := range matches {
			cmd := Command{
				cmd: string(group[0]),
			}
			if len(group) == 3 {
				a, _ := strconv.Atoi(string(group[1]))
				b, _ := strconv.Atoi(string(group[2]))

				cmd.a = a
				cmd.b = b
				cmd.result = a * b
			}

			commands = append(commands, cmd)
		}
	}

	if err := scanner.Err(); err != nil {
		return commands, err
	}

	return commands, nil
}
