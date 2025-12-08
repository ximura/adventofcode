package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var inputs = "../inputs/day08.txt"

type Input struct {
	Boxes []Point
}

type Point struct {
	X, Y, Z int
}

func main() {
	file, err := os.Open(inputs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d := parseFile(file)

	timeStart := time.Now()
	p1, p2 := calculate(d)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) *Input {
	p := &Input{
		Boxes: make([]Point, 0),
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r")
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		p.Boxes = append(p.Boxes, Point{X: x, Y: y, Z: z})

	}

	return p
}

type Edge struct {
	a, b int
	d2   int // squared distance
}

// =============== UNION FIND ==================

type UF struct {
	parent []int
	size   []int
}

func NewUF(n int) *UF {
	p := make([]int, n)
	s := make([]int, n)
	for i := range p {
		p[i] = i
		s[i] = 1
	}
	return &UF{p, s}
}

func (u *UF) Find(x int) int {
	for x != u.parent[x] {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *UF) Union(a, b int) bool {
	ra := u.Find(a)
	rb := u.Find(b)
	if ra == rb {
		return false
	}
	if u.size[ra] < u.size[rb] {
		ra, rb = rb, ra
	}
	u.parent[rb] = ra
	u.size[ra] += u.size[rb]
	return true
}

// Return all component sizes
func (u *UF) ComponentSizes() []int {
	sizes := map[int]int{}
	for i := range u.parent {
		root := u.Find(i)
		sizes[root] = u.size[root]
	}
	out := make([]int, 0, len(sizes))
	for _, v := range sizes {
		out = append(out, v)
	}
	return out
}

func calculate(data *Input) (int, int) {
	part1 := 0
	part2 := 0
	points := data.Boxes

	n := len(points)
	// ---- Build all edges O(n^2) ----
	edges := make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d2 := dist2(points[i], points[j])
			edges = append(edges, Edge{i, j, d2})
		}
	}

	// ---- Sort edges by distance O(n^2 log n) ----
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].d2 < edges[j].d2
	})

	uf := NewUF(n)
	added := 0

	for _, e := range edges {
		if uf.Union(e.a, e.b) {
			added++

			if added == n-1 {
				part2 = points[e.a].X * points[e.b].X
				break
			}
		}
	}

	return part1, part2
}

func dist2(a, b Point) int {
	x := a.X - b.X
	y := a.Y - b.Y
	z := a.Z - b.Z
	return x*x + y*y + z*z
}
