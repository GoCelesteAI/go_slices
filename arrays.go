package main

import "fmt"

func main() {
  // Declare an array with explicit size
  var numbers [5]int
  fmt.Println("Zero-valued array:", numbers)

  // Assign values
  numbers[0] = 10
  numbers[1] = 20
  numbers[2] = 30
  fmt.Println("After assignment:", numbers)

  // Array literal
  fruits := [3]string{"apple", "banana", "cherry"}
  fmt.Println("Fruits:", fruits)

  // Array with inferred size
  colors := [...]string{"red", "green", "blue", "yellow"}
  fmt.Println("Colors:", colors)
  fmt.Println("Length:", len(colors))

  // Arrays are values (copied on assignment)
  original := [3]int{1, 2, 3}
  copied := original
  copied[0] = 100
  fmt.Println("Original:", original)
  fmt.Println("Copied:", copied)
}
