package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPresentsDelivered(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2}, // some very lucky children
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, deliverPresents(tc.input, false), fmt.Sprintf("santa should have delivered %d presents", tc.want))
	}
}

func TestPresentsDeliveredWithRoboSanta(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, deliverPresents(tc.input, true), fmt.Sprintf("santa and robo-santa should have delivered %d presents", tc.want))
	}
}
