package main

import (
	"golang-restaurent-management/database"
	routes"golang-restaurent-management/routes"
	middleware "golang-restaurent-management/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main(){
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":", port)
}



