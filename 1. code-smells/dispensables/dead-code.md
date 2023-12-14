# Dead Code

A variable, parameter, field, method or class is no longer used (usually because it’s obsolete).

## What It Looks Like

```
// Function with dead code
func calculateSum(a, b int) int {
	// This block of code is not reachable
	if false {
		// Dead code - unreachable
		fmt.Println("This code is never executed")
	}

	// The following line is also dead code
	return a + b
}
```

## Why It Hurts

Dead code adds unnecessary complexity to the codebase, making it harder to understand and maintain.


## How To Fix It

- [Remove Dead Code](.././../2.%20refactorings/remove-dead-code.md).
- [Remove Parameter](.././../2.%20refactorings/remove-parameter.md) to remove unneeded parameters

## Refactor 

```
func calculateSum(a, b int) int {
	return a + b
}
```

## Payoff

- Reduced code size.
- Simpler support.

