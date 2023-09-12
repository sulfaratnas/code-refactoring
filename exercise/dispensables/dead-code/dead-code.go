package main

import (
	"fmt"
)

type Movie struct {
	Title     string
	ShowTimes []string
}

type Booking struct {
	MovieTitle string
	ShowTime   string
	NumTickets int
	CustomerID int
}

// This function is no longer used in the application
func cancelBooking(booking *Booking) {
	fmt.Printf("Booking for movie '%s' at %s canceled.\n", booking.MovieTitle, booking.ShowTime)
	// Implement cancellation logic (not shown here)
}

func main() {
	movie := &Movie{
		Title: "Avengers: Endgame",
		ShowTimes: []string{
			"2023-09-11 18:00",
			"2023-09-11 20:30",
		},
	}

	booking := &Booking{
		MovieTitle: "Avengers: Endgame",
		ShowTime:   "2023-09-11 18:00",
		NumTickets: 2,
		CustomerID: 123,
	}

	// In this example, the cancelBooking function is defined but not used anywhere in the application.
	// It's considered dead code and should be removed to clean up the codebase.

	fmt.Println("Booking confirmed for", booking.NumTickets, "tickets.")
}
