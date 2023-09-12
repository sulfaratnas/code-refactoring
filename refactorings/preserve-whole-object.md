# Preserve Whole Object

The "Preserve Whole Object" refactoring technique is used when a function or method is being called with multiple individual data parameters, but those parameters frequently belong to the same object or entity. 

Instead of passing individual parameters, you pass the entire object to simplify the function signature and make the code more maintainable.

## Example

### Before

```
func PrintX(x float64) {
	fmt.Printf("X: %.2f\n", x)
}

func PrintY(y float64) {
	fmt.Printf("Y: %.2f\n", y)
}

func main() {
	x := 10.0
	y := 5.0

	PrintX(x)
	PrintY(y)
}
```

### After

```
type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

func PrintPoint(point *Point) {
	fmt.Printf("X: %.2f\n", point.X)
	fmt.Printf("Y: %.2f\n", point.Y)
}

func main() {
	point := NewPoint(10.0, 5.0)
	PrintPoint(point)
}
```

In this refactored code, we've introduced a Point struct to represent a point with X and Y coordinates. The PrintPoint function now accepts a Point object as a parameter. This refactoring preserves the whole object and makes the code more readable and maintainable, especially if you have more functions operating on points.
