# Message Chains

## What It Looks Like

It occurs when you have a series of method calls or property accesses chained together to retrieve information from nested objects.

```
func main() {
	address := &Address{
		City:    "New York",
		State:   "NY",
		Country: "USA",
	}

	person := &Person{
		Name:    "John Doe",
		Address: address,
	}

	// Message chain to retrieve the person's country
	country := person.Address.Country

	fmt.Printf("%s lives in %s, %s, %s\n", person.Name, person.Address.City, person.Address.State, country)
}
```

## Why It Hurts

- Long chains of method calls or property accesses can make code harder to read and understand.
- Message chains create tight coupling between the caller and the objects in the chain.
- As the codebase evolves, maintaining and extending code with message chains becomes increasingly difficult.
- Writing unit tests for code with message chains can be challenging.

## How To Fix It

Sometimes it’s better to think of why the end object is being used. Perhaps it would make sense to use [Extract Method](.././../2.%20refactorings/extract-method.md) for this functionality and move it to the beginning of the chain, by using [Move Method](.././../2.%20refactorings/move-method.md).

## Refactor

```
type Address struct {
	City    string
	State   string
	Country string
}

type Person struct {
	Name    string
	Address *Address
}

func (p *Person) GetFullAddress() string {
	return fmt.Sprintf("%s, %s, %s", p.Address.City, p.Address.State, p.Address.Country)
}

func main() {
	address := &Address{
		City:    "New York",
		State:   "NY",
		Country: "USA",
	}

	person := &Person{
		Name:    "John Doe",
		Address: address,
	}

	fullAddress := person.GetFullAddress()

	fmt.Printf("%s lives in %s\n", person.Name, fullAddress)
}
```

## Payoff

- Reduces dependencies between classes of a chain.
- Reduces the amount of bloated code.