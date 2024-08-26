# Functions

- The variable type comes after the variable name.
- And when arguments are of the same type, the type is declared after the last argument.

```go
func sub(x, y int) int {
  return x-y
}
```

`func sub(x, y int) int` this expression is called a "function signature".

## Declaration of Syntax

### C-Style syntax

The type goes on the left and the expression on the right.

```c
int y;
```

### Go-style syntax

Clear and intuitive, you read them left to right, like you would in English.

```go
x int
p *int
a [33]int
```

## Passing Variables by Value

- When variables in Go are passed into a function, that function receives a **copy** of that variable.

```go
func main() {
    x := 5
    increment(x)

    fmt.Println(x)
    // it will still print 5
    // the increment function received a copy of x
    // the function is unable to mutate the caller's original data
}

func increment(x) {
    x++
}
```

## Ignoring Return Values

- A function can return a value that the caller doesn't care about.
- We can ignore variables using an underscore.
- The Go compiler will \*_throw an error_ if you have unused variable declarations in your code.

```go
func getPoint() (x, y int) {
    return 3, 4
}

// ignore y value
x, _ := getPoint()
```

## Named Return Values

```go
func getCoords() (x, y int) {
    return
}
```

Note:

- `x` and `y` are the return values
- they are initialized with zero values
- we could simply write `return` to return the values of these two variables, rather than writing `return x, y`. This is known as a "naked" return.
