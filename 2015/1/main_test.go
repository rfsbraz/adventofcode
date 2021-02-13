package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloors(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, floorFinder(tc.input), fmt.Sprintf("floor should be %d", tc.want))
	}
}

func TestCrossFloors(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{")", 1},
		{"()())", 5},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, crossFloorFinder(tc.input, -1), fmt.Sprintf("position should be %d", tc.want))
	}
}
