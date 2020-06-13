package main

import "fmt"

func main() {
  a := 3
  b := &a

  // Pointers will show the memory address where the variable reside
  fmt.Println(a, b)
  fmt.Printf("%T\n", a)
  fmt.Printf("%T\n", b)

  // Use * to read value from memory address
  fmt.Println(*b)
  // This does the same thing
  fmt.Println(*&a)

  // Changing value using pointer
  *b = 10
  fmt.Println(a)
}
