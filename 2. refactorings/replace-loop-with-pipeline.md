# Replace Loop With Pipeline

Separate data transformation from side effects by replacing
an imperative loop with a sequence of functional operations.

## Example

### Before

```
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result := []int{}

	for _, num := range numbers {
		result = append(result, num*2)
	}

	fmt.Println(result)
}
```

### After

```
func doubleNumbers(numbers []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, num := range numbers {
			ch <- num * 2
		}
	}()
	return ch
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	pipelineResult := []int{}

	for doubled := range doubleNumbers(numbers) {
		pipelineResult = append(pipelineResult, doubled)
	}

	fmt.Println("Using Pipeline:")
	fmt.Println(pipelineResult)
}
```