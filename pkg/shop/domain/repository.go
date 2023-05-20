package products

import (
	"errors"
)

var ErrNotFound = errors.New("Product not found")

type Repository interface {
	Save(*products) error
	ByID(ID) (*products, error)
}
