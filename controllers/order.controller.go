package controllers

import (
	"github.com/g-s-pai/go-order-service/initializers"
	"github.com/g-s-pai/go-order-service/models"
	"github.com/g-s-pai/go-order-service/pubsub"

	"github.com/kataras/iris/v12"
)

func GetOrders(ctx iris.Context) {
	var orders []models.Order
	if err := initializers.DB.Find(&orders).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.JSON(orders)
}

func GetOrderByID(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var order models.Order
	if err := initializers.DB.First(&order, "id = ?", id).Error; err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Order not found"})
		return
	}
	ctx.JSON(order)
}

func CreateOrder(ctx iris.Context) {
	var order models.Order
	if err := ctx.ReadJSON(&order); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	if err := initializers.DB.Create(&order).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	// Publish to PubSub
	_ = pubsub.PublishOrderEvent(map[string]interface{}{
		"order_id": order.ID,
		"user_id":  order.UserID,
		"amount":   order.Amount,
	})

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(order)
}

func UpdateOrder(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var order models.Order
	var checkOrder models.Order

	if err := initializers.DB.First(&checkOrder, "id = ?", id).Error; err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Order not found"})
		return
	}

	if err := ctx.ReadJSON(&order); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	if err := initializers.DB.Save(&order).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.JSON(order)
}
