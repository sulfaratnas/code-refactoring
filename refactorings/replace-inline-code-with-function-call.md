# Replace Inline Code With Function Call

Replace a chunk of code with a call to a method that does
the same thing. Often, you can find a standard library
function that does what you need.

This is one of the most valuable refactorings, because it
makes code more expressive and readable *and* reduces the
total amount of code in the codebase!

## Example

### Before

```
func main() {
	// Create a slice of strings
	strSlice := []string{"Hello", "World", "in", "Go"}

	// Manually join the strings
	var result string
	for _, str := range strSlice {
		result += str + " " // Add a space separator
	}
	fmt.Println(result)
}
```

### After

```
func main() {
	// Create a slice of strings
	strSlice := []string{"Hello", "World", "in", "Go"}

	// Use strings.Join to join the strings
	result := strings.Join(strSlice, " ")

	fmt.Println(result)
}
```

## References

- [Replace Inline Code With Function Call on Refactoring.com](https://refactoring.com/catalog/replaceInlineCodeWithFunctionCall.html)
