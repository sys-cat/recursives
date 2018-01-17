package main

import (
	"fmt"
	"math/big"
)

func recursive_for(n int) int {
	res := 0
	for i := 0; i >= n; i++ {
		if i == 0 {
			res = 1
		}
		res = res * i
	}
	return res
}

func recursive(n *big.Int) *big.Int {
	if n == big.NewInt(0) {
		return big.NewInt(1)
	}
	return new(big.Int).Mul(n, recursive(new(big.Int).Sub(n, big.NewInt(1))))
}

func main() {
	fmt.Println(recursive_for(100))
	fmt.Println(recursive(big.NewInt(100)))
}
