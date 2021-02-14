package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdventCoinMiner(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, mineAdventCoin(tc.input), fmt.Sprintf("Adventcoin leading number should be %d", tc.want))
	}
}
