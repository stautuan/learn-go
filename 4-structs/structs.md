# Structs in Go

It represents structured data (fields) to group different variables together

<!-- is this the equivalent of a class? -->

```go
type car struct {
    make string
    model string
    doors int
    mileage int
}
```

## Nested structs

It can be nested to represent more complex properties

```go
type car struct {
    make string
    model string
    doors int
    mileage int

    // nested struct
    frontWheel wheel
    backWheel wheel
}

type wheel struct {
    radius int
    material string
}
```

These fields can be accessed using the dot `.` operator.

```go
myCar := car{}
myCar := frontWheel.radius = 5
```

## Anonymous structs

- They are defined with NO NAME, therefore it cannot be referenced elsewhere in the code.
- A good reason to use them is if it is only being used once.
- They prevent you from re-using a struct definition you never intended to re-use

```go
myCar := struct {
    make string
    model string
} {
    make: "tesla",
    model: "model 3",
}
```

Nesting anonymous structs

```go
type care struct {
    make string
    model string
    doors int
    mileage int

    // wheel is a field containing an anonymous struct
    wheel struct {
        radius int
        material int
    }
}
```

## Embedded structs

Go is not object-oriented, it doesn't support classes or inheritance, but embedded structs are a way to elevate and share fields between struct definitions.

<!-- embedded struct is an equivalent of super() constructor? -->

```go
type car struct {
    make string
    model string
}

type truck struct {
    // "car" is embedded, so the definition of a
    // "truck" now also contains all of the
    // fields of the car struct
    car
    bedSize int
}
```

```go
// lanesStruck is an instance of truck
lanesTruck := truck{
    bedSize: 10,
    car: car{
        make: "toyota",
        model: "camry",
    }
}

fmt.Println(lanesTruck.bedSize)

// embedded fields promoted to the top-level
// instead of lanesTruck.car.make
fmt.Println(lanesTruck.make)
fmt.Println(lanesTruck.model)
```

## Struct methods

- Like other oop languages, Go also supports methods that can be defined on structs.
- They are just functions that have a reciever, a special parameter that goes _before_ the name of the function.

```go
type rect struct {
    width int
    height int
}

// method of the rect struct
// it has a receiver (r rect)
func (r rect) area() int {
    return r.width * r.height
}

var r = rect{
    width: 5,
    heigh: 10,
}

fmt.Println(r.area())
// prints 50
```

### What is a receiver?

- a special parameter in a function
- by convention, the first letter of the struct's name is used

## Memory Layout

- structs are stored in a container that holds different pieces of data (fields) with a specific type
- its fields are stored in memory in one continuous block, without any gaps between them
- this makes accessing data fast because they are laid out in an organized way
- the fields are also placed in order as they are defined
- so ordering matters (from largest to smallest) if you want to save memory

### ✅ Good

```go
type stats struct {
    Reach    uint16
    NumPosts uint8
    NumLikes uint8
}
```

### ❌ Bad

```go
type stats struct {
    NumPosts uint8
    Reach    uint16
    NumLikes uint8
}
// wasted space
```

## Empty struct

They are the smallest type in go; they take up **zero bytes of memory**.

```go
// anonymous empty struct type
empty := struct{}{}

// named empty struct type
type emptyStruct struct{}
empty := emptyStruct{}
```
