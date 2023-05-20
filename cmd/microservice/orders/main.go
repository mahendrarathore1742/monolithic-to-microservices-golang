package main

import (
	"Apis/go/pkg/mod/github.com/go-chi/chi@v4.0.1+incompatible"
	"Apis/go/pkg/mod/github.com/spf13/cobra@v1.1.3/cobra/cmd"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Staring the order microservice")

	ctx := cmd.Context()

	r, closeFn := createOrderMicroservic()

	defer closeFn()

	server := &http.Server{Addr: os.Getenv("SHOP_ORDERS_SERVICE_BIND_ADDR"), Handler: r}

	go func() {

		if err := server.ListenAndServe(); err != http.ErrServerClosed {

			panic(err)
		}()

		<-ctx.Done()

		log.Println("closing order microservice")

		if err := server.Close(); err != nil {
			panic(err)
		}

	}

}

func createOrderMicroservic() (router *chi.Mux, closeFn func()) {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	shopHTTPClint := orders_infra_product.NewHTTPClint(os.Getenv("SHOP_PRODUCT_SERVICE_ADDR"))

	r := cmd.CreateRouter()

	orders_public_http.AddRoutes(r, ordersService, ordersRepo)
	orders_private_http.AddRoutes(r, ordersService, ordersRepo)


	return r,func() {
		
	}

}
