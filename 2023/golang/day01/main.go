package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

var ErrRuneNotInt = errors.New("type: rune was not int")

func main() {
	file, err := os.Open("../inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		n := filterNum(line)
		result += n
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}

func filterNum(str string) int {
	nums := make([]int, 0, 2)
	for _, val := range str {
		num, err := charToNum(val)
		if err != nil {
			continue
		}
		if len(nums) < 2 {
			nums = append(nums, num)
		} else {
			nums[1] = num
		}
	}
	if len(nums) == 1 {
		nums = append(nums, nums[0])
	}

	return nums[0]*10 + nums[1]
}

func charToNum(r rune) (int, error) {
	if '0' <= r && r <= '9' {
		return int(r) - '0', nil
	}
	return 0, ErrRuneNotInt
}
