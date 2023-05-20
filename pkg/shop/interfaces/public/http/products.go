package http

import (
	"Apis/Desktop/project/online_shoping_microservice/pkg/common/price"
	"Apis/Desktop/project/online_shoping_microservice/pkg/shop/domain/products"
	"net/http"

	"Apis/go/pkg/mod/github.com/go-chi/chi@v4.0.1+incompatible"
)

func AddRouter(router *chi.Mux, productReadModels productReadModels) {

	resource := productsResourse{productReadModels}

	router.Get("/products", resource.GetAll)
}

type productsResourse struct {
	readModel productReadModels
}

type productReadModels interface {
	AllProducts() ([]products.Product, error)
}

type productView struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	price       priceView `json:"price"`
}

type priceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func priceViewFormPrice(p price.Price) priceView {

	return priceView{p.Cents(), p.Currency()}
}

func (p productsResourse) GetAll(w http.ResponseWriter, r *http.Request) {

	products, err := p.readModel.AllProducts()

	if err != nil {
		_ = render.Render(w, common_http.ErrInternal(err))
		return
	}

	view := []productView{}

	for _, product := range products {

		view = append(view, productView{
			string(product.ID()),
			product.name(),
			product.Description(),
			priceViewFormPrice(product.Price()),
		})
	}

	render.Respond(w, r, view)

}
