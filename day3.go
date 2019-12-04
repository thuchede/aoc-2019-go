package main

import (
	"fmt"
	"strconv"
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

func (g *Grid) equal(obj *Grid) bool {
	if len(g.grid) != len(obj.grid) {
		return false
	}
	for i := range g.grid {
		if len(g.grid[i]) != len(obj.grid[i]) {
			return false
		}
		for j := range g.grid[i] {
			if g.grid[i][j] != obj.grid[i][j] {
				return false
			}
		}
	}
	return true
}

// AddWire return the grid
func AddWire(grid *Grid, wirePath []string) *Grid {
	x, y := grid.findOrigin()
	for i := range wirePath {
		if string(wirePath[i][0]) == "R" {
			steps, _ := strconv.Atoi(string(wirePath[i][1:]))
			for j := 0; j <= steps; j++ {
				if grid.grid[x][y+j] == "|" {
					if j == 0 {
						grid.grid[x][y+j] = "+"
					} else {
						grid.grid[x][y+j] = "x"
					}
				} else if grid.grid[x][y+j] != "o" {
					grid.grid[x][y+j] = "-"
				}
			}
			y += steps
		}
		if string(wirePath[i][0]) == "L" {
			steps, _ := strconv.Atoi(string(wirePath[i][1:]))
			for j := 0; j <= steps; j++ {
				if grid.grid[x][y-j] == "|" {
					if j == 0 {
						grid.grid[x][y-j] = "+"
					} else {
						grid.grid[x][y-j] = "x"
					}
				} else if grid.grid[x][y-j] != "o" {
					grid.grid[x][y-j] = "-"
				}
			}
			y -= steps
		}
		if string(wirePath[i][0]) == "U" {
			steps, _ := strconv.Atoi(string(wirePath[i][1:]))
			for j := 0; j <= steps; j++ {
				if grid.grid[x-j][y] == "-" {
					if j == 0 {
						grid.grid[x-j][y] = "+"
					} else {
						grid.grid[x-j][y] = "x"
					}
				} else if grid.grid[x-j][y] != "o" {
					grid.grid[x-j][y] = "|"
				}
			}
			x -= steps
		}
		if string(wirePath[i][0]) == "D" {
			steps, _ := strconv.Atoi(string(wirePath[i][1:]))
			for j := 0; j <= steps; j++ {
				if grid.grid[x+j][y] == "-" {
					if j == 0 {
						grid.grid[x+j][y] = "+"
					} else {
						grid.grid[x+j][y] = "x"
					}
				} else if grid.grid[x+j][y] != "o" {
					grid.grid[x+j][y] = "|"
				}
			}
			x += steps
		}
	}
	return grid
}

// ComputeManhattanDistance return the manhattan distance between two coordinates
func ComputeManhattanDistance(originX int, originY int, destinationX int, destinationY int) int {
	x := originX - destinationX
	if x < 0 {
		x = destinationX - originX
	}
	y := originY - destinationY
	if y < 0 {
		y = destinationY - originY
	}
	return x + y
}

// FindLocation return position of key from the grid starting its search from (startX, startY)
func (g *Grid) FindLocation(startX int, startY int, key string) (int, int) {
	for i := startX; i < len(g.grid); i++ {
		for j := startY + 1; j < len(g.grid[i]); j++ {
			if g.grid[i][j] == key {
				return i, j
			}
		}
		startY = 0
	}
	return -1, -1
}

// FindLowestManhattanDistance return the lowest manhattan distance between intersection point and central port
func (g *Grid) FindLowestManhattanDistance() int {
	originX, originY := g.findOrigin()

	lowestManhattan := int(^uint(0) >> 1)

	for i, j := 0, -1; i >= 0 && i < len(g.grid); {
		i, j = g.FindLocation(i, j, "x")
		man := ComputeManhattanDistance(originX, originY, i, j)
		if man < lowestManhattan {
			lowestManhattan = man
		}
	}
	return lowestManhattan
}
