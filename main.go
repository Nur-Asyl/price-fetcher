package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/Nur-Asyl/pricefetcher/client"
)

func main() {
	client := client.New("http://localhost:3000")

	price, err := client.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", price)

	return

	listenAddr := flag.String("http://localhost", ":3000", "listen address the service is running")
	flag.Parse()

	svc := NewLogginService(NewMetricService(&priceFetcher{}))

	server := NewJSONService(*listenAddr, svc)

	server.Run()

}
