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

Code in loops and conditionals tends to conflate data transformation with side effects. This is because there's no
way for a loop or an `if` to affect the rest of the system unless it has a side effect. As a result, it's hard
to see at a glance what the code is doing, and where it might be unsafe to insert new code or reorder statements.

## How To Fix It

Use [Replace Loop With Pipeline](.../refactorings/replace-loop-with-pipeline.md) to
separate data transformation from side effects.
Additionally,[Replace Inline Code With Function Call](.../refactorings/replace-inline-code-with-function-call.md)
and [Extract Method](.../refactorings/extract-method.md) can help abstract away complex boolean logic in conditionals.

## Payoff

- Improves code readability.
- In many cases, splitting large classes into parts avoids duplication of code and functionality.
