package main

import (
	"testing"
)

func Compare(arr1 []string, arr2 []string) bool {
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

func TestExtractLayers(t *testing.T) {
	input0 := "111111222222"
	expected0 := []string{"111111", "222222"}
	if res := ExtractLayers(input0, 3, 2); !Compare(res, expected0) {
		t.Errorf("expectation failed, received %v\n", res)
	}
	input1 := "3333333344444444"
	expected1 := []string{"33333333", "44444444"}
	if res := ExtractLayers(input1, 2, 4); !Compare(res, expected1) {
		t.Errorf("expectation failed, received %v\n", res)
	}
	input2 := "1234"
	expected2 := []string{"1", "2", "3", "4"}
	if res := ExtractLayers(input2, 1, 1); !Compare(res, expected2) {
		t.Errorf("expectation failed, received %v\n", res)
	}
}

func TestComputeNumberOf(t *testing.T) {
	input0 := "000"
	if res := ComputeNumberOf(input0, "0"); res != 3 {
		t.Errorf("expected 3, received %v\n", res)
	}
	input1 := "1233"
	if res := ComputeNumberOf(input1, "3"); res != 2 {
		t.Errorf("expected 3, received %v\n", res)
	}
	input2 := "2109314934932492409324"
	if res := ComputeNumberOf(input2, "2"); res != 4 {
		t.Errorf("expected 4, received %v\n", res)
	}
}

func TestFindLowestCount(t *testing.T) {
	input0 := []string{"000000", "101010", "100111"}
	if res := FindLowestCount(input0); res != 2 {
		t.Errorf("expected 2, received %v", res)
	}
	input1 := []string{"110022000", "222121210", "22100111"}
	if res := FindLowestCount(input1); res != 1 {
		t.Errorf("expected 1, received %v", res)
	}
}

func TestComputeHash(t *testing.T) {
	input0 := "000000101010100111"
	if res := ComputeHash(input0, 3, 2); res != 0 {
		t.Errorf("expected 0, received %v", res)
	}
	input1 := "110022000022212121002210011100"
	if res := ComputeHash(input1, 5, 2); res != 15 {
		t.Errorf("expected 15, received %v", res)
	}
}
