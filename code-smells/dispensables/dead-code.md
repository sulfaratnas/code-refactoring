# Dead Code

A variable, parameter, field, method or class is no longer used (usually because it’s obsolete).

## What It Looks Like

```
func main() {
    fmt.Println("Hello, World!")

    // This is always false
    x := 2
    if x < 0 {
        fmt.Println("hello")
    }
}
```

## Why It Hurts

Dead code adds unnecessary complexity to the codebase, making it harder to understand and maintain.


## How To Fix It

[Remove Dead Code](.././../refactorings/remove-dead-code.md).

