package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"gihthub.com/shion0625/FYP/backend/database"
	"gihthub.com/shion0625/FYP/backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productQueryID := ctx.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		userQueryID := ctx.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var c, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.AddProductToCart(c, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err)
		}
		ctx.IndentedJSON(200, "Succesfully added to the cart")
	}
}
func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productQueryID := ctx.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		userQueryID := ctx.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var c, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()
		err = database.RemoveCartItem(c, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		ctx.IndentedJSON(http.StatusOK, "Successfully removed item from cart")
	}
}
func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productQueryID := ctx.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		userQueryID := ctx.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var c, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()
		err = database.InstantBuyer(c, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		ctx.IndentedJSON(http.StatusOK, "successfully placed order")
	}
}
func GetItemFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id := ctx.Query("id")
		if user_id == "" {
			ctx.Header("Content-Type", "applictaion/json")
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			ctx.Abort()
			return
		}
		usert_id, _ := primitive.ObjectIDFromHex(user_id)

		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		var filledcart models.User
		err := UserCollection.FindOne(c, bson.D{primitive.E{Key: "_id", Value: user_id}}).Decode(&filledcart)

		if err != nil {
			log.Println(err)
			ctx.IndentedJSON(http.StatusInternalServerError, "not found")
		}
		filter_match := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: usert_id}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
		grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "&usercart.price"}}}}}}
		pointcursor, err := UserCollection.Aggregate(c, mongo.Pipeline{filter_match, unwind, grouping})
		if err != nil {
			log.Println(err)
		}
		var listing []bson.M
		if err = pointcursor.All(c, &listing); err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)

		}
		for _, json := range listing {
			ctx.IndentedJSON(http.StatusOK, json["total"])
			ctx.IndentedJSON(http.StatusOK, filledcart.UserCart)
		}
		c.Done()

	}

}
func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userQueryID := ctx.Query("id")

		if userQueryID == "" {
			log.Panicln("user id is empty")
			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("UserID is empty"))
		}

		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		err := database.BuyItemFromCart(c, app.userCollection, userQueryID)

		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		ctx.IndentedJSON(http.StatusOK, "successfully placed the order")
	}
}

type Billboard struct {
	id string
	label string
	imageUrl string
}

type Category struct {
	id string
	name string
	billboard Billboard
}
var Categoryes []Category

func (app *Application) GetCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ビルボードデータの作成
		billboard1 := Billboard{
			id: "1",
			label: "Billboard Hot 100",
			imageUrl: "https://example.com/billboard/hot-100.png",
		}
		billboard2 := Billboard{
			id: "2",
			label: "Billboard 200",
			imageUrl: "https://example.com/billboard/200.png",
		}

		// カテゴリデータの作成
		category1 := Category{
			id: "1",
			name: "音楽",
			billboard: billboard1,
		}
		category2 := Category{
			id: "2",
			name: "映画",
			billboard: billboard2,
		}

		// カテゴリデータを配列に追加
		Categoryes = append(Categoryes, category1)
		Categoryes = append(Categoryes, category2)

		ctx.IndentedJSON(200, Categoryes)
	}
}
