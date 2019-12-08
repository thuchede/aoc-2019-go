package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileDay5b() {
	input := []int{3, 0, 4, 0, 99}
	ExecuteIntCodeWithMode(input)
	fmt.Println("done")
}

func ReadFileDay5c() {
	input := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
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
	if opcode != 3 && opcode != 4 && opcode != 99 {
		if p2 == 1 {
			param2 = index + 2
		} else {
			param2 = array[index+2]
		}
	}
	if opcode < 3 || (opcode > 6 && opcode < 99) {
		if p3 == 1 {
			param3 = index + 3
		} else {
			param3 = array[index+3]
		}
	}
	if opcode == 99 {
		return array
	} else if opcode == 1 {
		nextIndex = 4
		array[param3] = array[param1] + array[param2]
	} else if opcode == 2 {
		nextIndex = 4
		array[param3] = array[param1] * array[param2]
	} else if opcode == 3 {
		nextIndex = 2
		fmt.Println("input?")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		ii, _ := strconv.Atoi(input[:len(input)-1])
		array[param1] = ii
	} else if opcode == 4 {
		nextIndex = 2
		fmt.Printf("code: %v\n", array[param1])
	} else if opcode == 5 {
		if array[param1] != 0 {
			return ExecuteIntCodeFromIndexWithMode(array, array[param2])
		} else {
			nextIndex = 3
		}
	} else if opcode == 6 {
		if array[param1] == 0 {
			return ExecuteIntCodeFromIndexWithMode(array, array[param2])
		} else {
			nextIndex = 3
		}
	} else if opcode == 7 {
		nextIndex = 4
		if array[param1] < array[param2] {
			array[param3] = 1
		} else {
			array[param3] = 0
		}
	} else if opcode == 8 {
		nextIndex = 4
		if array[param1] == array[param2] {
			array[param3] = 1
		} else {
			array[param3] = 0
		}
	} else {
		log.Fatalf("unknow opcode %v", opcode)
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
