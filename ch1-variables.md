# Variables

## Printing text to the console

```go
package main

import "fmt"

func main() {
  fmt.Println("I am learning Go...")
}
```

## Declaring a variable

### The sad way 😢

```go
// initialize variable (it defaults to 0)
var mySkillIssues int

// overwrite the zero value
mySkillIssues = 42
```

### The GOATed way 😎

```go
mySkillIssues := 42
```

The walrus operator `:=` should be used whenever you can. It can't be used, however, outside a function. Moreover, Go infers that `mySkillIssues` is an `int` because of its value `42`. This is called **type inference**.

```go
package main

import "fmt"

func main() {
  messageStart := "Happy birthday! You are now"
  age := 21
  messageEnd := "years old!"

  fmt.Println(messageStart, age, messageEnd)

  // Happy birthday! You are now 21 years old!
}
```

## The Compilation Process

The code that we write in Go is not readable by the machine. Therefore, it needs to go through a compiler in order to convert it into machine language.

code -> compiler -> executable

```go
/* lets the compiler know that we want this code to compile
and run as a standalone program, as opposed to being a library */
package main

/* imports the fmt package from the standard library to allow
us to use the Println method to print to the console */
import "fmt"

/* the main function, entry point for a Go progam */
func main() {
  // code here
}
```

## 🏃💨 Fast and Compiled

Go is one of the fastest programming languages. Faster than interpreted languages like JavaScript and Python, but not as fast as low-level languages like as Rust, Zig or C.

However, it _compiles_ faster than all of them, which makes the developer experience super productive.
