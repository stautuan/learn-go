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
