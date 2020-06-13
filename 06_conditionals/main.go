package main

import "fmt"

func main() {
  x := 5
  y := 10

  // The common convention is to NOT use parenthesis

  // if...else
  if x <= y {
    fmt.Printf("%d is less than or equal to %d\n", x, y)
  } else {
    fmt.Printf("%d is less than %d\n", y, x)
  }

  // if...else if...else
  color := "red"
  if color == "green" {
    fmt.Println("Color is green")
  } else if color == "blue" {
    fmt.Println("Color is blue")
  } else {
    fmt.Println("Color is NOT green or blue")
  }

  // switch
  switch color {
  case "red":
    fmt.Println("Color is red")
  case "green":
    fmt.Println("Color is green")
  default:
    fmt.Println("Color is not red or green")
  }
}
