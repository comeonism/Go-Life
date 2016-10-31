package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"time"
)

const size int = 23

var grid [size][size]bool
var nextGrid [size][size]bool
var cycleNo int = 1

func main() {
	initialize()

	printGrid()

	// game loop that refreshes every second
	timer := time.Tick(1000 * time.Millisecond)
	for tick := range timer {
		_ = tick
		cycleNo++

		evolve()
		grid = nextGrid

		printGrid()
	}
}

//////////////////////////////////////
// Initialization-related functions
//////////////////////////////////////

func initialize() {
	gridFileContents := fileToString("grid.txt")
	stringToGridArray(gridFileContents)
}

func fileToString(path string) string {
	dat, err := ioutil.ReadFile(path)
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func stringToGridArray(gridFileContents string) {
	validCharacterCounter := 0

	for _, character := range gridFileContents {

		if string(character) == "0" || string(character) == "1" {
			initCell(validCharacterCounter/size, int(math.Mod(float64(validCharacterCounter), float64(size))), string(character) == "1")
		} else {
			validCharacterCounter--
		}

		validCharacterCounter++
	}
}

func initCell(x, y int, alive bool) {
	grid[x][y] = alive
	nextGrid[x][y] = alive // for the initialization nextGrid == grid
}

//////////////////////////////////////
// Grid evolution-related functions
//////////////////////////////////////

func evolve() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			aliveNeighbours := getNoOfAliveNeighbours(i, j)
			evolveCell(aliveNeighbours, &grid[i][j], &nextGrid[i][j])
		}
	}
}

func getNoOfAliveNeighbours(x, y int) int {
	aliveNeighbours := 0

	countAlives(getLeft(&x), getUp(&y), &aliveNeighbours)
	countAlives(getLeft(&x), y, &aliveNeighbours)
	countAlives(getLeft(&x), getDown(&y), &aliveNeighbours)
	countAlives(x, getUp(&y), &aliveNeighbours)
	countAlives(x, getDown(&y), &aliveNeighbours)
	countAlives(getRight(&x), getUp(&y), &aliveNeighbours)
	countAlives(getRight(&x), y, &aliveNeighbours)
	countAlives(getRight(&x), getDown(&y), &aliveNeighbours)

	return aliveNeighbours
}

func countAlives(x, y int, aliveNeighbours *int) {
	accountForEgdes(&x, &y)
	checkAliveAndIncrement(&x, &y, aliveNeighbours)
}

func accountForEgdes(x, y *int) {
	if *x < 0 {
		*x = size - 1
	} else if *x >= size {
		*x = 0
	}

	if *y < 0 {
		*y = size - 1
	} else if *y >= size {
		*y = 0
	}
}

func checkAliveAndIncrement(x, y, aliveNeighbours *int) {
	if grid[*x][*y] == true {
		*aliveNeighbours++
	}
}

//////////////////////////////////////
//  Neighbour getting-related function
//////////////////////////////////////

func getLeft(x *int) int {
	return *x - 1
}

func getRight(x *int) int {
	return *x + 1
}

func getUp(y *int) int {
	return *y - 1
}

func getDown(y *int) int {
	return *y + 1
}

func evolveCell(numberOfAliveNeighbours int, cellAlive, nextGridCell *bool) {
	if *cellAlive && numberOfAliveNeighbours < 2 { // underpopulation
		*nextGridCell = false
	} else if !*cellAlive && numberOfAliveNeighbours == 3 { // reproduction
		*nextGridCell = true
	} else if *cellAlive && numberOfAliveNeighbours > 3 { // overpopulation
		*nextGridCell = false
	} else if *cellAlive && (numberOfAliveNeighbours == 3 || numberOfAliveNeighbours == 2) { // propagation
		*nextGridCell = true // not necessary and is only added for completeness sake
	}
}

//////////////////////////////////////
// Current implementation of output
//////////////////////////////////////

func printGrid() {
	fmt.Println("===========", cycleNo, "===========")

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] == true {
				fmt.Print(1)
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println("")
	}
}
