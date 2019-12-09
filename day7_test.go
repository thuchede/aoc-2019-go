package main

import (
	"testing"
)

func TestExecuteIntCodeFromIndexWithModeAndReturnValue(t *testing.T) {
	amp := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	phases := []int{4, 3, 2, 1, 0}
	if res := RunAmplifierCode(amp, phases, 0); res != 43210 {
		t.Errorf("expected 43210 got %v", res)
	}
	amp1 := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
		101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	phases1 := []int{0, 1, 2, 3, 4}
	if res := RunAmplifierCode(amp1, phases1, 0); res != 54321 {
		t.Errorf("expected 54321 got %v", res)
	}
	amp2 := []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
		1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}
	phases2 := []int{1, 0, 4, 3, 2}
	if res := RunAmplifierCode(amp2, phases2, 0); res != 65210 {
		t.Errorf("expected 65210 got %v", res)
	}
	t.Logf("went great")
}