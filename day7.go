package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileDay7() {
	file, _ := os.Open("./input/day7.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	strcode := strings.Split(scanner.Text(), ",")

	intcode := make([]int, len(strcode))

	for i := 0; i < len(strcode); i++ {
		intcode[i], _ = strconv.Atoi(strcode[i])
	}
	// fmt.Printf("maxout: %v\n", ComputePermutation(intcode))
	fmt.Printf("maxout: %v\n", ComputePermutationFeedBack(intcode))
}

// ExecuteIntCodeFromIndexWithModeAndIO execute the program stored in array starting from index
func ExecuteIntCodeFromIndexWithModeAndIO(array []int, index int, inputs []int, output int) int {
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
		return output
	} else if opcode == 1 {
		nextIndex = 4
		array[param3] = array[param1] + array[param2]
	} else if opcode == 2 {
		nextIndex = 4
		array[param3] = array[param1] * array[param2]
	} else if opcode == 3 {
		nextIndex = 2
		array[param1] = inputs[0]
		return ExecuteIntCodeFromIndexWithModeAndIO(array, index+nextIndex, inputs[1:], array[param1])
	} else if opcode == 4 {
		nextIndex = 2
		return ExecuteIntCodeFromIndexWithModeAndIO(array, index+nextIndex, inputs, array[param1])
	} else if opcode == 5 {
		if array[param1] != 0 {
			return ExecuteIntCodeFromIndexWithModeAndIO(array, array[param2], inputs, output)
		}
		nextIndex = 3
	} else if opcode == 6 {
		if array[param1] == 0 {
			return ExecuteIntCodeFromIndexWithModeAndIO(array, array[param2], inputs, output)
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
	} else {
		log.Fatalf("unknow opcode %v", opcode)
	}
	return ExecuteIntCodeFromIndexWithModeAndIO(array, index+nextIndex, inputs, output)
}

func RunAmplifierCode(amplifierSoftware []int, phases []int, input int) int {
	programA := amplifierSoftware
	programB := amplifierSoftware
	programC := amplifierSoftware
	programD := amplifierSoftware
	programE := amplifierSoftware
	ampAOutput := ExecuteIntCodeFromIndexWithModeAndIO(programA, 0, []int{phases[0], input}, -1)
	ampBOutput := ExecuteIntCodeFromIndexWithModeAndIO(programB, 0, []int{phases[1], ampAOutput}, -1)
	ampCOutput := ExecuteIntCodeFromIndexWithModeAndIO(programC, 0, []int{phases[2], ampBOutput}, -1)
	ampDOutput := ExecuteIntCodeFromIndexWithModeAndIO(programD, 0, []int{phases[3], ampCOutput}, -1)
	ampEOutput := ExecuteIntCodeFromIndexWithModeAndIO(programE, 0, []int{phases[4], ampDOutput}, -1)
	return ampEOutput
}

func ComputePermutation(program []int) int {
	arr := []int{0, 1, 2, 3, 4}
	max := 0
	for _, phases := range permutations(arr) {
		if res := RunAmplifierCode(program, phases, 0); res > max {
			fmt.Printf("phases %v - %v\n", phases, res)
			max = res
		}
	}
	return max
}

func ComputePermutationFeedBack(program []int) int {
	arr := []int{5, 6, 7, 8, 9}
	max := 0
	for _, phases := range permutations(arr) {
		if res := RunAmplifierCodeFBL(program, phases, 0); res > max {
			max = res
		}
	}
	return max
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func RunAmplifierCodeFBL(amplifierSoftware []int, phases []int, input int) int {
	programA := make([]int, len(amplifierSoftware))
	copy(programA, amplifierSoftware)
	programB := make([]int, len(amplifierSoftware))
	copy(programB, amplifierSoftware)
	programC := make([]int, len(amplifierSoftware))
	copy(programC, amplifierSoftware)
	programD := make([]int, len(amplifierSoftware))
	copy(programD, amplifierSoftware)
	programE := make([]int, len(amplifierSoftware))
	copy(programE, amplifierSoftware)
	ampAOutput := []int{phases[1]}
	ampBOutput := []int{phases[2]}
	ampCOutput := []int{phases[3]}
	ampDOutput := []int{phases[4]}
	ampEOutput := []int{phases[0], input}

	indx := []int{0, 0, 0, 0, 0}
	loop := true

	for loop == true {
		if indx[0] != 0 {
			if indx[0] == -1 {
				indx[0] = 0
			}
		}
		indx[0] = ExecuteIntCodeFromIndexWithModeAndIOFBL(programA, indx[0], &ampEOutput, &ampAOutput)
		if indx[1] != 0 {
			if indx[1] == -1 {
				indx[1] = 0
			}
		}
		indx[1] = ExecuteIntCodeFromIndexWithModeAndIOFBL(programB, indx[1], &ampAOutput, &ampBOutput)
		if indx[2] != 0 {
			if indx[2] == -1 {
				indx[2] = 0
			}
		}
		indx[2] = ExecuteIntCodeFromIndexWithModeAndIOFBL(programC, indx[2], &ampBOutput, &ampCOutput)
		if indx[3] != 0 {
			if indx[3] == -1 {
				indx[3] = 0
			}
		}
		indx[3] = ExecuteIntCodeFromIndexWithModeAndIOFBL(programD, indx[3], &ampCOutput, &ampDOutput)
		if indx[4] != 0 {
			if indx[4] == -1 {
				indx[4] = 0
			}
		}
		indx[4] = ExecuteIntCodeFromIndexWithModeAndIOFBL(programE, indx[4], &ampDOutput, &ampEOutput)
		if programE[indx[4]] == 99 {
			loop = false
		}
	}
	return ampEOutput[len(ampEOutput)-1]
}

func ExecuteIntCodeFromIndexWithModeAndIOFBL(array []int, index int, inputs *[]int, outputs *[]int) int {
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
		return ExecuteIntCodeFromIndexWithModeAndIOFBL(array, index+nextIndex, inputs, outputs)
	} else if opcode == 4 {
		nextIndex = 2
		*outputs = append(*outputs, array[param1])
		return ExecuteIntCodeFromIndexWithModeAndIOFBL(array, index+nextIndex, inputs, outputs)
	} else if opcode == 5 {
		if array[param1] != 0 {
			return ExecuteIntCodeFromIndexWithModeAndIOFBL(array, array[param2], inputs, outputs)
		}
		nextIndex = 3
	} else if opcode == 6 {
		if array[param1] == 0 {
			return ExecuteIntCodeFromIndexWithModeAndIOFBL(array, array[param2], inputs, outputs)
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
	} else {
		log.Fatalf("unknow opcode %v", opcode)
	}
	return ExecuteIntCodeFromIndexWithModeAndIOFBL(array, index+nextIndex, inputs, outputs)
}
