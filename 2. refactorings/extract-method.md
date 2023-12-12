# Extract Method

We can use this technique when you have code fragment that can be grouped together. We just need to move this code fragment to a new method and change the old code by calling this method.

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
    name := "John Doe"
    age := 30

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