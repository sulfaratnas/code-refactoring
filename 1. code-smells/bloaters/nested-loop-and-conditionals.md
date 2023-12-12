# Nested Loops And Conditionals

Also known as the ["Arrow Anti-Pattern"](http://wiki.c2.com/?ArrowAntiPattern).

## What It Looks Like

```
func main() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Outer loop iteration %d\n", i)

		for j := 1; j <= 2; j++ {
			fmt.Printf("Inner loop iteration %d\n", j)

			if i == 2 && j == 1 {
				fmt.Println("Condition met: i equals 2 and j equals 1")
			}
		}
	}
}
```

## Why It Hurts

It's hard to see at a glance what the code is doing, and where it might be unsafe to insert new code or reorder statements.

## How To Fix It

- We can use [Replace Inline Code With Function Call](.././../2.%20refactorings/replace-inline-code-with-function-call.md)
and [Extract Method](.././../2.%20refactorings/extract-method.md) to simplify this complicated nested loop, so the method would be easier to understand.

## Refactor

```
func main() {
	for i := 1; i <= 3; i++ {
		outerLoop(i)
	}
}

func outerLoop(i int) {
	fmt.Printf("Outer loop iteration %d\n", i)

	for j := 1; j <= 2; j++ {
		innerLoop(i, j)
	}
}

func innerLoop(i, j int) {
	fmt.Printf("Inner loop iteration %d\n", j)

	if i == 2 && j == 1 {
		fmt.Println("Condition met: i equals 2 and j equals 1")
	}
}
```

## Payoff

- Improves code readability.
- In many cases, splitting large classes into parts avoids duplication of code and functionality.
