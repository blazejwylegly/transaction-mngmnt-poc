package models

import "github.com/google/uuid"

type Product struct {
	ProductId         uuid.UUID
	Name              string
	Description       string
	Price             float64
	QuantityAvailable uint16
}
