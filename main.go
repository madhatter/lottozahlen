package main

import (
	"fmt"
	"math/rand"
	"time"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	var winningNumbers []int

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for len(winningNumbers) < 5 {
		num := 0
		for num == 0 && !contains(winningNumbers, num) {
			num = r1.Intn(50)
			winningNumbers = append(winningNumbers, num)
		}
	}
	fmt.Println(winningNumbers)
}
