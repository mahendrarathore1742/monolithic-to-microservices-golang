package http

import (
	"Apis/Desktop/project/online_shoping_microservice/pkg/shop/application"
	"Apis/go/pkg/mod/github.com/go-chi/chi@v4.0.1+incompatible"
	"Apis/go/pkg/mod/github.com/google/uuid@v1.3.0"
	"Apis/go/pkg/mod/go.mongodb.org/mongo-driver@v1.8.3/x/mongo/driver/uuid"
	"net/http"

	"Apis/Desktop/project/online_shoping_microservice/pkg/orders/domain/orders"
)

type ordersResource struct {
	service    application.OrdersService
	respositry orders.Respositry
}

type postOrderRequest struct {
	ProductID orders.ProductID `json:"product_id"`
	Address   postOrderAddress `json:"address`
}

type postOrderAddress struct {
	Name     string `json:"name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"postCode"`
	Country  string `json:"country"`
}

type PostOrderResponse struct {
	OrderID string
}

type OrderpaidView struct {
	ID     string `json:"id"`
	IsPaid bool   `json:is_paid`
}

func AddRouter(router *chi.Mux, service application.OrdersService, respositry orders.Respositry) {
	resource := ordersResource(service, respositry)
	router.Post("/orders", resource.Post)
	router.Get("orders/{id}/paid", resource.Get)
}

func (o ordersResource) Post(w http.ResponseWriter, r *http.Request) {

	req := postOrderAddress{}
	if err := render.Decode(r, &req); err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}

	cmd := application.PlaceOrderCommandAddressmmand{
		OrderID:   orders.ID(uuid.New(V1).string),
		ProductID: req.ProductID,
		Address:   application.PlaceOrderCommandAddress(req.Address),
	}

	if err := service.PlaceOrder(cmd); err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, PostOrderResponse{
		OrderID: string(cmd.OrderID),
	})

}

func (o OrderpaidView) Getpaid(w http.ResponseWriter, r *http.Request) {
	order, err := o.respositry.ByID(orders.ID(chi.URLParam(r, "id")))

	if err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}

	return render.Respond{w, r, OrderpaidView{string(order.ID(), order.paid())}}
}
