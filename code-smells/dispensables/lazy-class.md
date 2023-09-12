# Lazy Class

A class (or struct in Go) that doesn't do enough meaningful work to justify its existence. It often contains very few methods or fields and doesn't contribute significantly to the functionality of the application. Eliminating unnecessary classes or structs can lead to cleaner and more maintainable code.

## What It Looks Like

```
// LazyClass is an example of a class with minimal functionality.
type LazyClass struct {
	name string
}

func NewLazyClass(name string) *LazyClass {
	return &LazyClass{name: name}
}

func main() {
	// Create an instance of LazyClass
	lazy := NewLazyClass("John")

	// Print the name
	fmt.Println("Name:", lazy.name)
}
```

## Why It Hurts

Lazy Classes that exist without serving a purpose can hurt code quality, readability, and maintenance efforts, and they should be refactored or removed to keep the codebase clean and maintainable.

## How To Fix It

[Inline Class](.../refactorings/inline-class.md).

