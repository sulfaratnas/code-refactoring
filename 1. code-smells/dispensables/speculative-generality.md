# Speculative Generality

There’s an unused class, method, field or parameter.

## What It Looks Like

```
// the second variable has been never used
func greetCustomer(customerName string, _ int) {
	fmt.Println("hello, ", customerName)
}
```

## Why It Hurts

Sometimes code is created “just in case” to support anticipated future features that never get implemented. As a result, code becomes hard to understand and support.

## How To Fix It

Unused fields can be simply deleted.
[Inline Class](.././../refactorings/inline-class.md).
[Remove Parameter](.././../refactorings/remove-parameter.md).
