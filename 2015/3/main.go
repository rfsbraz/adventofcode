package main

import (
	"bufio"
	"fmt"
	"os"
)

func initializeHouseGrid(len int) [][]int {
	grid := make([][]int, len)
	for i := range grid {
		grid[i] = make([]int, len)
	}
	return grid
}

// north (^), south (v), east (>), or west (<)
func processDirection(x int, y int, r rune) (int, int) {
	switch r {
	case '^':
		return x, y + 1
	case 'v':
		return x, y - 1
	case '<':
		return x - 1, y
	case '>':
		return x + 1, y
	default:
		return x, y
	}
}

func countPresents(grid [][]int) int {
	count := 0
	for _, i := range grid {
		for _, j := range i {
			if j > 0 {
				count += 1
			}
		}
	}
	return count
}

// Santa is delivering presents to an infinite two-dimensional grid of houses.
func deliverPresents(input string, withRoboSanta bool) int {
	// counting max chars of a given direction would optimize for memory, but this is good enough
	size := len(input)

	// create a two dimensional cartesian space
	grid := initializeHouseGrid((size + 1) * 2)

	// start at the center
	x, y, i, j := size, size, size, size

	// First house gets a gift! Adding robo-santa gift is not required, the house is still counted
	grid[x][y] += 1
	for p, cardinal := range input {
		if p%2 == 0 || !withRoboSanta {
			x, y = processDirection(x, y, cardinal)
			grid[x][y] += 1
		} else {
			i, j = processDirection(i, j, cardinal)
			grid[i][j] += 1
		}
	}

	return countPresents(grid)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
	for scanner.Scan() {
		delivered := deliverPresents(scanner.Text(), false)
		deliveredWithRobot := deliverPresents(scanner.Text(), true)
		fmt.Printf("Upon following some very tipsy instructions, santa delivered %d presents\n", delivered)
		fmt.Printf("With his trusty Robo-Santa, %d gifts were delivered the next year!\n", deliveredWithRobot)
	}
}
