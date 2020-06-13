package main

import "fmt"

func main() {
  fmt.Println(greeting("Rashid"))
  fmt.Println(getSum(1, 2))
  fmt.Println(anotherGetSum(2, 2))
}

// Function declarations
func greeting(name string) string {
  return "Hello, " + name + "!"
}

func getSum(num1 int, num2 int) int {
  return num1 + num2
}

func anotherGetSum(num1, num2 int) int {
  return num1 + num2
}
