package main

import "fmt"

type Reservation struct {
	CustomerName string
	// Other reservation-related fields
}

func NewReservation(customerName string, _ int) *Reservation {
	// Initialize and return a new reservation
	return &Reservation{CustomerName: customerName}
}

func main() {
	customerName := "Alice"
	// The second parameter is not used in the NewReservation function
	reservation := NewReservation(customerName, 42)
	fmt.Println("Customer Name:", reservation.CustomerName)
}
