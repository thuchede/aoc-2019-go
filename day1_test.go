package main

import (
	"os"
	"testing"
)

func TestComputeFuelForMassRecursively(t *testing.T) {
	if ComputeFuelForMassRecursively(0, 14) != 2 {
		t.Error("expected 2")
	}
	if ComputeFuelForMassRecursively(0, 1969) != 966 {
		t.Error("expected 966")
	}
	if ComputeFuelForMassRecursively(0, 100756) != 50346 {
		t.Error("expected 50346")
	}
}

func TestComputeFuelForMass(t *testing.T) {
	if ComputeFuelForMass(-3) != 0 {
		t.Error("expected 0")
	}
	if ComputeFuelForMass(0) != 0 {
		t.Error("expected 0")
	}
	if ComputeFuelForMass(6) != 0 {
		t.Error("expected 0")
	}
	if ComputeFuelForMass(9) != 1 {
		t.Error("expected 1")
	}
	if ComputeFuelForMass(12) != 2 {
		t.Error("expected 2")
	}
	if ComputeFuelForMass(15) != 3 {
		t.Error("expected 3")
	}
}
func TestComputeFuelForAllModule(t *testing.T) {
	file1, _ := os.Open("./input_test/day1.txt")

	if res, _ := ComputeFuelForAllModule(file1); res != 40 {
		t.Errorf("expected 40, got %d", res)
	}
}
