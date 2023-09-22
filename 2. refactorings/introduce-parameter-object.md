# Introduce Parameter Object

This refactoring changes the interface of a method.

It applies when the method has a large number of parameters
that form one or more logical groupings.

It would be better to pass fewer objects that group the data
together. This would allow people reading the code to more
easily verify that the arguments are passed correctly.

## Example

### Before

```
func distance(x1, y1, x2, y2) (result float64) {
    # ...
    return result
}

func main() {
    distance(-1, 2, 3, 5)
}
```

### After
```
type Point struct {
	X float64
	Y float64
}

func distance(a, b) (result float64) {
     # ...
     return result
}

func main() {
    distance(Point{X:-1, Y:2}, Point{X:3, Y:5})
}
```

The distance function now takes two Point objects as parameters, encapsulating the related data. 
This refactoring technique makes the function signature cleaner and more readable while grouping related parameters into a single object.
