package main

import (
	"bufio"
	"fmt"
	"os"
)

type Box struct {
	length, width, height int
}

// find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l.
func (b Box) surfaceArea() int {
	// 2(ab + ac + bc)
	return 2 * (b.length*b.width + b.width*b.height + b.height*b.length)
}

// The elves also need a little extra paper for each present: the area of the smallest side.
func (b Box) smallestSide() int {
	min := b.length * b.width
	for _, v := range []int{b.width * b.height, b.height * b.length} {
		if v < min {
			min = v
		}
	}
	return min
}

// The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face
func (b Box) smallestSidePerimeter() int {
	min := 2*b.length + 2*b.width
	for _, v := range []int{2*b.width + 2*b.height, 2*b.height + 2*b.length} {
		if v < min {
			min = v
		}
	}
	return min
}

// Each present also requires a bow made out of ribbon as well
// the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present
func (b Box) volume() int {
	return b.length * b.width * b.height
}

func (b Box) paperToOrder() int {
	return b.surfaceArea() + b.smallestSide()
}

func (b Box) ribbonToOrder() int {
	return b.smallestSidePerimeter() + b.volume()
}

func boxFromInput(input string) Box {
	var box Box

	fmt.Sscanf(input, "%dx%dx%d", &box.length, &box.width, &box.height)

	return box
}

func main() {
	wrappingPaperRequired, ribbonRequired := 0, 0

	scanner := bufio.NewScanner(os.Stdin)

	// All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
	for scanner.Scan() {
		box := boxFromInput(scanner.Text())
		wrappingPaperRequired += box.paperToOrder()
		ribbonRequired += box.ribbonToOrder()
	}

	fmt.Printf("Elves need to order %d feet of paper and %d feet of ribbon\n", wrappingPaperRequired, ribbonRequired)
}
