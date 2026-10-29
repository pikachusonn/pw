package domain

import "time"

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Plant struct {
	Id          string    `json:"id"`
	Position    Position  `json:"position"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Types       []string  `json:"types"`
	Images      []string  `json:"images"`
}

func NewPlant(id string, position Position, name string, description string, types []string, images []string) *Plant {
	return &Plant{
		Id:          id,
		Position:    position,
		Name:        name,
		Description: description,
		Types:       types,
		Images:      images,
		CreatedAt:   time.Now(),
	}
}
