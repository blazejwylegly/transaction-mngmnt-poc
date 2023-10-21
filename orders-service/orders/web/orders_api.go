package web

import (
	"encoding/json"
	"github.com/blazejwylegly/orders-service/orders/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type OrderApi struct {
	router *mux.Router
}

func InitOrderRouting(router *mux.Router) *OrderApi {
	ordersRouter := router.PathPrefix("/order").Subrouter()
	api := &OrderApi{
		router: ordersRouter,
	}
	api.initializeMappings()
	return api
}

func (api OrderApi) initializeMappings() {
	api.router.HandleFunc("", api.handleGetOrderMapping()).Methods("GET")
	api.router.HandleFunc("", api.handleCreateOrderMapping()).Methods("POST")
}

func (api OrderApi) handleGetOrderMapping() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		encoder := json.NewEncoder(writer)
		encoder.SetIndent("", "\t")
		err := encoder.Encode("")
		if err != nil {
			return
		}
		return
	}
}

func (api OrderApi) handleCreateOrderMapping() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		order := models.NewOrder()
		err := readOrder(request, order)
		if err != nil {
			log.Fatal("Error creating new order for customer!")
			return
		}

		err = defaultEncoder(writer).Encode(order)
		if err != nil {
			log.Fatal("Error encoding created order!")
			return
		}
	}
}
