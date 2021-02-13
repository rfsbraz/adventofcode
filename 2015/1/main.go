package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	fmt.Printf("Went into basement at position %d\n", floorFinder(input))

	fmt.Printf("Destination floor is %d\n", crossFloorFinder(input, -1))
}

func floorFinder(instructions string) int {
	floor := 0

	for _, letter := range instructions {
		// An opening parenthesis, (, means he should go up one floor, and a closing parenthesis, ), means he should go down one floor.
		floor += stairStep(letter)
	}
	// To what floor do the instructions take Santa?
	return floor
}

func crossFloorFinder(instructions string, enterLevel int) int {
	floor := 0

	for i, letter := range instructions {
		floor += stairStep(letter)

		// What is the position of the character that causes Santa to first enter the basement?
		if floor == enterLevel {
			return i + 1
		}
	}

	return 0
}

func stairStep(char rune) int {
	if char == '(' {
		return 1
	} else {
		return -1
	}
}
