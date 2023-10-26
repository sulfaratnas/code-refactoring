# Duplicated Code

Two code fragments look almost identical.


## What It Looks Like

```
func main() {
    x := 5
    y := 10

    // Calculation 1
    sum1 := x + y
    fmt.Printf("Sum 1: %d\n", sum1)

    // Calculation 2 (Identical to Calculation 1)
    sum2 := x + y
    fmt.Printf("Sum 2: %d\n", sum2)
}
```

## Why It Hurts

It can lead to maintenance challenges, increased risk of errors, and reduced code maintainability.

## How To Fix It

[Extract Method](.././../2.%20refactorings/extract-method.md).

## Refactor

```
func calculateSum(x, y int) int {
    return x + y
}

func main() {
    x := 5
    y := 10

    // Calculation 1
    sum1 := calculateSum(x, y)
    fmt.Printf("Sum 1: %d\n", sum1)

    // Calculation 2 (Identical to Calculation 1)
    sum2 := calculateSum(x, y)
    fmt.Printf("Sum 2: %d\n", sum2)
}
```

## Payoff

- Merging duplicate code simplifies the structure of your code and makes it shorter.
- Simplification + shortness = code that’s easier to simplify and cheaper to support.

