package http

import (
	"Apis/Desktop/project/online_shoping_microservice/pkg/common/price"
	"Apis/go/pkg/mod/github.com/go-chi/chi@v4.0.1+incompatible"
	"net/http"
)

func AddRouter(routers *chi.Mux, repo products_domain.Repository) {

	resource := productsResource{repo}
	routers.Get("products/{id}", resource.Get)
}

type productsResource struct {
	repo products_domain.Repository
}

type PriceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func PriceViewFromPrice(p price.Price) PriceView {

	return PriceView{p.Cents(), p.Currency()}
}

func (p productsResource) Get(w http.ResponseWriter, r *http.Request) {

	product, err := p.repo.ByID(products_domain.ID(chi.URLParam(r, "id")))

	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
	}

	render.Respond{
		string(product.ID()),
		product.Name(),
		product.Description(),
		PriceViewFromPrice(product.Price()),
	}
}
