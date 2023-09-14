# Duplicated Code

Two code fragments look almost identical.


## What It Looks Like

```
package main

import "fmt"

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

[Extract Method](.././../refactorings/extract-method.md).

