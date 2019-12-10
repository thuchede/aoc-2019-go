package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFileDay8() {
	file, _ := os.Open("./input/day8.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	rawLayers := scanner.Text()

	// fmt.Printf("%v\n", rawLayers)
	fmt.Printf("%v\n", ComputeHash(rawLayers, 25, 6))
}

func ExtractLayers(input string, width, height int) []string {
	res := []string{}
	for len(input) > 0 {
		res = append(res, input[:width*height])
		input = input[width*height:]
	}
	return res
}

func ComputeNumberOf(input string, c string) int {
	return strings.Count(input, c)
}

func FindLowestCount(inputs []string) int {
	min := int(^uint(0) >> 1)
	idx := -1

	for i, input := range inputs {
		if res := ComputeNumberOf(input, "0"); res < min {
			min = res
			idx = i
		}
	}
	return idx
}

func ComputeHash(input string, width, height int) int {
	layers := ExtractLayers(input, width, height)
	index := FindLowestCount(layers)
	validationLayer := layers[index]
	output := strings.Count(validationLayer, "1") * strings.Count(validationLayer, "2")
	return output
}
