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

func main() {
	movie := NewMovie("Inception", "Science Fiction", 2010)

	switch movie.Genre {
	case "Action":
		fmt.Println("Action movie")
	case "Drama":
		fmt.Println("Drama movie")
	case "Science Fiction":
		fmt.Println("Science Fiction movie")
	default:
		fmt.Println("Unknown genre")
	}

	switch {
	case movie.ReleaseYear >= 2000 && movie.ReleaseYear < 2010:
		fmt.Println("Released in the 2000s")
	case movie.ReleaseYear >= 2010 && movie.ReleaseYear < 2020:
		fmt.Println("Released in the 2010s")
	case movie.ReleaseYear >= 2020:
		fmt.Println("Released in the 2020s")
	default:
		fmt.Println("Release year unknown")
	}
}
