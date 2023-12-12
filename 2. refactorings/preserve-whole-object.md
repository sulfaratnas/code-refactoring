# Preserve Whole Object

The "Preserve Whole Object" refactoring technique is used when a function need several parameters from an object. You can try passing the whole object.

## Example

### Before

```
type Point struct{
	X float64
	Y float64
}

point := Point{
	X: 1,
	Y: 2,
}

printCoordinate(point.X, point.Y);
```

### After

```
type Point struct{
	X float64
	Y float64
}

point := Point{
	X: 1,
	Y: 2,
}

printCoordinate(point)
```

In this refactored code, we've introduced a Point struct to represent a point with X and Y coordinates. The printCoordinate function now accepts a Point object as a parameter. This refactoring preserves the whole object and makes the code more readable and maintainable, especially if you have more functions operating on points.
