package main

import (
	"fmt"

	"github.com/carlosdamazio/lotrsdk/pkg/api"
)

func main() {
	client := api.New()

	books, err := client.Book.List()
	if err != nil {
		panic(err)
	}

	for _, book := range books {
		fmt.Printf("ID: %s, Name: %s\n", book.ID, book.Name)
	}

	book, err := client.Book.GetByID("5cf5805fb53e011a64671582")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Book: %s", book.Name)

	chapters, err := client.Book.GetAllChaptersByBookID("5cf5805fb53e011a64671582")
	for _, chapter := range chapters {
		fmt.Printf("ID: %s, Name: %s\n", chapter.ID, chapter.ChapterName)
	}
}
