package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	count := 10

	channelToSquare := make(chan float64, count)
	resultChannel := make(chan float64, count)

	go createSlice(count, channelToSquare)
	go square(count, channelToSquare, resultChannel)

	for i := 0; i < count; i++ {
		fmt.Print(<-resultChannel, " ")
	}
}

func square(count int, numChannel chan float64, channel chan float64) {
	for i := 0; i < count; i++ {
		channel <- math.Pow(<-numChannel, 2)
	}
}

// just to be sure

func createSlice(count int, ch chan float64) {
	for i := 0; i < count; i++ {
		ch <- math.Ceil(rand.Float64() * 100)
	}
}
