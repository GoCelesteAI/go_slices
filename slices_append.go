package main

import "fmt"

func main() {
  // Start with empty slice
  var nums []int
  fmt.Println("Empty slice:", nums)
  fmt.Println("Is nil?", nums == nil)

  // Append elements
  nums = append(nums, 1)
  nums = append(nums, 2, 3)
  fmt.Println("\nAfter append:", nums)

  // Append another slice
  more := []int{4, 5, 6}
  nums = append(nums, more...)
  fmt.Println("After append slice:", nums)

  fmt.Println("\nUsing make():")
  buffer := make([]int, 3, 10)
  fmt.Println("Buffer:", buffer)
  fmt.Printf("Length: %d, Capacity: %d\n", len(buffer), cap(buffer))

  // Fill the buffer
  buffer[0] = 100
  buffer[1] = 200
  buffer[2] = 300
  fmt.Println("Filled:", buffer)

  // Append beyond length
  buffer = append(buffer, 400, 500)
  fmt.Printf("After append: %v (len=%d, cap=%d)\n", buffer, len(buffer), cap(buffer))
}
