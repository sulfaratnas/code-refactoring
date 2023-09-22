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
	Customer *Customer
	// Other reservation-related fields
}

func NewReservation(customer *Customer) *Reservation {
	return &Reservation{Customer: customer}
}

type Cinema struct {
	Name string
	// Other cinema-related fields
}

func NewCinema(name string) *Cinema {
	return &Cinema{Name: name}
}

func main() {
	// Creating objects
	customer := NewCustomer("Alice", "alice@example.com", "Gold")
	reservation := NewReservation(customer)
	cinema := NewCinema("Cineplex")

	// Accessing properties through a message chain
	customerName := reservation.Customer.Name
	cinemaName := cinema.Name

	fmt.Printf("Customer Name: %s\n", customerName)
	fmt.Printf("Cinema Name: %s\n", cinemaName)
}
