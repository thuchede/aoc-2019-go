package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadFileDay10() {
	file, _ := os.Open("./input/day10.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	asteroids := [][]string{}

	for scanner.Scan() {
		asteroids = append(asteroids, strings.Split(scanner.Text(), ""))
	}

	coor, res := FindHigest(asteroids)
	fmt.Printf("%v : %v\n", coor, res)
}

func CompareVectorDirection(a Key, b Key) bool {
	const delta = 0.000000001
	aNorm := math.Sqrt(float64(a.X*a.X + a.Y*a.Y))
	bNorm := math.Sqrt(float64(b.X*b.X + b.Y*b.Y))
	return 1.0-float64(a.X*b.X+a.Y*b.Y)/(aNorm*bNorm) < delta
}

func ComputeVectorFromCoordinates(a Key, b Key) Key {
	return Key{X: b.X - a.X, Y: b.Y - a.Y}
}

func CountVisibleAsteroid(input [][]string, coordinates Key) int {
	found := []Key{}
	res := 0
	for i := 0; i < len(input); i++ {
	next:
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "#" {
				if i == coordinates.Y && j == coordinates.X {
					continue
				}
				key := ComputeVectorFromCoordinates(coordinates, Key{X: j, Y: i})
				for _, vec := range found {
					if CompareVectorDirection(vec, key) {
						continue next
					}
				}
				found = append(found, key)
				res++
			}
		}
	}
	return res
}

func MapVisible(input [][]string) [][]string {
	res := make([][]string, len(input))
	for m := 0; m < len(input); m++ {
		res[m] = make([]string, len(input[m]))
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "#" {
				res[i][j] = strconv.Itoa(CountVisibleAsteroid(input, Key{X: j, Y: i}))
			} else {
				res[i][j] = "."
			}
		}
	}
	return res
}

func FindHigest(input [][]string) (Key, int) {
	max := 0
	var coordinates Key

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "#" {
				res := CountVisibleAsteroid(input, Key{X: j, Y: i})
				if res > max {
					max = res
					coordinates = Key{X: j, Y: i}
				}
			}
		}
	}
	return coordinates, max
}
