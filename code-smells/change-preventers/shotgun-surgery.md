# Shotgun Surgery

The "Shotgun Surgery" code smell occurs when a single change to a feature or functionality requires modifications in multiple places across the codebase. 

## What It Looks Like

```
// Circle represents a circle with a radius.
type Circle struct {
	Radius float64
}

// Rectangle represents a rectangle with width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// CalculateArea calculates the area of a circle.
func (c *Circle) CalculateArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

// CalculatePerimeter calculates the perimeter of a circle.
func (c *Circle) CalculatePerimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// CalculateArea calculates the area of a rectangle.
func (r *Rectangle) CalculateArea() float64 {
	return r.Width * r.Height
}

// CalculatePerimeter calculates the perimeter of a rectangle.
func (r *Rectangle) CalculatePerimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func main() {
	circle := &Circle{Radius: 5.0}
	rectangle := &Rectangle{Width: 4.0, Height: 6.0}

	// Calculate and print areas and perimeters
	circleArea := circle.CalculateArea()
	circlePerimeter := circle.CalculatePerimeter()
	rectangleArea := rectangle.CalculateArea()
	rectanglePerimeter := rectangle.CalculatePerimeter()

	fmt.Printf("Circle Area: %.2f\n", circleArea)
	fmt.Printf("Circle Perimeter: %.2f\n", circlePerimeter)
	fmt.Printf("Rectangle Area: %.2f\n", rectangleArea)
	fmt.Printf("Rectangle Perimeter: %.2f\n", rectanglePerimeter)
}
```
In this example, we have a Circle and Rectangle type, each with methods for calculating their area and perimeter. Whenever a new shape is introduced or if we want to change the calculation logic for any shape, we must make changes in multiple places, leading to the "Shotgun Surgery" code smell.

## Why It Hurts

- The more places you have to make changes, the greater the chance of introducing errors.
- Shotgun surgery often leads to changes being scattered across multiple files or classes.
- Frequent shotgun surgery can significantly increase the time it takes to implement new features or fix bugs. 
- Decreased Code Readability.

## How To Fix It

Use [Move Method](.././../refactorings/move-method.md) to move existing class behaviors into a single class. If there’s no class appropriate for this, create a new one.

If moving code to the same class leaves the original classes almost empty, try to get rid of these now-redundant classes via [Inline Class](.././../refactorings/inline-class.md).
