package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/carlosdamazio/lotrsdk/internal/models"
)

type ListDocs[T any] struct {
	Docs   []*T `json:"docs"`
	Total  int  `json:"total"`
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
	Page   int  `json:"page"`
	Pages  int  `json:"pages"`
}

func List[T models.Book |
	models.Movie |
	models.BookChapter |
	models.Character |
	models.Chapter |
	models.Quote](url string, client *http.Client, auth string) ([]*T, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth))

	resp, err := client.Do(req)
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respDocs := ListDocs[T]{}

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bodyBytes, &respDocs); err != nil {
		return nil, err
	}

	return respDocs.Docs, nil
}
