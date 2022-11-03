package main

import (
	"fmt"

	"github.com/carlosdamazio/lotrsdk/pkg/api"
)

func main() {
	client := api.New("SDzi1-GkN2i4hUA9R7Iy")

	movies, err := client.Movie.List()
	if err != nil {
		panic(err)
	}

	for _, movie := range movies {
		fmt.Printf("ID: %s, Name: %s\n", movie.ID, movie.Name)
	}

	movie, err := client.Movie.GetByID("5cd95395de30eff6ebccde56")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Movie: %s\n", movie.Name)

	quotes, err := client.Movie.GetAllQuotesByMovieID("5cd95395de30eff6ebccde5c")
	if err != nil {
		panic(err)
	}

	for _, quote := range quotes {
		fmt.Printf("ID: %s, Quote: %s\n", quote.ID, quote.Dialog)
	}
}
