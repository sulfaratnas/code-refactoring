# Move Method

We can implement this technique if there is a method is used more in another class than in its own class or that method could be more appropriately belong to another class.

## Example

### Before

```
type Order struct {
    ProductName string
    Price       float64
    Quantity    int
}

type Customer struct {
    Name   string
    Orders []Order
}

func (c *Customer) CalculateTotalPrice() float64 {
    total := 0.0
    for _, order := range c.Orders {
        total += order.Price * float64(order.Quantity)
    }
    return total
}

func main() {
    orders := []Order{
        {ProductName: "Item 1", Price: 10.0, Quantity: 2},
        {ProductName: "Item 2", Price: 15.0, Quantity: 1},
    }
    customer := Customer{Name: "John Doe", Orders: orders}

    totalPrice := customer.CalculateTotalPrice()
    fmt.Printf("%s's Total Price: $%.2f\n", customer.Name, totalPrice)
}
```

### After

```
type Order struct {
    ProductName string
    Price       float64
    Quantity    int
}

type Customer struct {
    Name   string
    Orders []Order
}

func (o *Order) CalculateTotalPrice() float64 {
    return o.Price * float64(o.Quantity)
}

func (c *Customer) CalculateTotalPrice() float64 {
    total := 0.0
    for _, order := range c.Orders {
        total += order.CalculateTotalPrice()
    }
    return total
}

func main() {
    orders := []Order{
        {ProductName: "Item 1", Price: 10.0, Quantity: 2},
        {ProductName: "Item 2", Price: 15.0, Quantity: 1},
    }
    customer := Customer{Name: "John Doe", Orders: orders}

    totalPrice := customer.CalculateTotalPrice()
    fmt.Printf("%s's Total Price: $%.2f\n", customer.Name, totalPrice)
}
```

In this refactoring example:

We've moved the CalculateTotalPrice method from the Customer struct to the Order struct, as it makes more sense for an order to calculate its own total price.

We've updated the CalculateTotalPrice method in the Customer struct to iterate through the orders and call the CalculateTotalPrice method of each order.

This refactoring eliminates the feature envy by ensuring that the calculation of the total price is performed within the relevant object, improving code maintainability and readability.
