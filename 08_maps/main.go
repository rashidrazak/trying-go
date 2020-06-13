package main

import "fmt"

func main() {
  // Maps are key-value pairs

  // Defining a map
  emails := make(map[string]string)

  // Assigning key-value
  emails["Rashid"] = "rashid@example.com"
  emails["Shai"] = "shai@example.com"

  fmt.Println(emails)
  fmt.Println(len(emails))
  fmt.Println(emails["Rashid"])

  // Deleting from map
  delete(emails, "Shai")
  fmt.Println(emails)

  // Declaring map and add key-value at one go
  otherEmails := map[string]string{"Ali": "ali@example.com", "Abu": "abu@example.com"}
  // Let's add another one
  otherEmails["Ah Seng"] = "ahseng@example.com"
  fmt.Println(otherEmails)
  fmt.Println(len(otherEmails))
  fmt.Println(otherEmails["Abu"])
}
