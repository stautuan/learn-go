# Loops

```go
for i := 0; i < 10; i++ {
    fmt.Println(1)
}
// Prints 0 through 9
```

## Omitting conditions from a for loop

Loops in Go can omit sections of a for loop.

```go
for INITIAL; ; AFTER {
    // do something
}
```

Therefore, there is no `while` loop in Go. It is just a for loop that only has a CONDITION.

```go
for CONDITION {
    // do something
}
```

for example:

```go
plantHeight := 1
for plantHeight < 5 {
    fmt.Println("still growing! current height:", plantHeight)
    plantHeight++
}
fmt.Println("plant has grown to ", plantHeight, "inches")
```

## Continue & Break

`continue`: stops the current iteration of a loop and continues to the next iteration

```go
for i := 0;  i < 10; i++ {
    if i % 2 == 0 {
        continue
    }
    fmt.Println(i)
}
// 1
// 3
// 5
// 7
// 9
```

`break`: stops the current iteration of a loop and exits the loop

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
// 0
// 1
// 2
// 3
// 4
```
