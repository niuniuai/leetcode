package main

import (
	"fmt"
	"strings"
)

func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	tokens := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		factorial[i] = factorial[i-1] * i
		tokens[i] = i
	}

	k--
	var b strings.Builder
	for n > 0 {
		n--
		a := k / factorial[n]
		k = k % factorial[n]
		fmt.Fprintf(&b, "%d", tokens[a]+1)
		tokens = append(tokens[:a], tokens[a+1:]...)
	}
	return b.String()
}

func main() {
	fmt.Println(getPermutation(9,1))
}
