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

func generateNumbers(n int, l int) []int {
	var winningNumbers []int

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for len(winningNumbers) < n {
		num := 0
		for num == 0 && !contains(winningNumbers, num) {
			num = r1.Intn(l)
			winningNumbers = append(winningNumbers, num)
		}
	}

	return winningNumbers
}

func main() {
	for i := 1; i < 5; i++ {
		fiveNumbers := generateNumbers(5, 50)
		twoNumbers := generateNumbers(2, 10)
		fmt.Println(fiveNumbers, twoNumbers)
	}
}
