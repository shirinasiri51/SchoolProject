  package main
  import (
    "math/rand"
    "time"
  )
  func main() {
    rand.Seed(time.Now().UnixNano())
    num := rand.Intn(10)
    if num > 5 {
      println("The number is greater than 5")
    } else {
      println("The number is lesser than 5")
    }
  }