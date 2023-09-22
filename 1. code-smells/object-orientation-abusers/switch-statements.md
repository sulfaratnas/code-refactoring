# Switch Statements

You have a complex switch operator or sequence of if statements.

## What It Looks Like


```
// BirdType represents the type of bird.
type BirdType int

const (
	EUROPEAN BirdType = iota
	AFRICAN
	NORWEGIAN_BLUE
)

// Bird represents a bird.
type Bird struct {
	BirdType        BirdType
	BaseSpeed       float64
	LoadFactor      float64
	NumberOfCoconuts int
	IsNailed        bool
	Voltage         float64
}

func NewBird(birdType BirdType) *Bird {
	return &Bird{BirdType: birdType}
}

func (b *Bird) GetSpeed() float64 {
	switch b.BirdType {
	case EUROPEAN:
		return b.getBaseSpeed()
	case AFRICAN:
		return b.getBaseSpeed() - b.getLoadFactor()*float64(b.NumberOfCoconuts)
	case NORWEGIAN_BLUE:
		if b.IsNailed {
			return 0
		}
		return b.getBaseSpeed(b.Voltage)
	default:
		panic(errors.New("Should be unreachable"))
	}
}

func (b *Bird) getBaseSpeed(voltage ...float64) float64 {
	if len(voltage) > 0 {
		return voltage[0]
	}
	return b.BaseSpeed
}

func main() {
	bird := NewBird(NORWEGIAN_BLUE)
	bird.BaseSpeed = 10.0
	bird.Voltage = 220.0
	bird.IsNailed = false

	speed := bird.GetSpeed()
	fmt.Printf("Bird Speed: %.2f\n", speed)
}
```

## Why It Hurts

Switch statements become overly complex, lengthy, or difficult to maintain.

## How To Fix It

[Replace Conditional With Polymorphism](.././../refactorings/replace-conditional-with-polymorphism.md).

