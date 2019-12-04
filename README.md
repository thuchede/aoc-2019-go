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

- creating struct to encapsulate behavior is useless
- do not try to reproduce example as it, use them as test case scenario only
- try to understand what is expected, input might be large and might drive possible solutions (or rule out some)
- initializing array boring

### What I've learned

`some_string[i:]` : to get a slice from index `i` to the end
