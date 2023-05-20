package products

import "Apis/Desktop/project/online_shoping_microservice/pkg/shop/infrastructure/products"

type MemoryRepository struct {
	products []products.Product
}

func NewMemoryRepositry() *MemoryRepository {
	return &MemoryRepository{
		[]products.Product{},
	}
}

func (m *MemoryRepository) Save(productTosave *products.Product) error {

	for i, p := range m.products {
		if p.ID() == productTosave.ID() {
			m.products[i] = *productTosave
			return nil
		}
	}

	m.products = append(m.products, productTosave)
	return nil
}

func (m MemoryRepository) byID(id products.ID) (*products.Product, error) {

	for _, p := range m.products {

		if p.ID == id {
			return &p, nil
		}
	}

	return nil, products.ErrNotFound

}

func (m MemoryRepository) AllProdcuts() ([]products.Product, error) {
	return m.products, nil
}
