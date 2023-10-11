package controllers

import (
	"context"
	"net/http"
	"time"

	"gihtub.com/SherzodAbdullajonov/ecommerce-yt/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id := ctx.Query("id")
		if user_id == "" {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Invalid code"})
			ctx.Abort()
			return
		}
		address, err := primitive.ObjectIDFromHex(user_id)

		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, "Internal Server error")

		}
		var addressess models.Address

		addressess.Address_id = primitive.NewObjectID()

		if err = ctx.BindJSON(&addressess); err != nil {
			ctx.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		match_filter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$address"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$address_id"}, {Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}
		pointcursor, err := UserCollection.Aggregate(c, mongo.Pipeline{match_filter, unwind, group})
		if err != nil {
			ctx.IndentedJSON(500, "Internal sever error")
		}

		var addressinfo []bson.M
		err = pointcursor.All(c, &addressinfo)
		if err != nil {
			panic(err)
		}
		var size int32

		for _, address_no := range addressinfo {
			count := address_no["count"]
			size = count.(int32)
		}
		if size < 2 {
			filter := bson.D{primitive.E{Key: "_id", Value: address}}
			update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: addressess}}}}
			_, err := UserCollection.UpdateOne(c, filter, update)
			if err != nil {
				ctx.IndentedJSON(500, "Internal Server Error")
			}
		} else {
			ctx.IndentedJSON(400, "Not Allowed")
		}
		defer cancel()
		c.Done()
	}
}
func EditHomeAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")

		if id == "" {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Invalid Search Index"})
			ctx.AbortWithStatus(404)
			return
		}

		user_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, "Internal sever error")
		}
		var editaddress models.Address
		if err := ctx.BindJSON(&editaddress); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		}
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address.0.house_name", Value: editaddress.House}, {Key: "address.0.street_name", Value: editaddress.Street}, {Key: "addrss.0.city_name", Value: editaddress.City}, {Key: "address.0.pin_code", Value: editaddress.Pincode}}}}
		_, err = UserCollection.UpdateOne(c, filter, update)
		if err != nil {
			ctx.IndentedJSON(500, "Internal Server error")
			return
		}
		defer cancel()
		c.Done()
		ctx.IndentedJSON(200, "successfully updated the home address")
	}
}
func EditWorkAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")

		if id == "" {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Invalid Search Index"})
			ctx.AbortWithStatus(404)
			return
		}

		user_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, "Internal sever error")
		}
		var editaddress models.Address
		if err := ctx.BindJSON(&editaddress); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		}
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address.0.house_name", Value: editaddress.House}, {Key: "address.0.street_name", Value: editaddress.Street}, {Key: "addrss.0.city_name", Value: editaddress.City}, {Key: "address.0.pin_code", Value: editaddress.Pincode}}}}

		_, err = UserCollection.UpdateOne(c, filter, update)
		if err != nil {
			ctx.IndentedJSON(500, "Internal Server error")
			return
		}
		defer cancel()
		c.Done()
		ctx.IndentedJSON(200, "successfully updated the home address")

	}
}

func DeleteAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")

		if id == "" {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Invalid Search Index"})
			ctx.AbortWithStatus(404)
			return
		}

		addresses := make([]models.Address, 0)
		user_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, "Internal sever error")
		}

		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
		_, err = UserCollection.UpdateOne(c, filter, update)

		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, "Wrong command")
			return
		}

		defer cancel()
		c.Done()
		ctx.IndentedJSON(http.StatusOK, "Successfully Deleted")
	}
}
