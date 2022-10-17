package models

type Field [3][3]string

type Game struct {
	ID    int64 `json:"id"`
	Field Field `json:"field"`
}
