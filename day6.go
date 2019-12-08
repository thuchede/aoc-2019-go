package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFileDay6() {
	file, _ := os.Open("./input/day6.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// tree := make(map[string][]string)
	tree := make(map[string]string)

	for scanner.Scan() {
		orbit := scanner.Text()
		relation := strings.Split(orbit, ")")

		// tree[relation[0]] = append(tree[relation[0]], relation[1])
		tree[relation[1]] = relation[0]
	}

	// fmt.Printf("%v", ComputeLength(tree))
	fmt.Printf("%v\n", ComputeLengthToAnotherOrbit(tree, "YOU", "SAN"))
}

func ComputeLength(tree map[string][]string) int {
	return ComputeLengthFromItem(tree, "COM", 0)
}

func ComputeLengthFromItem(tree map[string][]string, item string, acc int) int {
	if len(tree[item]) == 0 {
		return acc
	}
	total := acc
	for _, sitem := range tree[item] {
		prev := ComputeLengthFromItem(tree, sitem, acc+1)
		total += prev
	}
	return total
}

func contains(list []string, key string) bool {
	for _, item := range list {
		if item == key {
			return true
		}
	}
	return false
}

func pos(list []string, key string) int {
	for num, item := range list {
		if item == key {
			return num
		}
	}
	return -1
}

func ComputeLengthToAnotherOrbit(tree map[string]string, s1 string, s2 string) int {
	currentNode := tree[s1]
	var listOfOrbits []string
	for currentNode != "" {
		listOfOrbits = append(listOfOrbits, currentNode)
		currentNode = tree[currentNode]
	}
	// each node index now maps to number of steps
	commonNode := tree[s2]
	steps := 0
	for !contains(listOfOrbits, commonNode) {
		commonNode = tree[commonNode]
		steps++
	}
	return pos(listOfOrbits, commonNode) + steps
}
