package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileDay9() {
	file, _ := os.Open("./input/day9.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	strcode := strings.Split(scanner.Text(), ",")

	intcode := make([]int, 10000)

	for i := range intcode {
		intcode[i] = 0
	}

	for i := 0; i < len(strcode); i++ {
		intcode[i], _ = strconv.Atoi(strcode[i])
	}
	input := []int{1}
	output := []int{}
	res := ExecuteIntCodeRelative(intcode, 0, &input, &output, 0)
	fmt.Printf("%v : %v\n", output, res)
}

func ExecuteIntCodeRelative(array []int, index int, inputs *[]int, outputs *[]int, relative int) int {
	opcode, p1, p2, p3 := ExtractOpCodeAndMode(array[index])

	var nextIndex int

	var param1 int

	var param2 int

	var param3 int

	if opcode < 99 {
		if p1 == 1 {
			param1 = index + 1
		} else if p1 == 2 {
			param1 = relative + array[index+1]
		} else {
			param1 = array[index+1]
		}
	}
	if opcode != 3 && opcode != 4 && opcode != 9 && opcode != 99 {
		if p2 == 1 {
			param2 = index + 2
		} else if p2 == 2 {
			param2 = relative + array[index+2]
		} else {
			param2 = array[index+2]
		}
	}
	if opcode < 3 || (opcode > 6 && opcode < 99) {
		if p3 == 1 {
			param3 = index + 3
		} else if p3 == 2 {
			param3 = relative + array[index+3]
		} else {
			param3 = array[index+3]
		}
	}
	if opcode == 99 {
		return index
	} else if opcode == 1 {
		nextIndex = 4
		array[param3] = array[param1] + array[param2]
	} else if opcode == 2 {
		nextIndex = 4
		array[param3] = array[param1] * array[param2]
	} else if opcode == 3 {
		if len(*inputs) == 0 {
			return index
		}
		nextIndex = 2
		array[param1] = (*inputs)[0]
		*inputs = (*inputs)[1:]
		return ExecuteIntCodeRelative(array, index+nextIndex, inputs, outputs, relative)
	} else if opcode == 4 {
		nextIndex = 2
		*outputs = append(*outputs, array[param1])
		return ExecuteIntCodeRelative(array, index+nextIndex, inputs, outputs, relative)
	} else if opcode == 5 {
		if array[param1] != 0 {
			return ExecuteIntCodeRelative(array, array[param2], inputs, outputs, relative)
		}
		nextIndex = 3
	} else if opcode == 6 {
		if array[param1] == 0 {
			return ExecuteIntCodeRelative(array, array[param2], inputs, outputs, relative)
		}
		nextIndex = 3
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
	} else if opcode == 9 {
		nextIndex = 2
		fmt.Printf("relative base adjusted from %v with %v : %v\n", relative, param1, array[param1])
		relative = relative + array[param1]
		fmt.Printf("relative is now %v\n", relative)
	} else {
		log.Fatalf("unknow opcode %v", opcode)
	}
	return ExecuteIntCodeRelative(array, index+nextIndex, inputs, outputs, relative)
}
