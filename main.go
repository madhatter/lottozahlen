package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func generateNumbers(n int, l int) []int {
	winningNumbers := make([]int, n)
	usedNumbers := make(map[int]struct{})

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	//for len(winningNumbers) < n {
	for i := range winningNumbers {
		var num int
		for {
			num = r1.Intn(l)

			if _, ok := usedNumbers[num]; !ok {
				break
			}
		}
		usedNumbers[num] = struct{}{}
		winningNumbers[i] = num
	}

	return winningNumbers
}

func main() {
	arg, _ := strconv.Atoi(os.Args[1])
	log.SetFormatter(&log.JSONFormatter{})
	fmt.Printf("Set of %v.\n", arg)
	for i := 1; i < arg+1; i++ {
		fmt.Println(generateNumbers(5, 50))
	}
}
