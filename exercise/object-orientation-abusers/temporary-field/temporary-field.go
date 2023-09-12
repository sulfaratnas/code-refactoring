package main

import (
	"fmt"
)

type Reservation struct {
	CustomerName string
	Seats        []string
	MovieTitle   string
	TicketPrice  float64
	Currency     string
	TotalPrice   float64 // Temporary field for storing the total price
}

func NewReservation(customerName string, seats []string, movieTitle string, ticketPrice float64, currency string) *Reservation {
	return &Reservation{
		CustomerName: customerName,
		Seats:        seats,
		MovieTitle:   movieTitle,
		TicketPrice:  ticketPrice,
		Currency:     currency,
	}
}

func (r *Reservation) CalculateTotalPrice() {
	// Calculate and store the total price in the temporary field
	r.TotalPrice = r.TicketPrice * float64(len(r.Seats))
}

func main() {
	reservation := NewReservation("Alice", []string{"A1", "A2", "A3"}, "Avengers: Endgame", 12.5, "USD")
	reservation.CalculateTotalPrice()

	fmt.Println("Customer Reservation:")
	fmt.Println("Customer Name:", reservation.CustomerName)
	fmt.Println("Reserved Seats:", reservation.Seats)
	fmt.Println("Movie Title:", reservation.MovieTitle)
	fmt.Printf("Ticket Price: %.2f %s\n", reservation.TicketPrice, reservation.Currency)
	fmt.Printf("Total Price: %.2f %s\n", reservation.TotalPrice, reservation.Currency)
}
