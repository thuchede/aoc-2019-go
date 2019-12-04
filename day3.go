package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Grid represent the electrical panel where wires are connected
type Grid struct {
	grid [][]string
}

// NewGrid return a new grid of size length * width
func NewGrid(length int, width int) *Grid {
	g := new(Grid)
	g.grid = make([][]string, length)
	for i := 0; i < length; i++ {
		g.grid[i] = make([]string, width)
		for j := range g.grid[i] {
			g.grid[i][j] = "."
		}
	}
	return g
}

func (g *Grid) printGrid() {
	for i := range g.grid {
		fmt.Printf("%v\n", g.grid[i])
	}
}

func (g *Grid) findOrigin() (int, int) {
	for i := range g.grid {
		for j := range g.grid[i] {
			if g.grid[i][j] == "o" {
				return i, j
			}
		}
	}
	return -1, -1
}

type Key struct {
	X, Y int
}

func equal(m, obj map[Key][]int) bool {
	if len(m) != len(obj) {
		return false
	}
	for key, value := range m {
		if len(value) != len(obj[key]) {
			return false
		}
		for i := range value {
			if value[i] != obj[key][i] {
				return false
			}
		}
	}
	return true
}

func equalDist(m, obj map[Key][]Key) bool {
	if len(m) != len(obj) {
		return false
	}
	for key, value := range m {
		if len(value) != len(obj[key]) {
			return false
		}
		for i := range value {
			if value[i].X != obj[key][i].X && value[i].Y != obj[key][i].Y {
				return false
			}
		}
	}
	return true
}

// AddWire return the map with updated wire position
func AddWire(grid map[Key][]int, wirePath []string, num int) map[Key][]int {
	currentX, currentY := 0, 0
	for _, value := range wirePath {
		action := string(value[0])
		steps, _ := strconv.Atoi(string(value[1:]))
		for i := 1; i <= steps; i++ {
			if action == "R" {
				currentX++
			}
			if action == "L" {
				currentX--
			}
			if action == "U" {
				currentY++
			}
			if action == "D" {
				currentY--
			}
			if !Contains(grid[Key{X: currentX, Y: currentY}], num) {
				grid[Key{X: currentX, Y: currentY}] = append(grid[Key{X: currentX, Y: currentY}], num)
			}
		}
	}
	return grid
}

// ComputeManhattanDistance return the manhattan distance between two coordinates
func ComputeManhattanDistance(sourceX int, sourceY int) int {
	var x, y int
	if sourceX < 0 {
		x = -sourceX
	} else {
		x = sourceX
	}
	if sourceY < 0 {
		y = -sourceY
	} else {
		y = sourceY
	}
	return x + y
}

// FindIntersection return position of key from the grid starting its search from (startX, startY)
func FindIntersection(grid map[Key][]int) (int, int) {
	return -1, -1
}

// Contains return true if array a contains x
func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// ContainsDist return true if array a contains x
func ContainsDist(a []Key, x int) bool {
	for _, n := range a {
		if x == n.X {
			return true
		}
	}
	return false
}

// FindLowestManhattanDistance return the lowest manhattan distance between intersection point and central port
func FindLowestManhattanDistance(grid map[Key][]int) int {
	lowestManhattan := int(^uint(0) >> 1)

	for key, value := range grid {
		if len(value) > 1 {
			man := ComputeManhattanDistance(key.X, key.Y)
			if man < lowestManhattan {
				lowestManhattan = man
			}
		}
	}
	return lowestManhattan
}

// ReadFileDay3 executes day3 puzzle
func ReadFileDay3() int {
	file, err := os.Open("./input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make(map[Key][]int)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	path1 := strings.Split(scanner.Text(), ",")

	scanner.Scan()
	path2 := strings.Split(scanner.Text(), ",")

	grid1 := AddWire(grid, path1, 1)
	grid2 := AddWire(grid1, path2, 2)

	return FindLowestManhattanDistance(grid2)
}

// FindLowestPathDistance return the lowest path distance between intersection point and central port
func FindLowestPathDistance(grid map[Key][]Key) int {
	lowestDistance := int(^uint(0) >> 1)

	for _, value := range grid {
		if len(value) > 1 {
			man := value[0].Y + value[1].Y
			if man < lowestDistance {
				lowestDistance = man
			}
		}
	}
	return lowestDistance
}

// AddWireWithDistance return the lowest path distance between intersection point and central port
func AddWireWithDistance(grid map[Key][]Key, wirePath []string, num int) map[Key][]Key {
	currentX, currentY := 0, 0
	totalSteps := 0
	for _, value := range wirePath {
		action := string(value[0])
		steps, _ := strconv.Atoi(string(value[1:]))
		for i := 1; i <= steps; i++ {
			if action == "R" {
				currentX++
			}
			if action == "L" {
				currentX--
			}
			if action == "U" {
				currentY++
			}
			if action == "D" {
				currentY--
			}
			totalSteps++
			if !ContainsDist(grid[Key{X: currentX, Y: currentY}], num) {
				grid[Key{X: currentX, Y: currentY}] = append(grid[Key{X: currentX, Y: currentY}], Key{X: num, Y: totalSteps})
			}
		}
	}
	return grid
}

// ReadFileDay3Part2 executes day3 puzzle
func ReadFileDay3Part2() int {
	file, err := os.Open("./input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make(map[Key][]Key)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	path1 := strings.Split(scanner.Text(), ",")

	scanner.Scan()
	path2 := strings.Split(scanner.Text(), ",")

	grid1 := AddWireWithDistance(grid, path1, 1)
	grid2 := AddWireWithDistance(grid1, path2, 2)

	// fmt.Printf("\n%v", grid2)

	return FindLowestPathDistance(grid2)
}
