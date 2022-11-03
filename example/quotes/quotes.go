package main

import (
	"fmt"

	"github.com/carlosdamazio/lotrsdk/pkg/api"
)

func main() {
	client := api.New("SDzi1-GkN2i4hUA9R7Iy")

	quotes, err := client.Quote.List()
	if err != nil {
		panic(err)
	}

	for _, quote := range quotes {
		fmt.Printf("ID: %s, Name: %s\n", quote.ID, quote.Dialog)
	}
}

