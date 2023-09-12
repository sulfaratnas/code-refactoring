# Remove Middle Man

## Example

### Before

```
// Service represents a service that clients want to use.
type Service struct{}

// DoSomething is a method of the Service.
func (s *Service) DoSomething() string {
	return "Service is doing something."
}

// MiddleMan represents a middleman class that delegates work to the Service.
type MiddleMan struct {
	service *Service
}

// NewMiddleMan creates a new instance of MiddleMan.
func NewMiddleMan() *MiddleMan {
	return &MiddleMan{
		service: &Service{},
	}
}

// DoSomething is a method of the MiddleMan that delegates work to the Service.
func (m *MiddleMan) DoSomething() string {
	return m.service.DoSomething()
}

func main() {
	middleMan := NewMiddleMan()
	result := middleMan.DoSomething()
	fmt.Println(result)
}
```

### After

```
// Service represents a service that clients want to use.
type Service struct{}

// DoSomething is a method of the Service.
func (s *Service) DoSomething() string {
	return "Service is doing something."
}

func main() {
	service := &Service{}
	result := service.DoSomething()
	fmt.Println(result)
}
```