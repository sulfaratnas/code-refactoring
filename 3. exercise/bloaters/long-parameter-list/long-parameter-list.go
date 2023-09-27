package main

import (
	"fmt"
)

type Reservation struct {
	CustomerName string
	Seats        []string
}

type Movie struct {
	Title     string
	ShowTimes []string
}

func makeReservation(customerName string, customerAge int, seats []string, movieTitle string, movieGenre string, showTime string, ticketPrice float64, paymentMethod string) {
	// Logic to create a reservation
	fmt.Printf("Reservation made for %s\n", customerName)
	fmt.Printf("Customer Age %s\n", customerAge)
	fmt.Printf("Seats reserved: %v\n", seats)
	fmt.Printf("Movie: %s\n", movieTitle)
	fmt.Printf("Genre: %s\n", movieGenre)
	fmt.Printf("Showtime: %s\n", showTime)
	fmt.Printf("Ticket price: %.2f\n", ticketPrice)
	fmt.Printf("Payment method: %s\n", paymentMethod)
}

func main() {
	seats := []string{"A1", "A2", "A3"}
	movieTitle := "Avengers: Endgame"
	movieGenre := "Sci-fi"
	showTime := "2023-09-11 18:00"
	ticketPrice := 12.5
	paymentMethod := "Credit Card"

	// long paramater list
	makeReservation("Alice", 15, seats, movieTitle, movieGenre, showTime, ticketPrice, paymentMethod)
}
