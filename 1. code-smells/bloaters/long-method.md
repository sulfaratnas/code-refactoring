# Long Method

A method has many lines, possibly involving conditionals,
loops, variable assignments, and early exits.

## What It Looks Like

```
func ComplexOperation(value int) {
    fmt.Println("Starting complex operation...")

    if value > 10 {
        fmt.Println("Value is greater than 10.")
    } else {
        fmt.Println("Value is not greater than 10.")
    }

    // Long and complex calculation step 1
    // ...
    fmt.Println("Step 1 completed.")

    if value < 5 {
        fmt.Println("Value is less than 5.")
    } else {
        fmt.Println("Value is not less than 5.")
    }

    // Long and complex calculation step 2
    // ...
    fmt.Println("Step 2 completed.")

    // More conditional statements...

    fmt.Println("Complex operation finished.")
}
```

## Why It Hurts

Long methods tend to be hard to understand when they entangle multiple responsibilities or mutate state. When
reading the code, you have to understand not only what each line does, but how it interacts with the rest of the
method's code. The difficulty of understanding a method is thus something like O(n^2) in the method's length.


## How To Fix It

The best fix is to [Replace Inline Code With Function Call](.../refactorings/replace-inline-code-with-function-call.md)—that
is, call other existing methods to implement the logic. 
If no suitable helper methods exist, [Extract Method](.../refactorings/extract-method.md) can break up the
long method into more digestible chunks.
