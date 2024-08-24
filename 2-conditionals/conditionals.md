## The initial state of an IF block

An `if`conditional can have an **initial** statement.

```go
if length := getLength(email); length < 1 {
  fmt.Println("Email is invalid")
}
```

### Why would you use this?

1. It's shorter
2. It limits the scope of the initialized variable to the `if` block.

## Switch

- They are similar to if-else statement but more concise and readable.
- a `break` statement is not required; it is implicit in Go.
- use a `fallthrough` to go through the next case.
- a `default` case runs if there is no match.

```go
func getCreator(os string) string {
    var creator string
    switch os {
    case "linux":
        creator = "Linus Torvalds"
    case "windows":
        creator = "Bill Gates"

    // all three of these cases will set creator to "A Steve"
    case "Mac OS":
        fallthrough
    case "Mac OS X":
        fallthrough
    case "mac":
        creator = "A Steve"

    default:
        creator = "Unknown"
    }
    return creator
}
```
