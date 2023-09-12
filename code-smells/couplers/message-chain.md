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

Sometimes it’s better to think of why the end object is being used. Perhaps it would make sense to use [Extract Method](.../refactorings/extract-method.md) for this functionality and move it to the beginning of the chain, by using [Move Method](.../refactorings/move-method.md).