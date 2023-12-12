# Primitive Obsession

Primitive Obsession is a code smell in which primitive data types are used excessively to represent your data models. Instead of creating a new class for fields it is easy to use strings, integers or collections to simulate types.

## What It Looks Like

```
type Person struct {
    FirstName string
    LastName string
    Address string
    ...
}
```

## Why It Hurts

- This can lead to a lack of type safety, readability, and maintainability in your code.
- Your code becomes coupled to the underlying structure of the data, making it harder to test and reason about.


## How To Fix It

[Replace Data Value With Object](.././../2.%20refactorings/replace-data-value-with-object.md).

## Refactor

```
type Person struct {
    Name Name
    Address string
    ...
}

type Name struct {
    FirstName string
    LastName string
}
```
Fields that logically belong together can be combined by replacing data value with object. This improves relatedness.

## Payoff

- Better code organization and flexibility
- Improves code readability, no more guessing about the reason for all these strange constants