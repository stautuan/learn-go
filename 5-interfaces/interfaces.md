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
