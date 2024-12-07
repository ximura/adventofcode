package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"slices"
	"time"
)

var inputs = "../inputs/day06.txt"

type DirectionType int8

const (
	DirectionNone  DirectionType = -1
	DirectionUp    DirectionType = 0
	DirectionRight DirectionType = 1
	DirectionDown  DirectionType = 2
	DirectionLeft  DirectionType = 3
)

type moves map[point]struct{}
type states map[guard]struct{}
type obstacles map[int][]point

func (o obstacles) removeValue(k int, v point) {
	t, ok := o[k]
	if !ok {
		return
	}

	o[k] = slices.DeleteFunc(t, func(p point) bool {
		return p == v
	})
}

type point struct {
	x, y int
}

type guard struct {
	point
	direction DirectionType
}

func (g *guard) Direction() string {
	switch g.direction {
	case DirectionUp:
		return "up"

	case DirectionDown:
		return "down"

	case DirectionLeft:
		return "left"

	case DirectionRight:
		return "right"
	}

	return "none"
}

func (g *guard) changeDirection() {
	g.direction++
	g.direction %= 4
	if g.direction > DirectionLeft {
		g.direction = DirectionUp
	}
}

func (g *guard) moveToClosesObstacle(m *Map, obs *point) bool {
	if obs != nil {
		m.colObstacles[obs.x] = append(m.colObstacles[obs.x], *obs)
		m.rowObstacles[obs.y] = append(m.rowObstacles[obs.y], *obs)
	}

	defer func() {
		if obs != nil {
			m.colObstacles.removeValue(obs.x, *obs)
			m.rowObstacles.removeValue(obs.y, *obs)
		}
	}()

	switch g.direction {
	case DirectionUp:
		// look into columns
		obs := m.colObstacles[g.x]
		min := m.maxY
		targetY := -1
		for _, o := range obs {
			dist := g.y - o.y
			if dist > 0 && dist < min {
				min = dist
				targetY = o.y
			}
		}

		count := g.y - targetY
		for i := 1; i < count; i++ {
			if !addMove(g.x, g.y-i, g.direction, m) {
				return false
			}
		}

		g.y = targetY + 1

	case DirectionDown:
		// look into columns
		obs := m.colObstacles[g.x]
		min := m.maxY
		targetY := m.maxY + 1
		for _, o := range obs {
			dist := o.y - g.y
			if dist > 0 && dist < min {
				min = dist
				targetY = o.y
			}
		}

		count := targetY - g.y
		for i := 1; i < count; i++ {
			if !addMove(g.x, g.y+i, g.direction, m) {
				return false
			}
		}

		g.y = targetY - 1

	case DirectionLeft:
		// look into rows
		obs := m.rowObstacles[g.y]
		min := m.maxX
		targetX := -1
		for _, o := range obs {
			dist := g.x - o.x
			if dist > 0 && dist < min {
				min = dist
				targetX = o.x
			}
		}

		count := g.x - targetX
		for i := 1; i < count; i++ {
			if !addMove(g.x-i, g.y, g.direction, m) {
				return false
			}
		}

		g.x = targetX + 1

	case DirectionRight:
		// look into rows
		obs := m.rowObstacles[g.y]
		min := m.maxX
		targetX := m.maxX + 1
		for _, o := range obs {
			dist := o.x - g.x
			if dist > 0 && dist < min {
				min = dist
				targetX = o.x
			}
		}

		count := targetX - g.x
		for i := 1; i < count; i++ {
			if !addMove(g.x+i, g.y, g.direction, m) {
				return false
			}
		}

		g.x = targetX - 1
	}

	return true
}

func (g *guard) canLeave(maxX, maxY int) bool {
	result := false
	switch g.direction {
	case DirectionUp:
		result = g.y == 0
	case DirectionDown:
		result = g.y == maxX
	case DirectionLeft:
		result = g.x == 0
	case DirectionRight:
		result = g.x == maxY
	}

	return result
}

func newGuard(x, y int, c byte) guard {
	direction := DirectionNone
	switch c {
	case '^':
		direction = DirectionUp
	case '>':
		direction = DirectionLeft
	case '<':
		direction = DirectionRight
	case 'v':
		direction = DirectionDown
	}

	return guard{
		point:     point{x: x, y: y},
		direction: direction,
	}
}

func addMove(x, y int, d DirectionType, m *Map) bool {
	p := point{x: x, y: y}
	m.moves[p] = struct{}{}

	g := guard{point: p, direction: d}
	if _, ok := m.state[g]; ok {
		return false
	}
	m.state[g] = struct{}{}

	return true
}

type Map struct {
	//field        [][]byte
	maxX, maxY   int
	colObstacles obstacles // y corrdinate mapped to obstacles in column
	rowObstacles obstacles // x corrdinate mapped to obstacles in row
	moves        moves
	state        states
	guard        guard
	start        guard
}

func (m *Map) reset() {
	clear(m.state)
	clear(m.moves)

	m.guard.x = m.start.x
	m.guard.y = m.start.y
	m.guard.direction = m.start.direction
}

func part1(m Map) int {
	for {
		m.guard.moveToClosesObstacle(&m, nil)
		//fmt.Printf("%+v %s %d\n", m.guard.point, m.guard.Direction(), len(m.moves))

		if m.guard.canLeave(m.maxX, m.maxY) {
			break
		}

		m.guard.changeDirection()
	}

	return len(m.moves)
}

func part2(m Map) int {
	result := 0
	moves := make(moves, len(m.moves))
	maps.Copy(moves, m.moves)
	delete(moves, m.start.point)

	for move, _ := range moves {
		m.reset()
		for {
			if !m.guard.moveToClosesObstacle(&m, &move) {
				result++
				break
			}
			//fmt.Printf("%+v %s %d\n", m.guard.point, m.guard.Direction(), len(m.moves))

			if m.guard.canLeave(m.maxX, m.maxY) {
				break
			}

			m.guard.changeDirection()
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

	m, err := parseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(m))
	fmt.Printf("Part 2: %d\n", part2(m))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func parseFile(r io.Reader) (Map, error) {
	m := Map{
		//field:        make([][]byte, 0, 10),
		colObstacles: make(map[int][]point),
		rowObstacles: make(map[int][]point),
		moves:        make(moves),
		state:        make(states),
	}

	row := 0
	count := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		m.maxX = len(line) - 1
		//fmt.Println(string(line))
		for col := range line {
			c := line[col]
			if c == '#' {
				p := point{x: col, y: row}
				count++
				m.colObstacles[col] = append(m.colObstacles[col], p)
				m.rowObstacles[row] = append(m.rowObstacles[row], p)
			} else if c == '^' || c == '>' || c == 'v' || c == '<' {
				m.guard = newGuard(col, row, c)
				m.start = newGuard(col, row, c)
				addMove(col, row, m.guard.direction, &m)
			}
		}
		row++
	}
	m.maxY = row - 1
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
