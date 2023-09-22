package main

import (
	"fmt"
	"time"
)

type Movie struct {
	Title     string
	ShowTimes []time.Time
}

type Booking struct {
	MovieTitle       string
	ShowTime         time.Time
	NumTickets       int
	CustomerID       int
	ConfirmationCode string
}

func bookTickets(movie *Movie, showTime time.Time, numTickets int, customerID int) *Booking {
	// Step 1: Check if the showtime is valid
	validShowTime := false
	for _, st := range movie.ShowTimes {
		if st == showTime {
			validShowTime = true
			break
		}
	}
	if !validShowTime {
		fmt.Println("Invalid showtime")
		return nil
	}

	// Step 2: Check if there are enough available tickets
	availableTickets := 100 // Assume 100 tickets available for simplicity
	if numTickets > availableTickets {
		fmt.Println("Not enough tickets available")
		return nil
	}

	// Step 3: Generate a confirmation code
	confirmationCode := generateConfirmationCode()

	// Step 4: Create the booking
	booking := &Booking{
		MovieTitle:       movie.Title,
		ShowTime:         showTime,
		NumTickets:       numTickets,
		CustomerID:       customerID,
		ConfirmationCode: confirmationCode,
	}

	// Step 5: Deduct the booked tickets from available tickets
	availableTickets -= numTickets

	// Step 6: Save the booking information (not shown here)

	return booking
}

func generateConfirmationCode() string {
	// Generate a random confirmation code (not shown here)
	return "ABC123"
}

func main() {
	movie := &Movie{
		Title: "Avengers: Endgame",
		ShowTimes: []time.Time{
			time.Date(2023, time.September, 11, 18, 0, 0, 0, time.UTC),
			time.Date(2023, time.September, 11, 20, 30, 0, 0, time.UTC),
		},
	}

	booking := bookTickets(movie, time.Date(2023, time.September, 11, 18, 0, 0, 0, time.UTC), 3, 123)

	if booking != nil {
		fmt.Println("Booking successful!")
		fmt.Printf("Confirmation Code: %s\n", booking.ConfirmationCode)
	}
}
