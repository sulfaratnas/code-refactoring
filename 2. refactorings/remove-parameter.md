# Remove Parameter

A parameter isn’t used in the body of a method. Need to remove the unused parameter.

## Example

### Before

```
// customerName isn't used
func greetCustomer(customerName string) {
    fmt.Println("hello")
}
```

### After

```
func greetCustomer() {
    fmt.Println("hello")
}
```

