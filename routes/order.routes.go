package routes

import (
	"github.com/g-s-pai/go-order-service/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func OrderRoutes(rg router.Party) {
	router := rg.Party("/orders")

	router.Use(iris.Compression)
	router.Get("/", controllers.GetOrders)
	router.Get("/{id}", controllers.GetOrderByID)
	router.Post("/", controllers.CreateOrder)
	// router.Post("/item", controllers.CreateOrderItem)
	router.Patch("/{id}", controllers.UpdateOrder)
}
