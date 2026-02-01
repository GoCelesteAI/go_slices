package main

import "fmt"

func main() {
  // Create a slice literal
  nums := []int{10, 20, 30, 40, 50}
  fmt.Println("Slice:", nums)

  // Length and capacity
  fmt.Println("Length:", len(nums))
  fmt.Println("Capacity:", cap(nums))

  // Accessing elements
  fmt.Println("First:", nums[0])
  fmt.Println("Last:", nums[len(nums)-1])

  fmt.Println("\nSlicing operations:")
  fmt.Println("nums[:3]  =", nums[:3])
  fmt.Println("nums[2:]  =", nums[2:])
  fmt.Println("nums[1:4] =", nums[1:4])

  // Slices are references
  slice1 := nums[1:4]
  slice1[0] = 999
  fmt.Println("\nAfter modifying slice1:")
  fmt.Println("slice1:", slice1)
  fmt.Println("nums:", nums)
}
