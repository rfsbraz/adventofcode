package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

const AdventCoinDifficulty = 5

func mineAdventCoin(hash string) int {
	i := 0
	for {
		data := []byte(fmt.Sprintf("%s%d", hash, i))
		if strings.HasPrefix(fmt.Sprintf("%x", md5.Sum(data)), strings.Repeat("0", AdventCoinDifficulty)) {
			return i
		}
		i += 1
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		adventCoinNumber := mineAdventCoin(scanner.Text())
		fmt.Printf("Advent coin number for %s is %d! - %s%d\n", scanner.Text(), adventCoinNumber, scanner.Text(), adventCoinNumber)
	}
}
