package model

import "time"

// Product model
type Product struct {
	Name       string
	Desription string
	Image      []byte
	Price      int
	CreatedOn  time.Time
	UpdatedOn  time.Time
}
