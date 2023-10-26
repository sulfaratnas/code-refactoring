# Middle Man

A class performs only one action, delegating work to another class. It occurs when a class or object serves as an unnecessary intermediary between clients and the services it delegates to. It can add unnecessary complexity to the code and provide little or no value. 

## What It Looks Like

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

## Why It Hurts

 - It introduces unnecessary complexity by adding an intermediary layer between clients and the actual service or component.
 - The presence of a middleman can obscure the direct relationship between clients and the service they depend on. 
 - Middleman classes or methods often require maintenance themselves, even if they don't provide significant functionality.
 - Middlemen can introduce tighter coupling between clients and the underlying service or component.

## Refactor

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


## How To Fix It

If most of a method’s classes delegate to another class, [Remove Middle Man](.././../2.%20refactorings/remove-middle-man.md) is in order.