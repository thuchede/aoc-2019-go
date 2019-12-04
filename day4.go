package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadFileDay4() {
	file, _ := os.Open("./input/day4.txt")

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	arr := strings.Split(scanner.Text(), "-")
	start, _ := strconv.Atoi(arr[0])
	end, _ := strconv.Atoi(arr[1])

	// res := CountNumberOfPasswordBetween(start, end)
	res := CountStrictNumberOfPasswordBetween(start, end)
	fmt.Printf("%v\n", res)
}

// HasDoubleDigit return true if num has two adjacent digit that are identical
func HasDoubleDigit(num int) bool {
	i := 0
	previous := num / int(math.Pow(10, float64(i))) % 10
	for i < 5 {
		i++
		current := num / int(math.Pow(10, float64(i))) % 10
		if previous == current {
			return true
		}
		previous = current
	}
	return false
}

// HasDoubleDigitNotMore return true if num has two adjacent digit that are identical
func HasDoubleDigitNotMore(num int) bool {
	res := false
	i := 6
	group := 0
	previous := num / int(math.Pow(10, float64(i))) % 10
	for i > 0 {
		i--
		current := num / int(math.Pow(10, float64(i))) % 10
		if previous == current {
			group++
			if group < 2 {
				res = res || true
			} else {
				res = res && false
			}
		} else {
			if res {
				return res
			}
			group = 0
		}
		previous = current
	}
	return res
}

// HasIncreasingDigitOnly return true if the digit in num never decrease when read from left to right
func HasIncreasingDigitOnly(num int) bool {
	i := 0
	previous := num / int(math.Pow(10, float64(i))) % 10
	for i < 5 {
		i++
		current := num / int(math.Pow(10, float64(i))) % 10
		if previous < current {
			return false
		}
		previous = current
	}
	return true
}

// CountNumberOfPasswordBetween compute the number of password meeting criteria in a range
func CountNumberOfPasswordBetween(start, end int) int {
	total := 0
	for i := start; i <= end; i++ {
		if HasDoubleDigit(i) && HasIncreasingDigitOnly(i) {
			total++
		}
	}
	return total
}

// CountStrictNumberOfPasswordBetween compute the number of password meeting criteria in a range
func CountStrictNumberOfPasswordBetween(start, end int) int {
	total := 0
	for i := start; i <= end; i++ {
		if HasDoubleDigitNotMore(i) && HasIncreasingDigitOnly(i) {
			total++
		}
	}
	return total
}
