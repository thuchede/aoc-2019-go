package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFileDay5b() {
	input := []int{3, 0, 4, 0, 99}
	ExecuteIntCodeWithMode(input)
	fmt.Println("done")
}

func ReadFileDay5() {
	file, _ := os.Open("./input/day5.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	strcode := strings.Split(scanner.Text(), ",")

	intcode := make([]int, len(strcode))

	for i := 0; i < len(strcode); i++ {
		intcode[i], _ = strconv.Atoi(strcode[i])
	}

	ExecuteIntCodeWithMode(intcode)

	// fmt.Printf("%v", intcode)
}

// ExecuteIntCodeWithMode execute the program stored in array
func ExecuteIntCodeWithMode(array []int) []int {
	return ExecuteIntCodeFromIndexWithMode(array, 0)
}

// ExecuteIntCodeFromIndexWithMode execute the program stored in array starting from index
func ExecuteIntCodeFromIndexWithMode(array []int, index int) []int {
	opcode, p1, p2, p3 := ExtractOpCodeAndMode(array[index])
	if opcode == 99 {
		return array
	}
	var nextIndex int

	var param1 int

	var param2 int

	var param3 int

	if opcode < 99 {
		if p1 == 1 {
			param1 = index + 1
		} else {
			param1 = array[index+1]
		}
	}
	if opcode < 3 {
		if p2 == 1 {
			param2 = index + 2
		} else {
			param2 = array[index+2]
		}
		if p3 == 1 {
			param3 = index + 3
		} else {
			param3 = array[index+3]
		}
	}

	if opcode == 1 {
		nextIndex = 4
		array[param3] = array[param1] + array[param2]
	}
	if opcode == 2 {
		nextIndex = 4
		array[param3] = array[param1] * array[param2]
	}
	if opcode == 3 {
		nextIndex = 2
		fmt.Println("input?")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		ii, _ := strconv.Atoi(input[:len(input)-1])
		array[param1] = ii
	}
	if opcode == 4 {
		nextIndex = 2
		fmt.Printf("code : %v\n", array[param1])
	}
	return ExecuteIntCodeFromIndexWithMode(array, index+nextIndex)
}

// ExtractOpCodeAndMode return the current opcode and all its param modes
func ExtractOpCodeAndMode(num int) (int, int, int, int) {
	opcode := num % 100
	p1mode := num / 100 % 10
	p2mode := num / 1000 % 10
	p3mode := num / 10000 % 10
	return opcode, p1mode, p2mode, p3mode
}
