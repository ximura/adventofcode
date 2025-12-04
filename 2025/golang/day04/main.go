package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var inputs = "../inputs/day04.txt"

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
		if line == "" {
			break
		}
		d = append(d, line)
	}

	return d, nil
}

// Directions for 8 dirs
var dirs = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func calcualte(grid Data) (int, int) {
	h := len(grid)
	w := len(grid[0])

	// Represent rolls as a set of coordinates for fast removal
	rolls := make(map[[2]int]bool)
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if grid[r][c] == '@' {
				rolls[[2]int{r, c}] = true
			}
		}
	}

	firstRemoved := 0
	totalRemoved := 0
	round := 1

	for {
		accessible := make([][2]int, 0)

		// Find accessible rolls this round
		for pos := range rolls {
			r, c := pos[0], pos[1]
			adj := 0

			for _, d := range dirs {
				rr := r + d[0]
				cc := c + d[1]
				if rr < 0 || rr >= h || cc < 0 || cc >= w {
					continue
				}
				if rolls[[2]int{rr, cc}] {
					adj++
				}
			}

			if adj < 4 {
				accessible = append(accessible, pos)
			}
		}

		// No more rolls can be removed
		if len(accessible) == 0 {
			break
		}

		// Visualization grid
		//visualizeGrid(round, h, w, rolls, accessible)

		// Remove all accessible rolls
		for _, pos := range accessible {
			delete(rolls, pos)
		}

		totalRemoved += len(accessible)
		if firstRemoved == 0 {
			firstRemoved = totalRemoved
		}
		round++
	}

	return firstRemoved, totalRemoved
}

// non generic solution to part one
func part1(grid Data) int {
	h := len(grid)
	w := len(grid[0])

	// Copy grid for marking
	marked := make([][]rune, h)
	for i := range grid {
		marked[i] = []rune(grid[i])
	}

	accessibleCount := 0

	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if grid[r][c] != '@' {
				continue
			}

			// Count adjacent rolls
			adj := 0
			for _, d := range dirs {
				rr := r + d[0]
				cc := c + d[1]
				if rr >= 0 && rr < h && cc >= 0 && cc < w {
					if grid[rr][cc] == '@' {
						adj++
					}
				}
			}

			// Accessible if fewer than 4 neighbors
			if adj < 4 {
				marked[r][c] = 'x'
				accessibleCount++
			}
		}
	}

	// Print result
	// for _, row := range marked {
	// 	fmt.Println(string(row))
	// }
	// fmt.Println("Accessible count:", accessibleCount)

	return accessibleCount
}

// helper function for grid visualization
func visualizeGrid(round, h, w int, rolls map[[2]int]bool, accessible [][2]int) {
	viz := make([][]rune, h)
	for i := 0; i < h; i++ {
		viz[i] = make([]rune, w)
		for j := 0; j < w; j++ {
			if rolls[[2]int{i, j}] {
				viz[i][j] = '@'
			} else {
				viz[i][j] = '.'
			}
		}
	}

	// Mark the newly removed as 'x' in the visualization
	for _, pos := range accessible {
		r, c := pos[0], pos[1]
		viz[r][c] = 'x'
	}

	// Print the round
	fmt.Printf("=== Round %d (removing %d rolls) ===\n", round, len(accessible))
	for _, row := range viz {
		fmt.Println(string(row))
	}
	fmt.Println()
}
