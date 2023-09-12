# Primitive Obsession

Use of primitives instead of small objects for simple tasks (such as currency, ranges, special strings for phone numbers, etc.)

Use of constants for coding information (such as a constant USER_ADMIN_ROLE = 1 for referring to users with administrator rights.)

Use of string constants as field names for use in data arrays.

## What It Looks Like

```
func main() {
    // Original code with primitive obsession
    customerName := "John Doe"
    customerAge := 30

    fmt.Printf("Customer Name: %s\n", customerName)
    fmt.Printf("Customer Age: %d\n", customerAge)
}
```

## Why It Hurts

This can lead to a lack of type safety, readability, and maintainability in your code.


## How To Fix It

[Replace Data Value With Object](../refactorings/replace-data-value-with-object.md).