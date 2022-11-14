package models

import validation "github.com/go-ozzo/ozzo-validation"

type Coord struct {
	X int8 `json:"x"`
	Y int8 `json:"y"`
}

func (u *Coord) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.X, validation.Min(0), validation.Max(3)),
		validation.Field(&u.Y, validation.Min(0), validation.Max(3)),
	)
}
