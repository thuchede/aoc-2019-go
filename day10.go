package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	CreateDestructionMap(asteroids, coor)
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

func CosTheta(a Key, b Key) float64 {
	aNorm := math.Sqrt(float64(a.X*a.X + a.Y*a.Y))
	bNorm := math.Sqrt(float64(b.X*b.X + b.Y*b.Y))
	return float64(a.X*b.X+a.Y*b.Y) / (aNorm * bNorm)
}

func OrientedCosTheta(a Key, b Key) float64 {
	aNorm := math.Sqrt(float64(a.X*a.X + a.Y*a.Y))
	bNorm := math.Sqrt(float64(b.X*b.X + b.Y*b.Y))

	x := -b.X
	sign := float64(x) / math.Abs(float64(x))
	return sign * float64(a.X*b.X+a.Y*b.Y) / (aNorm * bNorm)
}

type byMagnitude []Key

func (s byMagnitude) Len() int {
	return len(s)
}
func (s byMagnitude) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byMagnitude) Less(i, j int) bool {
	a, b := ComputeVectorFromCoordinates(Key{X: 20, Y: 20}, s[i]), ComputeVectorFromCoordinates(Key{X: 11, Y: 13}, s[j])
	aNorm := math.Sqrt(float64(a.X*a.X + a.Y*a.Y))
	bNorm := math.Sqrt(float64(b.X*b.X + b.Y*b.Y))
	return aNorm < bNorm
}

func CreateDestructionMap(input [][]string, station Key) {

	vectors := make(map[float64][]Key)
	sortedKey := []float64{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "#" {
				if i == station.Y && j == station.X {
					continue
				}
				pos := Key{X: j, Y: i}
				item := ComputeVectorFromCoordinates(station, pos)

				theta := math.Pi/2 - math.Atan2(-float64(item.Y), float64(item.X))
				if theta < 0 {
					theta = theta + 2*math.Pi
				}
				_, found := vectors[theta]
				vectors[theta] = append(vectors[theta], pos)
				if !found {
					sortedKey = append(sortedKey, theta)
				}
			}
		}
	}

	sort.Float64s(sortedKey)

	for k, v := range vectors {
		sort.Sort(byMagnitude(v))
		vectors[k] = v
	}

	destroyed := []Key{}

	killcount := 0
	for {
		idx := 0
		for idx < len(sortedKey) {
			theta := sortedKey[idx]

			killcount++
			destroyed = append(destroyed, vectors[theta][0])

			vectors[theta] = vectors[theta][1:]
			if len(vectors[theta]) == 0 {
				delete(vectors, theta)
				sortedKey = append(sortedKey[:idx], sortedKey[idx+1:]...)
				continue
			}
			idx++
		}
		if len(vectors) == 0 {
			break
		}
	}
	fmt.Printf("200th => %v\n", destroyed[199])
}
