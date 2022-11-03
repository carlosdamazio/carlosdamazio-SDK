package main

import (
	"fmt"

	"github.com/carlosdamazio/carlosdamazio_SDK/pkg/api"
)

func main() {
	client := api.New()

	characters, err := client.Character.List()
	if err != nil {
		panic(err)
	}

	for _, character := range characters {
		fmt.Printf("ID: %s, Name: %s\n", character.ID, character.Name)
	}

	character, err := client.Character.GetByID("5cd99d4bde30eff6ebccfc15")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Book: %s", character.Name)

	quotes, err := client.Character.GetAllQuotesByCharacterID("5cd99d4bde30eff6ebccfc15")
	for _, quote := range quotes {
		fmt.Printf("ID: %s, Name: %s\n", quote.ID, quote.Dialog)
	}
}
