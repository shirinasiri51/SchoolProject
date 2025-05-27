package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Generate a random number between 1 and 10
    var num int = rand.Intn(10) + 1

    fmt.Println("Random number:", num)
}
