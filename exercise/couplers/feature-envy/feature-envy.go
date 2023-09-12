package main

import (
	"fmt"
)

type Customer struct {
	Name       string
	Email      string
	MemberType string
}

func NewCustomer(name, email, memberType string) *Customer {
	return &Customer{Name: name, Email: email, MemberType: memberType}
}

type Reservation struct {
	Customer  *Customer
	SeatCount int
	// Other reservation-related fields
}

func NewReservation(customer *Customer, seatCount int) *Reservation {
	return &Reservation{Customer: customer, SeatCount: seatCount}
}

func (r *Reservation) CalculateDiscountedPrice() float64 {
	var discountFactor float64

	// Determine the discount factor based on the customer's member type
	if r.Customer.MemberType == "Gold" {
		discountFactor = 0.9
	} else if r.Customer.MemberType == "Silver" {
		discountFactor = 0.95
	} else {
		discountFactor = 1.0
	}

	// Calculate and return the discounted price
	return float64(r.SeatCount) * 10.0 * discountFactor
}

func main() {
	customer := NewCustomer("Alice", "alice@example.com", "Gold")
	reservation := NewReservation(customer, 3)

	discountedPrice := reservation.CalculateDiscountedPrice()

	fmt.Println("Customer Reservation:")
	fmt.Printf("Customer Name: %s\n", customer.Name)
	fmt.Printf("Discounted Price: %.2f\n", discountedPrice)
}
