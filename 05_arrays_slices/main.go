package main

import "fmt"

func main() {
  // Declaring an array
  var fruitArr [2]string
  // Note: Array length must be specified

  // Assigning values to array
  fruitArr[0] = "Apple"
  fruitArr[1] = "Orange"

  fmt.Println(fruitArr)
  fmt.Println(fruitArr[1])

  // Declaring and assigning values at the same time
  fruits := [2]string{"Banana", "Watermelon"}
  fmt.Println(fruits)

  fruitSlice := []string{"Kiwi", "Lemon", "Durian", "Pear"}
  fmt.Println(fruitSlice)
  // Display array length
  fmt.Println(len(fruitSlice))
  // display range
  fmt.Println(fruitSlice[1:3])
}
