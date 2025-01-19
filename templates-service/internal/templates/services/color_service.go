package services

import (
	"cv-templates-service/internal/templates/repositories"
	"errors"
)

type ColorService struct {
	ColorRepo repositories.ColorRepository
}

type Color struct {
	Title       string
	AccentColor string
}

func NewColorService(colorRepo repositories.ColorRepository) *ColorService {
	return &ColorService{
		ColorRepo: colorRepo,
	}
}

func (s *ColorService) GetColorsMap() (map[string]string, error) {
	return s.ColorRepo.GetColors()
}

func (s *ColorService) GetColor(title string) (*Color, error) {
	colors, _ := s.ColorRepo.GetColors()

	hex, ok := colors[title]
	if !ok {
		return nil, errors.New("color not found")
	}

	return &Color{
		Title:       title,
		AccentColor: hex,
	}, nil
}
