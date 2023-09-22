# Extract Method

## Example

### Before

```
func main() {
    // Original code with a long method
    name := "John Doe"
    age := 30

    // Greet the person
    fmt.Println("Hello, my name is", name)

    // Introduce the person
    fmt.Println("I am", age, "years old.")
    fmt.Println("Nice to meet you!")
}
```

### After

```
func main() {
    // Original code with a long method
    name := "John Doe"
    age := 30

    // Refactored code with extracted methods
    greetPerson(name)
    introducePerson(age)
}

func greetPerson(name string) {
    fmt.Println("Hello, my name is", name)
}

func introducePerson(age int) {
    fmt.Println("I am", age, "years old.")
    fmt.Println("Nice to meet you!")
}
```

## References

- [Extract Method on Refactoring.com](https://refactoring.com/catalog/extractFunction.html).
- [Extract Method on Refactoring.guru](https://refactoring.guru/extract-method)