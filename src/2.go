package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42)
	randomNumber := rand.Intn(10)
	fmt.Println("The random number is:", randomNumber)
}
