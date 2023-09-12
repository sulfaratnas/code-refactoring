# Rename Method

The name of a method doesn’t explain what the method does.

## Example

### Before

```
func Calculate(a, b int) int {
    return a + b
}
```

### After

```
func AddNumbers(a, b int) int {
    return a + b
}
```

In this refactored code, the method name AddNumbers clearly conveys its purpose, making the comment unnecessary. The code is now more self-explanatory and easier to understand.
