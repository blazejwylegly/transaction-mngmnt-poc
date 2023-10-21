package main

import (
	"fmt"
	"github.com/blazejwylegly/orders-service/orders/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Orders-service started")

	baseRouter := mux.NewRouter()
	web.InitOrderRouting(baseRouter)
	log.Fatal(http.ListenAndServe("localhost:8080", baseRouter))
}
