package application

import (
	"Apis/Desktop/project/online_shoping_microservice/pkg/common/price"

	"Apis/Desktop/project/online_shoping_microservice/pkg/shop/domain/products"
	"errors"
)

type productReadModel interface {
	AllProducts() ([]products.Product, err)
}

type productsService struct {
	repo      products.Repository
	readModel productReadModel
}

func NewProductsService(repo products.Repository, readModel productReadModel) productsService {

	return productsService{repo, readModel}
}

func (s productsService) AllProducts() ([]products.Product, error) {

	return s.readModel.AllProducts()
}

type AddProductcommand struct {
	ID            int
	Name          string
	Description   string
	PriceCents    uint
	PriceCurrency string
}

func (s productsService) AddProduct(cmd AddProductcommand) error {

	price, err := price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)

	if err != nil {
		return errors.Wrap(err, "invalid product price")
	}

	p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)

	if err != nil {
		return errors.Wrap(err, "can not create product")
	}

	if err == s.repo.Save(p); err != nil {
		return errors.Wrap(err, "can not save products")
	}

	return nil
}
