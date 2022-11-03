package main

import (
	"fmt"

	"github.com/carlosdamazio/lotrsdk/pkg/api"
)

func main() {
	client := api.New("SDzi1-GkN2i4hUA9R7Iy")

	chapters, err := client.Chapter.List()
	if err != nil {
		panic(err)
	}

	for _, chapter := range chapters {
		fmt.Printf("ID: %s, Name: %s\n", chapter.ID, chapter.ChapterName)
	}
}
