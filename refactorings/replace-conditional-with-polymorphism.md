# Replace Conditional With Polymorphism

The name of a method doesn’t explain what the method does.

## Example

### Before

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

### After

```
/ Bird defines the Bird interface with a GetSpeed method.
type Bird interface {
	GetSpeed() float64
}

// European represents a European bird.
type European struct {
	BaseSpeed float64
}

func (e *European) GetSpeed() float64 {
	return e.BaseSpeed
}

// African represents an African bird.
type African struct {
	BaseSpeed       float64
	LoadFactor      float64
	NumberOfCoconuts int
}

func (a *African) GetSpeed() float64 {
	return a.BaseSpeed - a.LoadFactor*float64(a.NumberOfCoconuts)
}

// NorwegianBlue represents a Norwegian Blue bird.
type NorwegianBlue struct {
	BaseSpeed float64
	IsNailed  bool
	Voltage   float64
}

func (n *NorwegianBlue) GetSpeed() float64 {
	if n.IsNailed {
		return 0
	}
	return n.BaseSpeed * n.Voltage
}

func main() {
	european := &European{BaseSpeed: 10.0}
	african := &African{BaseSpeed: 10.0, LoadFactor: 1.5, NumberOfCoconuts: 2}
	norwegianBlue := &NorwegianBlue{BaseSpeed: 10.0, IsNailed: false, Voltage: 220.0}

	// Somewhere in client code
	birds := []Bird{european, african, norwegianBlue}
	for _, bird := range birds {
		speed := bird.GetSpeed()
		fmt.Printf("Bird Speed: %.2f\n", speed)
	}
}
```

In this refactored code, the method name AddNumbers clearly conveys its purpose, making the comment unnecessary. The code is now more self-explanatory and easier to understand.
