# Message Chains

## What It Looks Like

It occurs when you have a series of method calls or property accesses chained together to retrieve information from nested objects.

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
In this example, GetCustomerCityWithMessageChain in the Order struct directly access to Customer struct.

## Why It Hurts

- Long chains of method calls or property accesses can make code harder to read and understand.
- Message chains create tight coupling between the caller and the objects in the chain.
- As the codebase evolves, maintaining and extending code with message chains becomes increasingly difficult.
- Writing unit tests for code with message chains can be challenging.
- Lack of encapsulation

## How To Fix It

- [Hide Delegate](.././../2.%20refactorings/hide-delegate.md)
- Sometimes it’s better to think of why the end object is being used. Perhaps it would make sense to use [Extract Method](.././../2.%20refactorings/extract-method.md) for this functionality and move it to the beginning of the chain, by using [Move Method](.././../2.%20refactorings/move-method.md).

## Refactor

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

## Payoff

- Reduces dependencies between classes of a chain.
- Reduces the amount of bloated code.