package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range ints {
		if number%2 == 0 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}
	}
}
