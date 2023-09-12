package main

import (
	"fmt"
)

type Customer struct {
	Name string
}

func NewCustomer(name string) *Customer {
	return &Customer{Name: name}
}

func (c *Customer) ReserveSeat(seat string) {
	fmt.Printf("%s reserved seat %s\n", c.Name, seat)
}

type Reservation struct {
	Customer *Customer
}

func NewReservation(customer *Customer) *Reservation {
	return &Reservation{Customer: customer}
}

func (r *Reservation) MakeReservation(seat string) {
	//delegate to customer to do an action
	r.Customer.ReserveSeat(seat)
}

func main() {
	customer := NewCustomer("Alice")
	reservation := NewReservation(customer)

	reservation.MakeReservation("A1")
}
