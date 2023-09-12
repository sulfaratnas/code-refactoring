# Remove Dead Code

## Example

### Before

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

### After

```
func main() {
    fmt.Println("Hello, World!")
}
```