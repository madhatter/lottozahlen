package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func createSlice(max int) []int {
	min := 1
	arr := make([]int, max)

	for i := range arr {
		arr[i] = min + i
	}
	return arr
}

func generateNumbers(n int, l int) []int {
	numbrs := createSlice(l)
	rand.Shuffle(len(numbrs), func(i, j int) { numbrs[i], numbrs[j] = numbrs[j], numbrs[i] })

	return numbrs[0:n]
}

func main() {
	arg, _ := strconv.Atoi(os.Args[1])
	log.SetFormatter(&log.JSONFormatter{})
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Set of %v.\n", arg)
	for i := 1; i < arg+1; i++ {
		fmt.Println(generateNumbers(5, 50))
	}
}
