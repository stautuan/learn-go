# The Error Interface

```go
type error interface {
	Error() string
}
```

If a function might go wrong, it should return an error as the last result. When you call a function, it should check if the error is `nil` (no error, success), and handle if it is not.
