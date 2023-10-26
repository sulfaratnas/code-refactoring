# Temporary Field

The "Temporary Field" code smell occurs when a class or struct contains fields that are used only for a short duration, often for a specific operation or calculation.

## What It Looks Like

```
type Order struct {
    ID          int
    CustomerID  int
    TotalAmount float64

    // Temporary fields for calculating tax
    TaxRate    float64
    TaxAmount  float64
    IsTaxExempt bool
}

func (o *Order) CalculateTax() {
    // Calculate tax based on TotalAmount and TaxRate
    o.TaxAmount = o.TotalAmount * o.TaxRate
}

func main() {
    order := &Order{
        ID:          1,
        CustomerID:  1001,
        TotalAmount: 50.0,
        TaxRate:     0.1, // 10% tax rate
    }

    // Calculate tax
    order.CalculateTax()

    // Print tax amount
    fmt.Printf("Order #%d Tax Amount: $%.2f\n", order.ID, order.TaxAmount)
}

```

## Why It Hurts

These fields may not have a meaningful, long-term purpose in the class and can clutter the code.

## How To Fix It

Removed the temporary fields.

## Refactor

```
type Order struct {
    ID          int
    CustomerID  int
    TotalAmount float64
}

func NewOrder(id, customerID int, totalAmount float64) *Order {
    return &Order{
        ID:          id,
        CustomerID:  customerID,
        TotalAmount: totalAmount,
    }
}

func (o *Order) CalculateTaxRate(taxRate float64) float64 {
    return o.TotalAmount * taxRate
}

func main() {
    order := NewOrder(1, 1001, 50.0)
    taxRate := 0.1 // 10% tax rate

    // Calculate tax
    taxAmount := order.CalculateTaxRate(taxRate)

    // Print tax amount
    fmt.Printf("Order #%d Tax Amount: $%.2f\n", order.ID, taxAmount)
}
```

