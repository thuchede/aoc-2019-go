package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day2() {

}

// ExecuteIntCode execute the program stored in array
func ExecuteIntCode(array []int) ([]int, error) {
	return ExecuteIntCodeFromIndex(array, 0)
}

// ExecuteIntCodeFromIndex execute the program stored in array starting from index
func ExecuteIntCodeFromIndex(array []int, index int) ([]int, error) {
	if index > len(array)-1 {
		return nil, errors.New("Index out of bound")
	}
	opcode := array[index]

	if opcode == 1 {
		array[array[index+3]] = array[array[index+1]] + array[array[index+2]]
	}
	if opcode == 2 {
		array[array[index+3]] = array[array[index+1]] * array[array[index+2]]
	}
	if opcode == 99 {
		return array, nil
	}
	return ExecuteIntCodeFromIndex(array, index+4)
}

// ReadFileDay2 reads a file and prints all its value
func ReadFileDay2() {
	file, err := os.Open("./input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	strcode := strings.Split(scanner.Text(), ",")

	intcode := make([]int, len(strcode))

	for i := 0; i < len(strcode); i++ {
		intcode[i], _ = strconv.Atoi(strcode[i])
	}

	intcode[1] = 12
	intcode[2] = 2

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res, err := ExecuteIntCode(intcode)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res[0])
}
