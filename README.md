# Go Arrays & Slices

Code examples from **Go Tutorial for Beginners #9 - Arrays & Slices**

## Contents

- `arrays.go` - Fixed-size arrays and value semantics
- `slices_basic.go` - Slice literals, slicing, and references
- `slices_append.go` - Append and make functions
- `slices_copy.go` - Copy function and nil vs empty slices

## Running the Examples

```bash
go run arrays.go
go run slices_basic.go
go run slices_append.go
go run slices_copy.go
```

## Key Concepts

### Arrays (Fixed Size)
```go
var arr [5]int                    // Zero-valued
fruits := [3]string{"a", "b", "c"} // Literal
colors := [...]string{"r", "g", "b"} // Inferred size
```

### Slices (Dynamic)
```go
nums := []int{10, 20, 30, 40, 50}
fmt.Println(nums[:3])  // [10 20 30]
fmt.Println(nums[2:])  // [30 40 50]
```

### Append
```go
nums = append(nums, 60, 70)
nums = append(nums, more...) // Spread operator
```

### Make
```go
buffer := make([]int, 5, 10) // length 5, capacity 10
```

### Copy
```go
dst := make([]int, len(src))
copy(dst, src) // Independent copy
```

### Nil vs Empty
```go
var nilSlice []int      // nil, len 0
emptySlice := []int{}   // not nil, len 0
madeSlice := make([]int, 0) // not nil, len 0
```

## Watch the Tutorial

[Go Tutorial for Beginners #9 - Arrays & Slices](https://youtube.com/@CelesteAI)

## License

MIT
