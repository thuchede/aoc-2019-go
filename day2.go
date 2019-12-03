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

// ExecDay2 print part1 solution
func ExecDay2() {
	res, err := ReadFileDay2(12, 2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res[0])
}

// ReadFileDay2 reads a file and prints all its value
func ReadFileDay2(noun int, verb int) ([]int, error) {
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ExecuteIntCodeWithNounAndVerb(intcode, noun, verb)
}

// ReadFileDay2Part2 reads a file and prints all its value
func ReadFileDay2Part2() {
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			res, _ := ReadFileDay2(i, j)
			if res[0] == 19690720 {
				fmt.Printf("noun: %v, verb: %v", i, j)
				return
			}
		}
	}
}

// ExecuteIntCodeWithNounAndVerb return the executed intcode program with specific verb and noun
func ExecuteIntCodeWithNounAndVerb(intcode []int, noun int, verb int) ([]int, error) {
	intcode[1] = noun
	intcode[2] = verb

	return ExecuteIntCode(intcode)
}
