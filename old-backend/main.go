package main

import (
	"log"
	"os"

	"gihthub.com/shion0625/FYP/backend/controllers"
	"gihthub.com/shion0625/FYP/backend/database"
	"gihthub.com/shion0625/FYP/backend/middleware"
	"gihthub.com/shion0625/FYP/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	r := gin.New()
	r.Use(gin.Logger())
	router := r.Group("/v1")
	router.GET("/categories", app.GetCategory())

	authentication := r.Group("/v1")
	authentication.Use(middleware.Authentication())
	authentication.GET("/addtocart", app.AddToCart())
	authentication.GET("/removeitem", app.RemoveItem())
	authentication.GET("/cartcheckout", app.BuyFromCart())
	authentication.GET("/instantbuy", app.InstantBuy())

	routes.UserRoutes(r)

	log.Fatal(r.Run(":8000"))
}
