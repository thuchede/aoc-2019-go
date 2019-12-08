package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func compareSliceHelperWithMode(arr1 []int, arr2 []int) bool {
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

func TestExecuteIntCodeFromIndexWithMode(t *testing.T) {
	input1 := []int{1, 0, 0, 0, 99}
	output1 := []int{2, 0, 0, 0, 99}
	res1 := ExecuteIntCodeFromIndexWithMode(input1, 0)
	if !compareSliceHelperWithMode(res1, output1) {
		t.Error("expected array to match")
	}
	input2 := []int{2, 3, 0, 3, 99}
	output2 := []int{2, 3, 0, 6, 99}
	res2 := ExecuteIntCodeFromIndexWithMode(input2, 0)
	if !compareSliceHelperWithMode(res2, output2) {
		t.Error("expected array to match")
	}
	input3 := []int{99}
	output3 := []int{99}
	res3 := ExecuteIntCodeFromIndexWithMode(input3, 0)
	if !compareSliceHelperWithMode(res3, output3) {
		t.Error("expected array to match")
	}
	input4 := []int{2, 4, 4, 5, 99, 0}
	output4 := []int{2, 4, 4, 5, 99, 9801}
	res4 := ExecuteIntCodeFromIndexWithMode(input4, 0)
	if !compareSliceHelperWithMode(res4, output4) {
		t.Error("expected array to match")
	}
	input5 := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	output5 := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	res5 := ExecuteIntCodeFromIndexWithMode(input5, 0)
	if !compareSliceHelperWithMode(res5, output5) {
		t.Errorf("expected array to match, %v", res5)
	}
}

func TestExecuteIntCodeWithMode(t *testing.T) {
	input1 := []int{1, 0, 0, 0, 99}
	output1 := []int{2, 0, 0, 0, 99}
	res1 := ExecuteIntCodeWithMode(input1)
	if !compareSliceHelperWithMode(res1, output1) {
		t.Error("expected array to match")
	}
	input2 := []int{2, 3, 0, 3, 99}
	output2 := []int{2, 3, 0, 6, 99}
	res2 := ExecuteIntCodeWithMode(input2)
	if !compareSliceHelperWithMode(res2, output2) {
		t.Error("expected array to match")
	}
	input3 := []int{2, 4, 4, 5, 99, 0}
	output3 := []int{2, 4, 4, 5, 99, 9801}
	res3 := ExecuteIntCodeWithMode(input3)
	if !compareSliceHelperWithMode(res3, output3) {
		t.Error("expected array to match")
	}
	input4 := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	output4 := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	res4 := ExecuteIntCodeWithMode(input4)
	if !compareSliceHelperWithMode(res4, output4) {
		t.Error("expected array to match")
	}
	input := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	output := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
	res := ExecuteIntCodeWithMode(input)
	if !compareSliceHelperWithMode(res, output) {
		t.Error("expected value")
	}
	inputf := []int{1002, 4, 3, 4, 33}
	outputf := []int{1002, 4, 3, 4, 99}
	resf := ExecuteIntCodeWithMode(inputf)
	if !compareSliceHelperWithMode(resf, outputf) {
		t.Error("expected value")
	}
}

func TestExtractOpCodeAndMode(t *testing.T) {
	if o, p1, p2, p3 := ExtractOpCodeAndMode(1); o != 1 && p1 != 0 && p2 != 0 && p3 != 0 {
		t.Error("failed")
	}
	if o, p1, p2, p3 := ExtractOpCodeAndMode(2); o != 2 && p1 != 0 && p2 != 0 && p3 != 0 {
		t.Error("failed")
	}
	if o, p1, p2, p3 := ExtractOpCodeAndMode(99); o != 99 && p1 != 0 && p2 != 0 && p3 != 0 {
		t.Error("failed")
	}
	if o, p1, p2, p3 := ExtractOpCodeAndMode(101); o != 1 && p1 != 1 && p2 != 0 && p3 != 0 {
		t.Error("failed")
	}
	if o, p1, p2, p3 := ExtractOpCodeAndMode(1002); o != 2 && p1 != 0 && p2 != 1 && p3 != 0 {
		t.Error("failed")
	}
	if o, p1, p2, p3 := ExtractOpCodeAndMode(10199); o != 2 && p1 != 1 && p2 != 0 && p3 != 1 {
		t.Error("failed")
	}
}

func ExampleExecuteIntCodeFromIndexWithMode() {
	content := []byte("8\n")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile

	input1 := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	ExecuteIntCodeFromIndexWithMode(input1, 0)
	// Output:
	// input?
	// code: 1
}
func ExampleExecuteIntCodeFromIndexWithMode2() {
	content := []byte("9\n")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile

	input1 := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	ExecuteIntCodeFromIndexWithMode(input1, 0)
	// Output:
	// input?
	// code: 0
}

func ExampleExecuteIntCodeFromIndexWithMode3() {
	content := []byte("1\n")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile

	input1 := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	ExecuteIntCodeFromIndexWithMode(input1, 0)
	// Output:
	// input?
	// code: 1
}
func ExampleExecuteIntCodeFromIndexWithMode4() {
	content := []byte("9\n")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile

	input1 := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	ExecuteIntCodeFromIndexWithMode(input1, 0)
	// Output:
	// input?
	// code: 0
}
