# Hide Delegate

When the client gets object B from a field or method of object А, then the client calls a method of object B, create a new method in class A that delegates the call to object B. Now the client doesn’t know about, or depend on, class B.

## Example

### Before

```
package main

type Address struct {
	City string
}

type Customer struct {
	Address *Address
}

func (c *Customer) GetCity() string {
	if c.Address != nil {
		return c.Address.City
	}
	return "Unknown"
}

type Order struct {
	Customer *Customer
}

func (o *Order) GetCustomerCityWithMessageChain() string {
	if o.Customer != nil && o.Customer.Address != nil {
		return o.Customer.Address.City
	}
	return "Unknown"
}

func main() {
	address := &Address{City: "ExampleCity"}
	customer := &Customer{Address: address}
	order := &Order{Customer: customer}

	fmt.Println(order.GetCustomerCityWithMessageChain())
}
```

### After

```
package main

type Address struct {
	City string
}

type Customer struct {
	Address *Address
}

func (c *Customer) GetCity() string {
	if c.Address != nil {
		return c.Address.City
	}
	return "Unknown"
}

type Order struct {
	Customer *Customer
}

func (o *Order) GetCustomerCity() string {
	if o.Customer != nil {
		return o.Customer.GetCity()
	}
	return "Unknown"
}

func main() {
	address := &Address{City: "ExampleCity"}
	customer := &Customer{Address: address}
	order := &Order{Customer: customer}

	fmt.Println(order.GetCustomerCity())
}
```

## References

- [Extract Method on Refactoring.com](https://refactoring.com/catalog/extractFunction.html).
- [Extract Method on Refactoring.guru](https://refactoring.guru/extract-method)