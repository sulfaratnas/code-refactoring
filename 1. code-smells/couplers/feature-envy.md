# Feature Envy

## What It Looks Like

A method accesses the data of another object more than its own data. It occurs when a method in one class seems more interested in the fields or methods of another class rather than its own. It often indicates that the method should belong to the class it is envious of.


```
// Wallet represents a simple wallet with a balance.
type Wallet struct {
	balance float64
}

// Deposit allows adding money to the wallet.
func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

// Withdraw allows withdrawing money from the wallet.
func (w *Wallet) Withdraw(amount float64) {
	w.balance -= amount
}

// Balance returns the wallet balance.
func (w *Wallet) Balance() float64 {
	return w.balance
}

// EnviousFunction is an example of a function in another class
// that seems envious of the Wallet's balance.
func EnviousFunction(wallet *Wallet) float64 {
	return wallet.Balance() * 2 // Envious of Wallet's balance
}

func main() {
	wallet := &Wallet{}
	wallet.Deposit(100.0)
	balance := EnviousFunction(wallet) // Feature Envy
	fmt.Printf("Wallet Balance: $%.2f\n", balance)
}
```

## Why It Hurts

Feature envy is a code smell that points to a potential design issue, where methods are not residing in the most appropriate classes. Identifying and addressing feature envy through refactoring can lead to more maintainable, modular, and understandable code.


## How To Fix It

If a method clearly should be moved to another place, use [Move Method](.././../refactorings/move-method.md).
If only part of a method accesses the data of another object, use [Extract Method](.././../refactorings/extract-method.md).

## Payoff
- Less code duplication (if the data handling code is put in a central place).
- Better code organization (methods for handling data are next to the actual data).