# Interfaces

- Interfaces are collection of **method** signatures.
- They are implemented _implicitly_. A type never declares that it implements a given interface.
- A type can implement multiple interfaces too!

```go
type Shape interface {
    area() float64
    perimeter() float64
}

type rect struct {
    width, height float64
}
// a type implements an interface if it has methods
// that match the interface's method signature
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}
```

## Name Your Interface Arguments

```go
type Copier interface {
    Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
```

It is not required to name the arguments of an interface in order for your code to compile properly BUT it's good to have for **readability and clarity**.

## Type assertions

```go
type shape interface {
    area() float64
}

type circle struct {
    radius float64
}

// s is an instance of the shape interface
// we want to check if s is a circle in order to cast it into its type
// c is the new circle struct cast from s
// ok is true if s is indeed a circle
c, ok := s.(circle)
if !ok {
    // log an error if s isn't a circle
    log.Fatal("s is not a circle")
}

radius := c.radius
```

## Type switches

A type switch is similar to regular switches, but the cases specify _types_ instead of _values_.

```go
func printNumericValue(num interface{}) {
    switch v:= num.(type) {
        case int:
        fmt.Printf("%T\n", v)
        case string:
        fmt.Printf("%T\n", v)
        default:
        fmt.Printf("%T\n", v)
    }
}

func main() {
    printNumericValue(1)
    // prints "int"
    printNumericValue("1")
    // prints "string"
    printNumericValue(struct{}{})
    // prints "struct {}"
}
```

## Guidelines on writing clean interfaces

### 1. Keep interfaces small

- Very important
- Interfaces are meant to define the _minimal_ behaviour to represent an idea or concept.

```go
type File interface {
    io.Close
    io.Reader
    io.Seeker
    Readdir(count int) ([]os.FileInfo, error)
    Stat() (os.FileInfo, error)
}
```

### 2. Interfaces should have no knowledge of satisfying types

```go
type car interface {
    Color() string
    Speed() int
    IsFiretruck() bool
}
```

- This is an interface that describes the components NECESSARY to define a car.
- `Color()` and `Speed()` makes sense, but `IsFiretruck()` don't, because there are different types of a car. The list would go on and on...

We should rely on the functionality of type assertion.

```go
type firetruck interface {
    car
    HoseLength() int
}
// it inherits all the methods from car interface and
// adds one additional method to make the car a firetruck
```

### 3. Interfaces are NOT classes

- They are more lightweight than classes.
- They don't have contructors or deconstructors. They don't create or destroy data. They just describe what functions should exist.
- They aren't hierarchical by nature (parent-child relationship).
- They only say what functions are required, NOT how they should work. Even if multiple types use the same interface, each type has to implement its own version of the function.
