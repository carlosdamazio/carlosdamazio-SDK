package book

import (
	"net/http"

	"github.com/carlosdamazio/carlosdamazio_SDK/internal/models"
	"github.com/carlosdamazio/carlosdamazio_SDK/internal/service"
)

type (
	Service struct {
		Client *http.Client
	}
)

func New() *Service {
	return &Service{}
}

func (s *Service) WithHTTPClient(client *http.Client) *Service {
	s2 := s
	s2.Client = client
	return s2
}

func (s *Service) List() ([]*models.Book, error) {
	res, err := service.List[models.Book](service.BaseUrl+"/book", s.Client, "")
	if err != nil {
		return nil, err
	}
	return any(res).([]*models.Book), nil
}

func (s *Service) GetByID(id string) (*models.Book, error) {
	res, err := service.List[models.Book](service.BaseUrl+"/book/"+id, s.Client, "")
	if err != nil {
		return nil, err
	}

	return any(res).([]*models.Book)[0], nil
}

func (s *Service) GetAllChaptersByBookID(id string) ([]*models.BookChapter, error) {
	res, err := service.List[models.BookChapter](service.BaseUrl+"/book/"+id+"/chapter", s.Client, "")
	if err != nil {
		return nil, err
	}
	return any(res).([]*models.BookChapter), nil
}
