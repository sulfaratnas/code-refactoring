# Replace Parameter With Method Call

The "Replace Parameter with Method Call" refactoring technique involves replacing a method or function parameter with a call to a method or function that returns the same value. This can make the code more maintainable and reduce the need to pass parameters around.

## Example

### Before

```
func main() {
	basePrice := quantity * itemPrice;
    seasonDiscount := this.getSeasonalDiscount();
    fees := this.getFees();
    finalPrice := discountedPrice(basePrice, seasonDiscount, fees);
    fmt.Println(finalPrice)
}

func discountedPrice(basePrice, seasonDiscount, fees float64) {
    // some action
}

func getSeasonalDiscount() {
    // some action
}

func getFees() {
    // some action
}
```

### After

```
func main() {
	basePrice := quantity * itemPrice;
    finalPrice := discountedPrice(basePrice, seasonDiscount, fees);
    fmt.Println(finalPrice)
}

func discountedPrice(basePrice) {
    seasonDiscount := getSeasonalDiscount();
    fees := getFees();
    // some action
}

func getSeasonalDiscount() {
    // some action
}

func getFees() {
    // some action
}
```