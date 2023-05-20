package main

import (
	"Apis/go/pkg/mod/github.com/go-chi/chi@v4.0.1+incompatible"
	"Apis/go/pkg/mod/github.com/spf13/cobra@v1.1.3/cobra/cmd"
	"log"
	"net/http"
	"os"
)

func main() {

	log.Println("string shop microservice")

	ctx := cmd.Context()

	r := createShopMicroservic()

	server := &http.Server{Addr: os.Getenv("SHOP_SERVICE_BIND_ADDR"), Handler: r}

	go func() {

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}

	}()

	<-ctx.Done()

	log.Println("Closing shop microservice")
	if err := server.Close(); err != nil {
		panic(err)
	}

}

func createShopMicroservic() *chi.Mux {
	shopProductRepo := shop_infra_product.NewMemoryRepositry()

	r := cmd.CreateRouter()

	shop_interface_public_http.AddRoutes(r, shopProductRepo)
	shop_interface_private_http.AddRoutes(r, shopProductRepo)

	return r
}
