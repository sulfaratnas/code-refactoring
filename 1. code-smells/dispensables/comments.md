# Comments

## What It Looks Like

Comments are usually created with the best of intentions, when the author realizes that his or her code isn’t intuitive or obvious. In such cases, comments are like a deodorant masking the smell of fishy code that could be improved.

```
// Calculate adds two numbers and returns the result.
func Calculate(a, b int) int {
    return a + b
}

func main() {
    result := Calculate(5, 10)

    fmt.Println("The result is:", result)
}
```
In this example, the Calculate function has a name that's not very descriptive, so a comment is added to explain its purpose. This is a code smell because the method name should ideally be clear enough to convey its purpose without needing a comment.


## Why It Hurts

Redundant comments don't improve code readability.

## How To Fix It

If a comment is intended to explain a complex expression, the expression should be split into understandable subexpressions using [Extract Variable](.././../2.%20refactorings/extract-variable.md).

If a comment explains a section of code, this section can be turned into a separate method via [Extract Method](.././../2.%20refactorings/extract-method.md). The name of the new method can be taken from the comment text itself, most likely.

If a method has already been extracted, but comments are still necessary to explain what the method does, give the method a self-explanatory name. Use [Rename Method](.././../2.%20refactorings/rename-method.md) for this.

## Refactor

```
func AddNumbers(a, b int) int {
    return a + b
}

func main() {
    sum := AddNumbers(5, 10)
    fmt.Println("Sum:", sum)
}
```

## Payoff

Code becomes more intuitive and obvious.
