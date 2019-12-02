package main

import (
	"testing"
)

func compareSliceHelper(arr1 []int, arr2 []int) bool {
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

func TestExecuteIntCodeFromIndex(t *testing.T) {
	input0 := []int{}
	_, err0 := ExecuteIntCodeFromIndex(input0, 0)
	if err0 == nil {
		t.Error("expected error")
	}
	input1 := []int{1, 0, 0, 0, 99}
	output1 := []int{2, 0, 0, 0, 99}
	res1, err1 := ExecuteIntCodeFromIndex(input1, 0)
	if !compareSliceHelper(res1, output1) || err1 != nil {
		t.Error("expected array to match")
	}
	input2 := []int{2, 3, 0, 3, 99}
	output2 := []int{2, 3, 0, 6, 99}
	res2, err2 := ExecuteIntCodeFromIndex(input2, 0)
	if !compareSliceHelper(res2, output2) || err2 != nil {
		t.Error("expected array to match")
	}
	input3 := []int{99}
	output3 := []int{99}
	res3, err3 := ExecuteIntCodeFromIndex(input3, 0)
	if !compareSliceHelper(res3, output3) || err3 != nil {
		t.Error("expected array to match")
	}
	input4 := []int{2, 4, 4, 5, 99, 0}
	output4 := []int{2, 4, 4, 5, 99, 9801}
	res4, err4 := ExecuteIntCodeFromIndex(input4, 0)
	if !compareSliceHelper(res4, output4) || err4 != nil {
		t.Error("expected array to match")
	}
	input5 := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	output5 := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	res5, err5 := ExecuteIntCodeFromIndex(input5, 0)
	if !compareSliceHelper(res5, output5) || err5 != nil {
		t.Errorf("expected array to match, %v", res5)
	}
}

func TestExecuteIntCode(t *testing.T) {
	input0 := []int{}
	_, err0 := ExecuteIntCode(input0)
	if err0 == nil {
		t.Error("expected error")
	}
	input1 := []int{1, 0, 0, 0, 99}
	output1 := []int{2, 0, 0, 0, 99}
	res1, err1 := ExecuteIntCode(input1)
	if !compareSliceHelper(res1, output1) || err1 != nil {
		t.Error("expected array to match")
	}
	input2 := []int{2, 3, 0, 3, 99}
	output2 := []int{2, 3, 0, 6, 99}
	res2, err2 := ExecuteIntCode(input2)
	if !compareSliceHelper(res2, output2) || err2 != nil {
		t.Error("expected array to match")
	}
	input3 := []int{2, 4, 4, 5, 99, 0}
	output3 := []int{2, 4, 4, 5, 99, 9801}
	res3, err3 := ExecuteIntCode(input3)
	if !compareSliceHelper(res3, output3) || err3 != nil {
		t.Error("expected array to match")
	}
	input4 := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	output4 := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	res4, err4 := ExecuteIntCode(input4)
	if !compareSliceHelper(res4, output4) || err4 != nil {
		t.Error("expected array to match")
	}
}

func TestExecuteIntCodeFinal(t *testing.T) {
	input := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	output := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
	res, err := ExecuteIntCode(input)
	if !compareSliceHelper(res, output) || err != nil {
		t.Error("expected value")
	}
}
