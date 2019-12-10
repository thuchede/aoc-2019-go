package main

import (
	"testing"
)

func CompareI(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func TestExecuteIntCodeRelative(t *testing.T) {
	intcode0 := make([]int, 10000)
	copy(intcode0, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99})
	expected0 := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	input0 := []int{}
	output0 := []int{}
	if ExecuteIntCodeRelative(intcode0, 0, &input0, &output0, 0); !CompareI(output0, expected0) {
		t.Errorf("expectation failed, received %v\n", output0)
	}
}
