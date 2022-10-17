package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(4, 30)),
		validation.Field(&u.Password, validation.Required, validation.Length(4, 30)),
	)
}
