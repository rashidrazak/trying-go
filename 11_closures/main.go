package main

import "fmt"

func main() {
  sum := adder()
  for i := 0; i < 10; i++ {
    fmt.Println(sum(i))
  }
}

// adder is a function that returns a function.
// This returned function takes an int and returns an int
func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}
