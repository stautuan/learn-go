# Interfaces

- Interfaces are collection of method signatures.
- A type "implements" an interface if it has methods that match the interface's method signatures.
- They are implemented _implicitly_.
- A type never declares that it implements a given interface.
- A type can implement multiple interfaces.

```go
type Shape interface {
    area() float64
    perimeter() float64
}

type rect struct {
    width, height float64
}
// methods
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}
```

## Multiple Interfaces
