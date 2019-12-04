package main

import (
	"testing"
)

func TestEqual(t *testing.T) {
	g1 := NewGrid(4, 3)
	g2 := NewGrid(4, 3)
	if !g1.equal(g2) {
		t.Error("expected new grids to have same content")
	}
	g3 := NewGrid(4, 3)
	g4 := NewGrid(3, 3)
	if g3.equal(g4) {
		t.Error("expected different sized grid to not have same content")
	}
	g5 := NewGrid(4, 3)
	g5.grid[1][1] = "o"
	g6 := NewGrid(4, 3)
	if g5.equal(g6) {
		t.Error("expected grid with different content to not have equal content")
	}
}

func TestAddWire(t *testing.T) {
	grid0 := NewGrid(10, 11)
	grid0.grid[8][1] = "o"
	path0 := []string{"R2"}
	expected0 := Grid{}
	expected0.grid = [][]string{
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", "o", "-", "-", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}

	res0 := AddWire(grid0, path0)
	if !res0.equal(&expected0) {
		t.Error("failed 0")
	}

	grid1 := NewGrid(10, 11)
	grid1.grid[8][1] = "o"
	path1 := []string{"R8", "U5", "L5", "D3"}
	expected1 := Grid{}
	expected1.grid = [][]string{
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "+", "-", "-", "-", "-", "+", "."},
		[]string{".", ".", ".", ".", "|", ".", ".", ".", ".", "|", "."},
		[]string{".", ".", ".", ".", "|", ".", ".", ".", ".", "|", "."},
		[]string{".", ".", ".", ".", "|", ".", ".", ".", ".", "|", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", "|", "."},
		[]string{".", "o", "-", "-", "-", "-", "-", "-", "-", "+", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}

	res1 := AddWire(grid1, path1)
	if !res1.equal(&expected1) {
		res1.printGrid()
		t.Error("failed 1")
	}

	grid2 := NewGrid(10, 11)
	grid2.grid[8][1] = "o"
	path21 := []string{"R8", "U5", "L5", "D3"}
	path22 := []string{"U7", "R6", "D4", "L4"}
	expected2 := Grid{}
	expected2.grid = [][]string{
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", "+", "-", "-", "-", "-", "-", "+", ".", ".", "."},
		[]string{".", "|", ".", ".", ".", ".", ".", "|", ".", ".", "."},
		[]string{".", "|", ".", ".", "+", "-", "-", "x", "-", "+", "."},
		[]string{".", "|", ".", ".", "|", ".", ".", "|", ".", "|", "."},
		[]string{".", "|", ".", "-", "x", "-", "-", "+", ".", "|", "."},
		[]string{".", "|", ".", ".", "|", ".", ".", ".", ".", "|", "."},
		[]string{".", "|", ".", ".", ".", ".", ".", ".", ".", "|", "."},
		[]string{".", "o", "-", "-", "-", "-", "-", "-", "-", "+", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}

	res21 := AddWire(grid2, path21)
	res22 := AddWire(res21, path22)
	if !res22.equal(&expected2) {
		res22.printGrid()
		expected2.printGrid()
		t.Error("failed")
	}
}

func TestComputeManhattanDistance1(t *testing.T) {
	if ComputeManhattanDistance(8, 1, 7, 2) != 2 {
		t.Error("expected ComputeManhattanDistance to be 2")
	}
	if ComputeManhattanDistance(8, 1, 0, 0) != 9 {
		t.Error("expected ComputeManhattanDistance to be 9")
	}
	if ComputeManhattanDistance(8, 1, 5, 3) != 5 {
		t.Error("expected ComputeManhattanDistance to be 5")
	}
	if ComputeManhattanDistance(0, 0, 10, 9) != 19 {
		t.Error("expected ComputeManhattanDistance to be 19")
	}
}

func TestFindLocation(t *testing.T) {
	expected := Grid{}
	expected.grid = [][]string{
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", "o", "-", "-", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}
	if x, y := expected.FindLocation(0, -1, "o"); x != 8 || y != 1 {
		t.Error("expected 8, 1")
	}
	if x, y := expected.FindLocation(0, -1, "-"); x != 8 || y != 2 {
		t.Error("expected 8, 2")
	}
	if x, y := expected.FindLocation(8, 2, "-"); x != 8 || y != 3 {
		t.Errorf("expected 8, 3 - got %v, %v", x, y)
	}
}

func TestFindLowestManhattanDistance(t *testing.T) {
	expected := Grid{}
	expected.grid = [][]string{
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", "+", "-", "-", "-", "-", "-", "+", ".", ".", "."},
		[]string{".", "|", ".", ".", ".", ".", ".", "|", ".", ".", "."},
		[]string{".", "|", ".", ".", "+", "-", "-", "x", "-", "+", "."},
		[]string{".", "|", ".", ".", "|", ".", ".", "|", ".", "|", "."},
		[]string{".", "|", ".", "-", "x", "-", "-", "+", ".", "|", "."},
		[]string{".", "|", ".", ".", "|", ".", ".", ".", ".", "|", "."},
		[]string{".", "|", ".", ".", ".", ".", ".", ".", ".", "|", "."},
		[]string{".", "o", "-", "-", "-", "-", "-", "-", "-", "+", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}
	if res := expected.FindLowestManhattanDistance(); res != 6 {
		t.Errorf("expected 6 got %v", res)
	}
}

func TestDrawAndFindLowestManhattan(t *testing.T) {
	grid := NewGrid(1000, 1000)
	grid.grid[500][500] = "o"
	path1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	path2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	grid1 := AddWire(grid, path1)
	grid2 := AddWire(grid1, path2)
	if res := grid2.FindLowestManhattanDistance(); res != 159 {
		t.Errorf("failed 1st man: %v", res)
	}

	gridb := NewGrid(1000, 1000)
	gridb.grid[500][500] = "o"
	pathb1 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	pathb2 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
	gridb1 := AddWire(gridb, pathb1)
	gridb2 := AddWire(gridb1, pathb2)
	if res := gridb2.FindLowestManhattanDistance(); res != 135 {
		t.Errorf("failed 2nd man: %v", res)
	}
}
