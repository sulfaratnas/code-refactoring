# Replace Data Value With Object

## Example

### Before

In this "before" example, we represent a customer's name and age using primitive data types, which can lead to a lack of type safety and readability.


```
func main() {
    // Original code with primitive obsession
    customerName := "John Doe"
    customerAge := 30

    fmt.Printf("Customer Name: %s\n", customerName)
    fmt.Printf("Customer Age: %d\n", customerAge)
}
```

### After

```
type Customer struct {
    Name string
    Age  int
}

func NewCustomer(name string, age int) *Customer {
    return &Customer{
        Name: name,
        Age:  age,
    }
}

func main() {
    // Refactored code using custom types
    customer := NewCustomer("John Doe", 30)

    fmt.Printf("Customer Name: %s\n", customer.Name)
    fmt.Printf("Customer Age: %d\n", customer.Age)
}
```

This refactoring improves code readability, type safety, and maintainability by encapsulating the customer's name and age within a custom type, rather than using primitive data types. It also makes it easier to add additional fields or behavior related to customers in the future.