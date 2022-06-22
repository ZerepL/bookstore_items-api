package app

import (
	"github.com/ZerepL/bookstore_items-api/controllers"
	"net/http"
)

func mapUrl() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
