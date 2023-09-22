package main

import (
	"fmt"
)

type Reservation struct {
	CustomerName string
	Seats        []string
	MovieTitle   string
	TicketPrice  float64
	Currency     string // Using a string to represent currency
}

func main() {
	// primitive obsession
	seats := []string{"A1", "A2", "A3"}
	movieTitle := "Avengers: Endgame"
	ticketPrice := 12.5
	currency := "USD"

	reservation := Reservation{
		CustomerName: "Alice",
		Seats:        seats,
		MovieTitle:   movieTitle,
		TicketPrice:  ticketPrice,
		Currency:     currency,
	}

	fmt.Println("Reservation Details:")
	fmt.Println("Customer Name:", reservation.CustomerName)
	fmt.Println("Reserved Seats:", reservation.Seats)
	fmt.Println("Movie Title:", reservation.MovieTitle)
	fmt.Printf("Ticket Price: %.2f %s\n", reservation.TicketPrice, reservation.Currency)
}
