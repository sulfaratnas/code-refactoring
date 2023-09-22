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

func main() {
	reservations := []Reservation{
		{
			CustomerName: "Alice",
			Seats:        []string{"A1", "A2", "A3"},
		},
		{
			CustomerName: "Bob",
			Seats:        []string{"B1", "B2"},
		},
	}

	movies := []Movie{
		{
			Title: "Avengers: Endgame",
			ShowTimes: []string{
				"2023-09-11 18:00",
				"2023-09-11 20:30",
			},
		},
		{
			Title: "Inception",
			ShowTimes: []string{
				"2023-09-12 19:00",
				"2023-09-12 21:30",
			},
		},
	}

	// Nested loop to match reservations with movie showtimes
	for _, reservation := range reservations {
		for _, movie := range movies {
			for _, showTime := range movie.ShowTimes {
				// Check if the reservation's seat matches a seat for the movie and showtime
				for _, seat := range reservation.Seats {
					if seatAvailable(movie, showTime, seat) {
						fmt.Printf("%s reserved seat %s for %s at %s\n", reservation.CustomerName, seat, movie.Title, showTime)
					} else {
						fmt.Printf("%s's reservation for seat %s at %s is invalid\n", reservation.CustomerName, seat, showTime)
					}
				}
			}
		}
	}
}

func seatAvailable(movie Movie, showTime string, seat string) bool {
	// Complex logic to check seat availability (not shown here)
	// Return true if the seat is available for the specified movie and showtime
	return true
}
