# The Error Interface

```go
type error interface {
	Error() string
}
```

If a function might go wrong, it should return an error as the last result. When you call a function, it should check if the error is `nil` (no error, success), and handle if it is not.

## Panic

`panic` is another way to handle errors in Go. When a function calls `panic`, the program crashes and prints a stack trace.

As a general rule, DO NOT USE PANIC!

`log.Fatal()` is a good alternative way to panic, if you want your program to cleanly exit in an unrecoverable way.
