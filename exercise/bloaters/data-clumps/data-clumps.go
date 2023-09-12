package main

import (
	"fmt"
	"time"
)

type Screening struct {
	MovieTitle string
	StartDate  time.Time
	EndDate    time.Time
}

func main() {
	// Example 1: Creating a Screening
	startDate1 := time.Date(2023, time.September, 11, 15, 0, 0, 0, time.UTC)
	endDate1 := time.Date(2023, time.September, 11, 17, 0, 0, 0, time.UTC)
	screening1 := Screening{
		MovieTitle: "Avengers: Endgame",
		StartDate:  startDate1,
		EndDate:    endDate1,
	}

	// Example 2: Creating another Screening
	startDate2 := time.Date(2023, time.September, 12, 14, 30, 0, 0, time.UTC)
	endDate2 := time.Date(2023, time.September, 12, 16, 30, 0, 0, time.UTC)
	screening2 := Screening{
		MovieTitle: "Spider-Man: No Way Home",
		StartDate:  startDate2,
		EndDate:    endDate2,
	}

	// Printing Screening Details
	fmt.Println("Screening 1 Details:")
	fmt.Println("Movie Title:", screening1.MovieTitle)
	fmt.Println("Start Date:", screening1.StartDate.Format("2006-01-02 15:04:05"))
	fmt.Println("End Date:", screening1.EndDate.Format("2006-01-02 15:04:05"))

	fmt.Println("\nScreening 2 Details:")
	fmt.Println("Movie Title:", screening2.MovieTitle)
	fmt.Println("Start Date:", screening2.StartDate.Format("2006-01-02 15:04:05"))
	fmt.Println("End Date:", screening2.EndDate.Format("2006-01-02 15:04:05"))
}
