package main

func main() {
	g := NewGrid(11, 10)
	g.grid[9][1] = "o"
	g.printGrid()

	AddWire(g, []string{"R3", "U2"})
}
