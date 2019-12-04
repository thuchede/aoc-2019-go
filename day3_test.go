package main

import (
	"testing"
)

func TestEqual(t *testing.T) {
	g1 := make(map[Key][]int)
	g2 := make(map[Key][]int)
	if !equal(g1, g2) {
		t.Error("expected empty panel to have same content")
	}
	g3 := make(map[Key][]int)
	g4 := make(map[Key][]int)
	g3[Key{X: 1, Y: 2}] = []int{1}
	if equal(g3, g4) {
		t.Error("expected different sized panel to not have same content")
	}
	g5 := make(map[Key][]int)
	g6 := make(map[Key][]int)
	g5[Key{X: 1, Y: 2}] = []int{1}
	g3[Key{X: 1, Y: 2}] = []int{1, 2}
	if equal(g5, g6) {
		t.Error("expected grid with different content (value size) to not have equal content")
	}
	g7 := make(map[Key][]int)
	g8 := make(map[Key][]int)
	g7[Key{X: 1, Y: 2}] = []int{1}
	g8[Key{X: 1, Y: 2}] = []int{2}
	if equal(g7, g8) {
		t.Error("expected grid with different content (pure content) to not have equal content")
	}
	g9 := make(map[Key][]int)
	g10 := make(map[Key][]int)
	g9[Key{X: 1, Y: 2}] = []int{1}
	g10[Key{X: 2, Y: 2}] = []int{1}
	if equal(g9, g10) {
		t.Error("expected grid with different content (keys) to not have equal content")
	}
}

func TestAddWire(t *testing.T) {
	grid0 := make(map[Key][]int)
	expected0 := make(map[Key][]int)
	expected0[Key{X: 1, Y: 0}] = []int{1}
	expected0[Key{X: 2, Y: 0}] = []int{1}
	path0 := []string{"R2"}
	res0 := AddWire(grid0, path0, 1)
	if !equal(res0, expected0) {
		t.Errorf("expected grid to be the same 0 : %v, %v", res0, expected0)
	}
	grid1 := make(map[Key][]int)
	expected1 := make(map[Key][]int)
	expected1[Key{X: 1, Y: 0}] = []int{2}
	expected1[Key{X: 2, Y: 0}] = []int{2}
	expected1[Key{X: 2, Y: -1}] = []int{2}
	expected1[Key{X: 2, Y: -2}] = []int{2}
	path1 := []string{"R2", "D2"}
	res1 := AddWire(grid1, path1, 2)
	if !equal(res1, expected1) {
		t.Errorf("expected grid to be the same 1 : %v, %v", res1, expected1)
	}
	grid2 := make(map[Key][]int)
	expected2 := make(map[Key][]int)
	expected2[Key{X: 1, Y: 0}] = []int{1, 2}
	expected2[Key{X: 2, Y: 0}] = []int{1, 2}
	expected2[Key{X: 2, Y: -1}] = []int{1, 2}
	expected2[Key{X: 2, Y: -2}] = []int{1, 2}
	path21 := []string{"R2", "D2"}
	path22 := []string{"R2", "D2"}
	res21 := AddWire(grid2, path21, 1)
	res22 := AddWire(res21, path22, 2)
	if !equal(res22, expected2) {
		t.Errorf("expected grid to be the same 2 : \n%v\n%v\n", res22, expected2)
	}
}

func TestComputeManhattanDistance1(t *testing.T) {
	if ComputeManhattanDistance(-7, 2) != 9 {
		t.Error("expected ComputeManhattanDistance to be 9")
	}
	if ComputeManhattanDistance(0, 0) != 0 {
		t.Error("expected ComputeManhattanDistance to be 0")
	}
	if ComputeManhattanDistance(5, -3) != 8 {
		t.Error("expected ComputeManhattanDistance to be 8")
	}
	if ComputeManhattanDistance(10, 9) != 19 {
		t.Error("expected ComputeManhattanDistance to be 19")
	}
}

func TestFindLowestManhattanDistance(t *testing.T) {
	grid0 := make(map[Key][]int)
	path01 := []string{"R2", "D2"}
	path02 := []string{"D2", "R2"}
	res01 := AddWire(grid0, path01, 1)
	res02 := AddWire(res01, path02, 2)
	if res := FindLowestManhattanDistance(res02); res != 4 {
		t.Errorf("expected 4 got %v", res02)
	}
	grid1 := make(map[Key][]int)
	path11 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	path12 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	res11 := AddWire(grid1, path11, 1)
	res12 := AddWire(res11, path12, 2)
	if res := FindLowestManhattanDistance(res12); res != 159 {
		t.Errorf("expected 159 got %v", res12)
	}
	grid2 := make(map[Key][]int)
	path21 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	path22 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
	res21 := AddWire(grid2, path21, 1)
	res22 := AddWire(res21, path22, 2)
	if res := FindLowestManhattanDistance(res22); res != 135 {
		t.Errorf("expected 135 got %v", res22)
	}
}

func TestAddWireWithDistance(t *testing.T) {
	grid0 := make(map[Key][]Key)
	expected0 := make(map[Key][]Key)
	expected0[Key{X: 1, Y: 0}] = []Key{Key{X: 1, Y: 1}}
	expected0[Key{X: 2, Y: 0}] = []Key{Key{X: 1, Y: 2}}
	path0 := []string{"R2"}
	res0 := AddWireWithDistance(grid0, path0, 1)
	if !equalDist(res0, expected0) {
		t.Errorf("expected grid to be the same 0 : %v, %v", res0, expected0)
	}
}

func TestFindLowestPathDistance(t *testing.T) {
	grid0 := make(map[Key][]Key)
	path01 := []string{"R8", "U5", "L5", "D3"}
	path02 := []string{"U7", "R6", "D4", "L4"}
	res01 := AddWireWithDistance(grid0, path01, 1)
	res02 := AddWireWithDistance(res01, path02, 2)
	if res := FindLowestPathDistance(res02); res != 30 {
		t.Errorf("expected 30 got %v", res)
	}
	grid1 := make(map[Key][]Key)
	path11 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	path12 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	res11 := AddWireWithDistance(grid1, path11, 1)
	res12 := AddWireWithDistance(res11, path12, 2)
	if res := FindLowestPathDistance(res12); res != 610 {
		t.Errorf("expected 610 got %v", res)
	}
	grid2 := make(map[Key][]Key)
	path21 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	path22 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
	res21 := AddWireWithDistance(grid2, path21, 1)
	res22 := AddWireWithDistance(res21, path22, 2)
	if res := FindLowestPathDistance(res22); res != 410 {
		t.Errorf("expected 410 got %v", res)
	}
}
