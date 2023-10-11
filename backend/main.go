package main

import (
	"log"
	"os"

	"gihtub.com/SherzodAbdullajonov/ecommerce-yt/controllers"
	"gihtub.com/SherzodAbdullajonov/ecommerce-yt/database"
	"gihtub.com/SherzodAbdullajonov/ecommerce-yt/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":8000"))
}
