# Replace Parameter With Method Call

The "Replace Parameter with Method Call" refactoring technique involves replacing a method or function parameter with a call to a method or function that returns the same value. This can make the code more maintainable and reduce the need to pass parameters around.

## Example

### Before

```
func main() {
    basePrice := quantity * itemPrice;
    seasonDiscount := getSeasonalDiscount(basePrice);
    fees := getFees(basePrice);
    finalPrice := discountedPrice(basePrice, seasonDiscount, fees); 
}

func getSeasonalDiscount(basePrice float64) {
    // some logic here
}

func getFees(basePrice float64) {
    // some logic here
}
```

### After

```
func main() {
	basePrice := quantity * itemPrice;
    finalPrice := discountedPrice(basePrice);
}

func discountedPrice(basePrice) {
    seasonDiscount := getSeasonalDiscount(basePrice);
    fees := getFees(basePrice);
    // some action
}

func getSeasonalDiscount(basePrice float64) {
    // some logic here
}

func getFees(basePrice float64) {
    // some logic here
}
```