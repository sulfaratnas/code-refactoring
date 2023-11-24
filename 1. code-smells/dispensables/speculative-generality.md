# Speculative Generality

There’s an unused class, method, field or parameter. This smell it’s really similar to Dead code. The difference is only on purpose. If the Dead code appears because of unmaintained code, the Speculative generality appears because the engineer has too much initiative. The engineer tries to cover something in the future but is not really used by now. Then after months, we realize that the code isn’t used by anyone.

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
[Inline Class](.././../2.%20refactorings/inline-class.md).
[Remove Parameter](.././../2.%20refactorings/remove-parameter.md).

## Refactor

```
func greetCustomer(customerName string) {
	fmt.Println("Hello, ", customerName)
}

func main() {
	customerName := "John"
	greetCustomer(customerName)
}
```

