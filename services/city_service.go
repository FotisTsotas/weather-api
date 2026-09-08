package services

import "weather-api/repositories"

type CityService struct {
	repo *repositories.CityRepository
}

func NewCityService(repo *repositories.CityRepository) *CityService {
	return &CityService{repo: repo}
}

func (s *CityService) GetCities() ([]repositories.City, error) {
	cities, err := s.repo.GetCities()
	if err != nil {
		return nil, err
	}
	if len(cities) == 0 {
		return []repositories.City{}, nil
	}
	return cities, nil
}

func (s *CityService) GetCityByID(id int) (*repositories.City, error) {
	city, err := s.repo.GetCityByID(id)
	if err != nil {
		return nil, err
	}
	if city == nil {
		return nil, nil
	}
	return city, nil
}

func (s *CityService) CreateCity(name string, latitude, longitude float64) (*repositories.City, error) {
	city, err := s.repo.CreateCity(name, latitude, longitude)
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (s *CityService) UpdateCity(id int, name string, latitude, longitude float64) (*repositories.City, error) {
	city, err := s.repo.UpdateCity(id, name, latitude, longitude)
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (s *CityService) DeleteCity(id int) error {
	err := s.repo.DeleteCity(id)
	if err != nil {
		return err
	}
	return nil
}
