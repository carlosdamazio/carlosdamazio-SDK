package main

import (
	"fmt"

	"github.com/carlosdamazio/carlosdamazio_SDK/pkg/api"
)

func main() {
	client := api.New()

	chapters, err := client.Chapter.List()
	if err != nil {
		panic(err)
	}

	for _, chapter := range chapters {
		fmt.Printf("ID: %s, Name: %s\n", chapter.ID, chapter.ChapterName)
	}
}
