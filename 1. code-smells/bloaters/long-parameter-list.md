# Long Parameter list

The "Long Parameter List" code smell occurs when a function or method has too many parameters, which can make the code harder to read and maintain.

## What It Looks Like

```
func CalculateSum(a, b, c, d, e int) int {
    return a + b + c + d + e
}
```

## Why It Hurts

It can become a code smell when the number of parameters grows, especially if they are of different types or have complex meanings.

## How To Fix It

The best fix is to [Introduce Parameter Object](.../refactorings/introduce-parameter-object.md) that group data together, 
[Preserve Whole Object](.../refactorings/preserve-whole-object.md) to simplify function signature, 
[Replace Parameter With Method Call](.../refactorings/replace-parameter-with-method-call.md).

## Payoff

- More readable, shorter code.
- Refactoring may reveal previously unnoticed duplicate code.