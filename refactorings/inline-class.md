# Inline Class

Involves removing a class that doesn't have much functionality and moving its fields and methods directly into the calling class.

## Example

### Before

```
type Person struct {
    name    string
    address string
}

func NewPerson(name, address string) *Person {
    return &Person{
        name:    name,
        address: address,
    }
}

func (p *Person) GetDetails() string {
    return fmt.Sprintf("Name: %s, Address: %s", p.name, p.address)
}

func main() {
    person := NewPerson("John Doe", "123 Main St")
    details := person.GetDetails()
    fmt.Println(details)
}
```

### After

```
func main() {
    name := "John Doe"
    address := "123 Main St"
    details := fmt.Sprintf("Name: %s, Address: %s", name, address)
    fmt.Println(details)
}
```

we've eliminated the Person class and directly used the name and address variables in the main function. This simplifies the code by removing unnecessary abstraction introduced by the Person class.