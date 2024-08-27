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

The opposite way - you read them left to right, which is more clear and intuitive, like you would read in English.

```go
x int
p *int
a [33]int
```

## Passing Variables by Value

When variables in Go are passed into a function, that function receives a **copy** of that variable.

```go
func main() {
    x := 5
    increment(x)

    fmt.Println(x)
    // still prints 5, which is the x that is defined above
    // the increment function received a copy of x
    // it did not mutate the caller's original data
}

func increment(x) {
    x++
}
```

## Ignoring Return Values

- A function can return a value that the caller doesn't care about.
- We can ignore variables using an underscore.
- The Go compiler will **throw an error** if you have unused variable declarations in your code.

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
- naked returns should only be for short and simple functions

## The Benefits of Named Returns

### Good for Documentation (and understanding)

- named returns are important in longer functions with many returns
- we know what the function is returning directly, no need for a comment

## Early Returns

Guard Clauses leverage the ability to `return` early from a function to make nested conditionals one-dimensional (instead of using if/else chains).

```go
// with error handling
func divide(divident, divisor int) (int, error) {
    if divisor == 0 {
        return 0, errors.New("Can't divide by zero")
    }
    return dividend/divisor, nil
}
```

## Functions as values

Also called as "first class" and "higher-order functions"

Function as a parameter to another function:

```go
func add(x, y int) int {
    return x + y
}

func mul(x, y int) int {
    return x * y
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
    firstSum := arithmetic(a, b)
    secondSum := arithmetic(firstSum, c)
    return secondSum
}

func main() {
    sum := aggregate(2, 3, 4, add)
    // sum is 9
    product := aggregate(2, 3, 4, mul)
    // product is 24
}
```

## Anonymous Functions

- True to its name, they simply have _no name_.
- Use: for functions that will only be used once or to create a quick closure.

```go
package main

import "fmt"

func printReports(intro, body, outro string) {
    printCostReport(func(message string) int {
        return 2 * len(message)
    }, intro)
    printCostReport(func(message string) int {
        return 3 * len(message)
	}, body)
    printCostReport(func(message string) int {
        return 4 * len(message)
	}, outro)

    // the function inside printCostReport is an anonymous function
    // it takes a string as a parameter and returns an integer.
}

func main() {
    printReports(
        "Welcome to the Hotel California",
        "Such a lovely place",
        "Plenty of room at the Hotel California",
    )
}

func printCostReport(costCalculator func(string) int, message string) {
    cost := costCalculator(message)
    fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
    fmt.Println()
}
```

## Defer

- It delays the execution of a function or code until the surrounding function finishes
- Use: to clean up resources that are no longer being used; close database connections, file handlers, and the like.

```go
func main() {
    defer fmt.Println("Goodbye!")
    fmt.Println("Hello!")
}
```

- The program runs and `Hello` is printed.
- When everything inside the function has been executed, `Goodbye` will be printed.

## Block Scope

Variables declared inside a function cannot be used outside the function.

```go
package main

// scoped to the entire "main" package (basically global)
var age = 19

func sendEmail() {
    // scoped to the "sendEmail" function
    name := "Jon Snow"

    for i := 0; i < 5; i++ {
        // scoped to the "for" body
        email := "snow@winterfell.net"
    }
}
```
