package main

import (
	"fmt"

	"github.com/carlosdamazio/carlosdamazio_SDK/pkg/api"
)

func main() {
	client := api.New()

	quotes, err := client.Quote.List()
	if err != nil {
		panic(err)
	}

	for _, quote := range quotes {
		fmt.Printf("ID: %s, Name: %s\n", quote.ID, quote.Dialog)
	}
}
