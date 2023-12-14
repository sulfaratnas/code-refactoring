# Switch Statements

You have a complex switch operator or sequence of if statements.

## What It Looks Like


```
type ShapeType int

const (
	CircleType ShapeType = iota
	RectangleType
)

type Shape struct {
	Type ShapeType
}

func CalculateArea(shape Shape) float64 {
	switch shape.Type {
	case CircleType:
		// Complex logic for Circle
		return calculateCircleArea()
	case RectangleType:
		// Complex logic for Rectangle
		return calculateRectangleArea()
	default:
		return 0
	}
}

func calculateCircleArea() float64 {
	// Complex logic for calculating Circle area
	return 3.14 * 5 * 5 // Assuming a fixed radius for simplicity
}

func calculateRectangleArea() float64 {
	// Complex logic for calculating Rectangle area
	return 4 * 6 // Assuming fixed dimensions for simplicity
}

func main() {
	circle := Shape{Type: CircleType}
	rectangle := Shape{Type: RectangleType}

	circleArea := CalculateArea(circle)
	rectangleArea := CalculateArea(rectangle)

	fmt.Printf("Circle area: %f\n", circleArea)
	fmt.Printf("Rectangle area: %f\n", rectangleArea)
}
```
In this example, when we want to add a new shape, we need to modify the switch case. It violates open/closed principle from SOLID.

## Why It Hurts

Switch statements become overly complex, lengthy, or difficult to maintain.

## How To Fix It

- To isolate switch and put it in the right class, you may need [Extract Method](.././../2.%20refactorings/extract-method.md).
- After specifying the inheritance structure, use [Replace Conditional With Polymorphism](.././../2.%20refactorings/replace-conditional-with-polymorphism.md).

## When to ignore

- When a switch operator performs simple actions, there’s no reason to make code changes.
- Often switch operators are used by factory design patterns (Factory Method or Abstract Factory) to select a created class.

## Refactor

```
// Shape interface representing common behavior for all shapes
type Shape interface {
	CalculateArea() float64
}

// Circle type implementing the Shape interface
type Circle struct {
	Radius float64
}

// CalculateArea method for Circle
func (c Circle) CalculateArea() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle type implementing the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// CalculateArea method for Rectangle
func (r Rectangle) CalculateArea() float64 {
	return r.Width * r.Height
}

func main() {
	// Creating instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Using polymorphism to calculate areas
	circleArea := calculateArea(circle)
	rectangleArea := calculateArea(rectangle)

	// Printing the results
	fmt.Printf("Circle area: %f\n", circleArea)
	fmt.Printf("Rectangle area: %f\n", rectangleArea)
}

// calculateArea function accepting any type that implements the Shape interface
func calculateArea(shape Shape) float64 {
	return shape.CalculateArea()
}
```
After implementing polymorphism, when we want to add a new shape, we do not need to modify existing code/function, we just need to create a new class that implements `Shape` interface.

