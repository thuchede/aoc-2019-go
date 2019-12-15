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

- generating all permutation (n!) is hard
- the concept of non-shared memory was not properly tested, resulting in really hard-to-debug assumption
- use pointer when you need to share a resizeable array/slice between differents objects

### What I've learned

`array_1 := array_0` creates a copy of the array, ie. 2 distinct objects
`slice_1 := slice_0` creates a pointer to the slice, ie. 1 object

## Day 8

### What I've learned

`string` in go is an immutable type, ie. you can't set :

```
my_string[index] = some_char // invalid!
```

`int(^uint(0) >> 1)` gets you the maximum signed integer without the math package
`math.MaxInt32` is max value on 32 bits
`math.MaxInt64` is max value on 64 bits

## Day 9

`The computer's available memory should be much larger than the initial program` force to initialize an large array. In go you'd have to loop through the array to set each initial value individually

## Day 10

To check all, immediately, visible items from a set of coordinates, we have to find a way to check if the line of sight to an item is blocked.
From a point of observation with coordinates O(a, b) the closest point C(x, y) is visible if the vector OC is the shortest in that direction and all points Vi(xi, yi) are hidden if the dot product OC.OVi = cos COVi = 1
Can't check for colinear vector as opposite vector will be colinear.

## What I've learned

- when comparing float number, use a very small delta
- trigonometry can be hard
- manipulating vectors and origin is troublesome
- `math.Atan2` is a useful function to get a angle between two vectors
- `delete(myMap, entry)` : delete a entry from a map
