# Data Clump

## What It Looks Like

Multiple parameters or variables consistently appear together in function or method signatures, indicating that they may belong together as a single data structure.

```
func CalculateDistance(x1, y1, x2, y2 float64) float64 {
	deltaX := x2 - x1
	deltaY := y2 - y1
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
```

In this code, `x1` and `y1` form one data clump—they
represent a point in the Cartesian plane. `x2` and `y2` form
another clump.

## Why It Hurts

When a method takes many positional arguments, calls to the
method become difficult to read. It's too easy to pass
arguments in the wrong order, or reference the wrong parameter within the method implementation.

## How To Fix It

Instead of extracting individual fields from an object to pass them to a function, 
use [Preserve Whole Object](.././../refactorings/preserve-whole-object.md). If there is no such object, 
use [Introduce Parameter Object](.././../refactorings/introduce-parameter-object.md).

