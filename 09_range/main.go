package main

import "fmt"

func main() {
  ids := []int{33, 76, 23, 67, 33, 3}

  // Loop through ids
  for i, id := range ids {
    fmt.Printf("%d - ID: %d\n", i, id)
  }

  // Use _ if we are not using index
  for _, id := range ids {
    fmt.Printf("ID: %d\n", id)
  }

  // Sum up all ids
  sum := 0
  for _, id := range ids {
    sum += id
  }
  fmt.Println("Sum", sum)

  // Range with map
  emails := map[string]string{"Ali": "ali@example.com", "Abu": "abu@example.com"}
  for k, v := range emails {
    fmt.Printf("%s: %s\n", k, v)
  }

  // Getting just the keys
  for k := range emails {
    fmt.Println("Name: " + k)
  }
}
