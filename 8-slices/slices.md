## Array

These are fixed-size groups of the same type, like "ice cubes" -- ice (type) in a tray (array) 😅

`[n]T` is an array of n values of type `T`

To declare an array of 10 integers:

```go
var myInts [10]int
```

To declare an initialized literal:

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
```

## Slices

Arrays are fixed in size. Once you make an array like `[10]int` you cannot add an 11th element.

The zero value of slice is `nil`.

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}
```

The syntax is `arrayname[lowIndex:highIndex]`, where `lowIndex` is inclusive and `highIndex` is exclusive. Both can also be omitted `arrayname[:]`.

## Make

We can create a new slice using the `make` function!

```go
// func make([]T, len, capacity) []T
mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)
```

When we make slices, it is filled with zero values.

Here is how we can fill it:

```go
mySlice := []string{"I", "love", "go"}

// notice that the square brackets didn't have a 3 in them
// if they did, it would be an array instead of a slice
```

- `len()`: is the current length of the slice
- `cap()`:is the maximum length of the slice
- `len()` and `cap()` return 0 when the slice is `nil`
