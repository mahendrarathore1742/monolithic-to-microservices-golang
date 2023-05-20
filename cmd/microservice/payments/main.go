package main

import (
	"Apis/go/pkg/mod/github.com/spf13/cobra@v1.1.3/cobra/cmd"
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Payment Microservice start")

	defer log.Println("closing payments microservice")

	ctx := cmd.Context()

	paymentsInterface := createPaymentsMicoservice()

	if err := paymentsInterface.Run(ctx); err != nil {
		panic(err)

	}

}

func createPaymentsMicoservice() amqp.paymentsInterface {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	paymentsService := payments_app.NewPaymentsService(
		payments_infra_orders.NewHTTPClint(os.Getenv("SHOP_ORDERS_SERVICE_ADDER")),
	)

	paymentsInterface, err := amqp.NewPaymentsInterface(
		fmt.Sprint("ampq://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDER_TO_PAY_QUEUE"),
		paymentsService,
	)

	if err != nil {
		panic(err)
	}

	return paymentsInterface
}
