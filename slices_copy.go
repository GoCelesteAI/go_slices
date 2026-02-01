package main

import "fmt"

func main() {
  // Copy creates independent slice
  src := []int{1, 2, 3, 4, 5}
  dst := make([]int, len(src))

  copied := copy(dst, src)
  fmt.Println("Source:", src)
  fmt.Println("Destination:", dst)
  fmt.Println("Elements copied:", copied)

  // Modify destination - source unchanged
  dst[0] = 999
  fmt.Println("\nAfter modifying dst:")
  fmt.Println("Source:", src)
  fmt.Println("Destination:", dst)

  fmt.Println("\nPartial copy:")
  partial := make([]int, 3)
  copy(partial, src)
  fmt.Println("Partial (3 elements):", partial)

  fmt.Println("\nNil vs Empty:")
  var nilSlice []int
  emptySlice := []int{}
  madeSlice := make([]int, 0)

  fmt.Printf("nil slice: %v, is nil: %t\n", nilSlice, nilSlice == nil)
  fmt.Printf("empty literal: %v, is nil: %t\n", emptySlice, emptySlice == nil)
  fmt.Printf("make(0): %v, is nil: %t\n", madeSlice, madeSlice == nil)
}
