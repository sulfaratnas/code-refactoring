package main

import (
	"fmt"
)

type Reservation struct {
	CustomerName string   // Customer's name
	Seats        []string // List of reserved seats
	MovieTitle   string   // Title of the movie
	TicketPrice  float64  // Price per ticket
	Currency     string   // Currency code (e.g., USD)
}

func calculateTotalPrice(price float64, quantity int) float64 {
	// TODO: Improve this calculation logic
	// Calculate the total price by multiplying the price and quantity
	return price * float64(quantity)
}

func printReservationDetails(reservation Reservation) {
	// Print reservation details
	fmt.Println("Reservation Details:")
	fmt.Println("Customer Name:", reservation.CustomerName)
	fmt.Println("Reserved Seats:", reservation.Seats)
	fmt.Println("Movie Title:", reservation.MovieTitle)
	fmt.Printf("Ticket Price: %.2f %s\n", reservation.TicketPrice, reservation.Currency)
}

func main() {
	// Create a reservation
	seats := []string{"A1", "A2", "A3"} // List of available seats
	movieTitle := "Avengers: Endgame"   // Title of the movie
	ticketPrice := 12.5                 // Price per ticket
	currency := "USD"                   // Currency code (e.g., USD)

	reservation := Reservation{
		CustomerName: "Alice",     // Customer's name
		Seats:        seats,       // List of reserved seats
		MovieTitle:   movieTitle,  // Title of the movie
		TicketPrice:  ticketPrice, // Price per ticket
		Currency:     currency,    // Currency code (e.g., USD)
	}

	total := calculateTotalPrice(reservation.TicketPrice, len(reservation.Seats)) // Calculate total price

	fmt.Println("Customer Reservation:")
	printReservationDetails(reservation)                              // Print reservation details
	fmt.Printf("Total Price: %.2f %s\n", total, reservation.Currency) // Print total price
}
