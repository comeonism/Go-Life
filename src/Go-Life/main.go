package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

const size int = 3

var grid [size][size]Cell

type Cell struct {
	x, y  int
	alive bool
}

func main() {
	initialize()
}

func initialize() {
	gridFileContents := fileToString("grid.txt")
	stringToGridArray(gridFileContents)

	fmt.Printf("%v", grid)
}

func stringToGridArray(gridFileContents string) {
	validCharacterCounter := 0

	for _, character := range gridFileContents {

		if string(character) == "0" || string(character) == "1" {
			processCell(validCharacterCounter/size, int(math.Mod(float64(validCharacterCounter), float64(size))), string(character) == "1")
		} else {
			validCharacterCounter--
		}

		validCharacterCounter++
	}
}

func processCell(x, y int, alive bool) {
	grid[x][y] = Cell{x, y, alive}
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
