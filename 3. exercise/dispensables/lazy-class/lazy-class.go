package main

import "fmt"

type Reservation struct {
	CustomerName string
}

func NewReservation(customerName string) *Reservation {
	return &Reservation{CustomerName: customerName}
}

func (r *Reservation) PrintCustomerName() {
	fmt.Println("Customer Name:", r.CustomerName)
}

func main() {
	reservation := NewReservation("Alice")
	reservation.PrintCustomerName()
}
