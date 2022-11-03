package movie

import (
	"net/http"

	"github.com/carlosdamazio/lotrsdk/internal/models"
	"github.com/carlosdamazio/lotrsdk/internal/service"
)

type (
	Service struct {
		AuthHeader string
		Client     *http.Client
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

func (s *Service) WithAuthorizationHeader(header string) *Service {
	s2 := s
	s2.AuthHeader = header
	return s2
}

func (s *Service) List() ([]*models.Movie, error) {
	res, err := service.List[models.Movie](service.BaseUrl+"/movie", s.Client, s.AuthHeader)
	if err != nil {
		return nil, err
	}
	return any(res).([]*models.Movie), nil
}

func (s *Service) GetByID(id string) (*models.Movie, error) {
	res, err := service.List[models.Movie](service.BaseUrl+"/movie/"+id, s.Client, s.AuthHeader)
	if err != nil {
		return nil, err
	}

	return any(res).([]*models.Movie)[0], nil
}

func (s *Service) GetAllQuotesByMovieID(id string) ([]*models.Quote, error) {
	res, err := service.List[models.Quote](service.BaseUrl+"/movie/"+id+"/quote", s.Client, s.AuthHeader)
	if err != nil {
		return nil, err
	}
	return any(res).([]*models.Quote), nil
}
