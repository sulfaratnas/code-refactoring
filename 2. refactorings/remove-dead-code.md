# Remove Dead Code

## Example

### Before

```
// Function with dead code
func calculateSum(a, b int) int {
	// This block of code is not reachable
	if false {
		// Dead code - unreachable
		fmt.Println("This code is never executed")
	}

	// The following line is also dead code
	return a + b
}
```

### After

```
func calculateSum(a, b int) int {
	return a + b
}
```