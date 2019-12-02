package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// ComputeFuelForMassRecursively return fuel needed to move a module based on its mass plus an accumulator
func ComputeFuelForMassRecursively(accumulator int, mass int) int {
	res := mass/3 - 2
	if res < 0 {
		return accumulator
	}
	return ComputeFuelForMassRecursively(accumulator+res, res)
}

// ComputeFuelForMass return fuel needed to move a module based on its mass
func ComputeFuelForMass(mass int) int {
	return ComputeFuelForMassRecursively(0, mass)
}

// ComputeFuelForAllModule return fuel needed to move all modules in file
// each module is described by its respective mass
func ComputeFuelForAllModule(file *os.File) (int, error) {
	fuel := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		fuel = fuel + ComputeFuelForMass(mass)
	}

	return fuel, scanner.Err()
}

// ReadFile reads a file and prints all its value
func ReadFile() {
	file, err := os.Open("./input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fuel, err := ComputeFuelForAllModule(file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fuel)
}
