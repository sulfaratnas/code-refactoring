# Lazy Class

A class (or struct in Go) that doesn't do enough meaningful work to justify its existence. It often contains very few methods or fields and doesn't contribute significantly to the functionality of the application. It might caused by refactoring and move some method from a large class to new class to implement Single Responsibility Principle, but this new class has minimum functionality, extremely only has one method. Eliminating unnecessary classes or structs can lead to cleaner and more maintainable code.

## What It Looks Like

```
// PriceValidator represents a price validator.
type PriceValidator struct{}

// Validate checks if the value is greater than or equal to 0.
func (pv *PriceValidator) Validate(value int) bool {
	return value >= 0
}

// Price represents a price with validation.
type Price struct {
	value int
}

// NewPrice creates a new Price instance with validation.
func NewPrice(value int) (*Price, error) {
	validator := PriceValidator{}

	if !validator.Validate(value) {
		return nil, errors.New("price not valid")
	}

	return &Price{value: value}, nil
}
```
In this example, PriceValidator is only used for validate price.

## Why It Hurts

Lazy Classes that exist without serving a purpose can hurt code quality, readability, and maintenance efforts, and they should be refactored or removed to keep the codebase clean and maintainable.

## How To Fix It

[Inline Class](.././../2.%20refactorings/inline-class.md).

## Refactor

```
// NewPrice creates a new Price instance with validation.
func NewPrice(value int) (*Price, error) {
	if value < 0 {
		return nil, errors.New("price not valid")
	}

	return &Price{value: value}, nil
}

// Price represents a price with validation.
type Price struct {
	value int
}

// GetValue returns the value of the price.
func (p *Price) GetValue() int {
	return p.value
}
```

## Payoff

- Reduced code size.
- Easier maintenance.

