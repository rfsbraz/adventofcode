package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSurfaceAreas(t *testing.T) {
	testCases := []struct {
		box  Box
		want int
	}{
		{Box{2, 3, 4}, 52},
		{Box{1, 1, 10}, 42},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, tc.box.surfaceArea(), fmt.Sprintf("surface area should be %d", tc.want))
	}
}

func TestSmallestSides(t *testing.T) {
	testCases := []struct {
		box  Box
		want int
	}{
		{Box{2, 3, 4}, 6},
		{Box{10, 1, 1}, 1},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, tc.box.smallestSide(), fmt.Sprintf("smallest side should have area %d", tc.want))
	}
}

func TestPaperToOrder(t *testing.T) {
	testCases := []struct {
		box  Box
		want int
	}{
		{Box{2, 3, 4}, 58},
		{Box{10, 1, 1}, 43},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, tc.box.paperToOrder(), fmt.Sprintf("elves should order %d feet", tc.want))
	}
}

func TestSmallestPerimeter(t *testing.T) {
	testCases := []struct {
		box  Box
		want int
	}{
		{Box{2, 3, 4}, 10},
		{Box{1, 1, 10}, 4},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, tc.box.smallestSidePerimeter(), fmt.Sprintf("smallest perimeter should be %d", tc.want))
	}
}

func TestVolume(t *testing.T) {
	testCases := []struct {
		box  Box
		want int
	}{
		{Box{2, 3, 4}, 24},
		{Box{1, 1, 10}, 10},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, tc.box.volume(), fmt.Sprintf("volume should be %d", tc.want))
	}
}

func TestRibbonToORder(t *testing.T) {
	testCases := []struct {
		box  Box
		want int
	}{
		{Box{2, 3, 4}, 34},
		{Box{1, 1, 10}, 14},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, tc.box.ribbonToOrder(), fmt.Sprintf("elves should order %d feet", tc.want))
	}
}

func TestBoxParse(t *testing.T) {
	assert.Equal(t, Box{3, 11, 24}, boxFromInput("3x11x24"), "box should be 3x11x24")
}
