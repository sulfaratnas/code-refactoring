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

[Replace Conditional With Polymorphism](.././../2.%20refactorings/replace-conditional-with-polymorphism.md).

## Refactor

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

type SpeedCalculator interface {
	CalculateSpeed(bird *Bird) float64
}

type EuropeanSpeedCalculator struct{}

func (c *EuropeanSpeedCalculator) CalculateSpeed(bird *Bird) float64 {
	return bird.getBaseSpeed()
}

type AfricanSpeedCalculator struct{}

func (c *AfricanSpeedCalculator) CalculateSpeed(bird *Bird) float64 {
	return bird.getBaseSpeed() - bird.getLoadFactor()*float64(bird.NumberOfCoconuts)
}

type NorwegianBlueSpeedCalculator struct{}

func (c *NorwegianBlueSpeedCalculator) CalculateSpeed(bird *Bird) float64 {
	if bird.IsNailed {
		return 0
	}
	return bird.getBaseSpeed(bird.Voltage)
}

func (b *Bird) GetSpeed() float64 {
	var calculator SpeedCalculator

	switch b.BirdType {
	case EUROPEAN:
		calculator = &EuropeanSpeedCalculator{}
	case AFRICAN:
		calculator = &AfricanSpeedCalculator{}
	case NORWEGIAN_BLUE:
		calculator = &NorwegianBlueSpeedCalculator{}
	default:
		panic(errors.New("Should be unreachable"))
	}

	return calculator.CalculateSpeed(b)
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

