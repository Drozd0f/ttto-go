package service

import (
	"github.com/Drozd0f/ttto-go/models"
)

func (s *Service) Reg(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	return nil
}
