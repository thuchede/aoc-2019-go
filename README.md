# AOC 2019

## What ?

Advent of Code is an Advent calendar of small programming puzzles. I'm attempting to solve all puzzles in go

## Why?

I found out AOC would be a nice opportunity to learn and write go code by solving challenges out of my confort zone.

## How?

## When?

Since the main goal is to learn a new programming language, I'm not trying to reach the leaderboard, but to have a basic understanding at how to write go code and exercice my algorithm solving aptitude.

# Learnings

## Day 1

- recusivity rocks
- testing in go is easy (for now)

### What I've learned

`os.Open("./path/to/file.txt")` : to open a file
`bufio.NewScanner(file)` : to create a scanner and read a file

## Day 2

- do not prepare for tricky situation too much
  - input might not go out of bound, don't check too early
  - checking all error code in these puzzle might be a loss of time

### What I've learned

- there isn't much "built-in" helper method in go (like compare array, min or max for integer)
- `fmt.Printf("%v", value)` : to print formatted value

## Day 3

### first attempt

- creating struct to encapsulate behavior is useless
- do not try to reproduce example as it, use them as test case scenario only
- try to understand what is expected, input might be large and might drive possible solutions (or rule out some)
- initializing array boring

### What I've learned

`some_string[i:]` : to get a slice from index `i` to the end

### second attempt

- working with map is really a nice way to track unbounded values
- seems like easily extensible to part two problem to keep track of number of step needed to reach each location
- since there is no generics in go, there is no helper method to sort/find element in a array/map and you have to rewrite everything

### What I've learned

`for key, value := range my_map` : allow to loop through all values of a map

## Day 4

- brute force is sometime the easy way

### What I've learned

`num / int(math.Pow(10, float64(n))) % 10` : get the nth digit of num without converting to string

## Day 5

### What I've learned

The included testing package in go allow to define example and read output in stdin. If the function `MyFunc` should output `blabla`, then you can write the following test example:

```
func ExampleMyFunc(){
  MyFunc()
  // Output
  // blabla
}
```

## Day 6

- The lack of generics feels cumbersome, always forcing us to rewrite small utilities like `contains` and `indexOf`
- reversing a map hierarchy node => childrens for node => parent can help solving a problem

## Day 7
