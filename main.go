package main

import (
	"log"

	"github.com/g-s-pai/go-order-service/initializers"
	"github.com/g-s-pai/go-order-service/routes"
	"github.com/g-s-pai/go-order-service/pubsub"
	"github.com/g-s-pai/go-order-service/middleware"
	
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

var (
	app *iris.Application
	config initializers.Config
)

func init() {
	_ = godotenv.Load()

	initializers.ConnectDB()

	if err := pubsub.InitPubSub(); err != nil {
		log.Fatal("PubSub init failed:", err)
	}
	
	app = iris.Default()
}

func main() {
	app.Post("/login", func(ctx iris.Context) {
		var req struct {
			UserID string `json:"user_id"`
		}
		if err := ctx.ReadJSON(&req); err != nil {
			ctx.StopWithStatus(400)
			return
		}

		token, err := middleware.GenerateJWT(req.UserID)
		if err != nil {
			ctx.StopWithStatus(500)
			return
		}

		ctx.JSON(iris.Map{
			"token": token,
		})
	})

	router := app.Party("/api/v1", middleware.JWTMiddleware)

	routes.OrderRoutes(router)

	app.Listen(":8081")
}