package main

import (
	"fmt"
)

type Movie struct {
	Title       string
	Genre       string
	ReleaseYear int
}

func NewMovie(title, genre string, releaseYear int) *Movie {
	return &Movie{Title: title, Genre: genre, ReleaseYear: releaseYear}
}

type Reservation struct {
	MovieTitle  string
	MovieGenre  string
	Seats       []string
	ReleaseYear int
}

func NewReservation(movieTitle, movieGenre string, seats []string, releaseYear int) *Reservation {
	return &Reservation{
		MovieTitle:  movieTitle,
		MovieGenre:  movieGenre,
		Seats:       seats,
		ReleaseYear: releaseYear,
	}
}

func PrintReservationDetails(reservation *Reservation) {
	fmt.Println("Reservation Details:")
	fmt.Println("Movie Title:", reservation.MovieTitle)
	fmt.Println("Movie Genre:", reservation.MovieGenre)
	fmt.Println("Reserved Seats:", reservation.Seats)
	fmt.Println("Release Year:", reservation.ReleaseYear)
}

func main() {
	movie := NewMovie("Inception", "Science Fiction", 2010)
	reservation := NewReservation(movie.Title, movie.Genre, []string{"A1", "A2", "A3"}, movie.ReleaseYear)

	// Later, a requirement change occurs
	movie.Title = "The Matrix"

	// Now, we need to update the reservation as well
	reservation.MovieTitle = movie.Title

	PrintReservationDetails(reservation)
}
