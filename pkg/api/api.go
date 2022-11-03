package api

import (
	"net/http"
	"os"

	"github.com/carlosdamazio/lotrsdk/internal/models"
	"github.com/carlosdamazio/lotrsdk/internal/service/book"
	"github.com/carlosdamazio/lotrsdk/internal/service/chapter"
	"github.com/carlosdamazio/lotrsdk/internal/service/character"
	"github.com/carlosdamazio/lotrsdk/internal/service/movie"
	"github.com/carlosdamazio/lotrsdk/internal/service/quote"
)

type (
	BookSvc interface {
		List() ([]*models.Book, error)
		GetByID(id string) (*models.Book, error)
		GetAllChaptersByBookID(id string) ([]*models.BookChapter, error)
	}
	MovieSvc interface {
		List() ([]*models.Movie, error)
		GetByID(id string) (*models.Movie, error)
		GetAllQuotesByMovieID(id string) ([]*models.Quote, error)
	}
	CharacterSvc interface {
		List() ([]*models.Character, error)
		GetByID(id string) (*models.Character, error)
		GetAllQuotesByCharacterID(id string) ([]*models.Quote, error)
	}
	QuoteSvc interface {
		List() ([]*models.Quote, error)
		GetByID(id string) (*models.Quote, error)
	}
	ChapterSvc interface {
		List() ([]*models.Chapter, error)
		GetByID(id string) (*models.Chapter, error)
	}

	Requester struct {
		Authorization string
		Book          BookSvc
		Movie         MovieSvc
		Character     CharacterSvc
		Quote         QuoteSvc
		Chapter       ChapterSvc
	}
)

func New() *Requester {
	authHeader := os.Getenv("GO_LOTR_API")
	client := &http.Client{}
	return &Requester{
		Book: book.New().WithHTTPClient(client),
		Movie: movie.New().WithHTTPClient(client).WithAuthorizationHeader(authHeader),
		Character: character.New().WithHTTPClient(client).WithAuthorizationHeader(authHeader),
		Quote: quote.New().WithHTTPClient(client).WithAuthorizationHeader(authHeader),
		Chapter: chapter.New().WithHTTPClient(client).WithAuthorizationHeader(authHeader),
	}
}

func (r *Requester) WithAuthorizationHeader(authorization string) *Requester {
	r2 := r
	r2.Authorization = authorization
	return r2
}
