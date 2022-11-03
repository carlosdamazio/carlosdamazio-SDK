package character

import (
	"net/http"

	"github.com/carlosdamazio/carlosdamazio_SDK/internal/models"
	"github.com/carlosdamazio/carlosdamazio_SDK/internal/service"
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

func (s *Service) List() ([]*models.Character, error) {
	res, err := service.List[models.Character](service.BaseUrl+"/character", s.Client, s.AuthHeader)
	if err != nil {
		return nil, err
	}
	return any(res).([]*models.Character), nil
}

func (s *Service) GetByID(id string) (*models.Character, error) {
	res, err := service.List[models.Character](service.BaseUrl+"/character/"+id, s.Client, s.AuthHeader)
	if err != nil {
		return nil, err
	}

	return any(res).([]*models.Character)[0], nil
}

func (s *Service) GetAllQuotesByCharacterID(id string) ([]*models.Quote, error) {
	res, err := service.List[models.Quote](service.BaseUrl+"/character/"+id+"/quote", s.Client, s.AuthHeader)
	if err != nil {
		return nil, err
	}
	return any(res).([]*models.Quote), nil
}
