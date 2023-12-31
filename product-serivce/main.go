package main

import (
	"fmt"
	"github.com/blazejwylegly/product-service/config"
	"github.com/blazejwylegly/product-service/product/data"
	"github.com/blazejwylegly/product-service/product/service"
	"github.com/blazejwylegly/product-service/product/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const configFileName = "config.yaml"

func main() {
	appConfig := config.New(configFileName)
	productRepo := data.NewProductRepository(*appConfig.GetDatabaseConfig())
	productService := service.NewProductService(&productRepo)
	router := mux.NewRouter()

	web.InitProductApi(router, productService)
	web.InitDevApi(router, *appConfig)

	serverUrl := fmt.Sprintf("%s:%s", appConfig.Server.Host, appConfig.Server.Port)
	err := http.ListenAndServe(serverUrl, router)
	if err != nil {
		log.Fatal(err)
	}
}
