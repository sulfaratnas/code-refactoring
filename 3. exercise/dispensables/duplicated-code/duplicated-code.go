package main

import (
	"fmt"
)

type Reservation struct {
	CustomerName string
	Seats        []string
}

func printReservationDetails(reservation Reservation) {
	fmt.Println("Customer Name:", reservation.CustomerName)
	fmt.Println("Reserved Seats:")
	for _, seat := range reservation.Seats {
		fmt.Println(seat)
	}
}

func main() {
	// Duplicate code for creating & printing reservation details
	reservation1 := Reservation{
		CustomerName: "Alice",
		Seats:        []string{"A1", "A2", "A3"},
	}
	fmt.Println("Reservation 1 Details:", reservation1)
	printReservationDetails(reservation1)

	reservation2 := Reservation{
		CustomerName: "Bob",
		Seats:        []string{"B1", "B2"},
	}
	fmt.Println("Reservation 2 Details:", reservation2)
	printReservationDetails(reservation2)
}
