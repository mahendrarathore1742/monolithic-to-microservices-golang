package orders

import (
	"Apis/Desktop/project/online_shoping_microservice/pkg/common/price"
	"errors"
)

type ProductID string

var ErrEmptyProductID = errors.New("empty product id")

type Product struct {
	id    ProductID
	name  string
	price price.Price
}

func NewProudct(id ProductID, name string, price price.Price) (Product, error) {

	if len(id) == 0 {
		return Product{}, ErrEmptyProductID
	}

	return Product{id, name, price}, nil
}

func (p Product) ID() ProductID {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Price() price.Price {
	return p.price
}
