package service

import "github.com/Drozd0f/ttto-go/models"

func (s *Service) GetGameByID(id int64) models.Game {
	return models.Game{
		ID: id,
		Field: [3][3]string{
			{"X", "0", " "},
			{" ", "X", " "},
			{" ", "0", " "},
		},
	}
}
