package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	generate "gihtub.com/SherzodAbdullajonov/ecommerce-yt/tokens"
	"gihtub.com/SherzodAbdullajonov/ecommerce-yt/database"
	"gihtub.com/SherzodAbdullajonov/ecommerce-yt/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""

	if err != nil {
		msg = "Login or Password is incorrect"
		valid = false
	}

	return valid, msg
}
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var founduser models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password is incorrect"})
			return
		}

		PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)

		defer cancel()

		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			fmt.Println(msg)
			return
		}
		token, refresToken, _ := generate.TokenGenrator(*founduser.Email, *founduser.First_name, *founduser.Last_name, founduser.User_ID)
		defer cancel()

		generate.UpdateAllTokens(token, refresToken, founduser.User_ID)
		c.JSON(http.StatusNotFound, founduser)
	}

}
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validateErr := Validate.Struct(user)
		if validateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": validateErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"eror": "user already exist"})
			return
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"phone": "this phone number is already used"})
			return
		}
		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshtoken, _ := generate.TokenGenrator(*user.Email, *user.First_name, *user.Last_name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_, insertErr := UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the user did not get created"})
		}
		defer cancel()
		c.JSON(http.StatusCreated, "Successfully signed in!")
	}
}
func ProductViewerAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var products models.Product
		defer cancel()
		if err:= ctx.BindJSON(&products); err!= nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		products.Product_ID = primitive.NewObjectID()
		_, anyerr := ProductCollection.InsertOne(c, products)
		if anyerr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": anyerr.Error()})
			return
		}
		defer cancel()
		ctx.JSON(http.StatusOK, "successfully added")
	}
}

func SearchProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var productlist []models.Product
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := ProductCollection.Find(c, bson.D{{}})
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, "something went wrong please try another time")
			return
		}
		err = cursor.All(c, &productlist)

		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		defer cursor.Close(c)

		if err := cursor.Err(); err != nil {
			log.Println(err)
			ctx.IndentedJSON(400, "invalid")
			return
		}
		defer cancel()
		ctx.JSON(http.StatusOK, productlist)

	}
}
func SearchProductByOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var searchProducts []models.Product
		queryParam := ctx.Query("name")

		//you want to check if it's empty

		if queryParam == "" {
			log.Println("query is empty")
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search index"})
			ctx.Abort()
			return
		}

		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		searchquerydb, err := ProductCollection.Find(c, bson.M{"product_name": bson.M{"$regex": queryParam}})

		if err!= nil {
			ctx.IndentedJSON(404, "something went wrong fetching the data")
			return
		}

		err = searchquerydb.All(c, &searchProducts)
		if err!= nil {
			log.Println(err)
			ctx.IndentedJSON(400, "invalid request")
			return
		}
		defer searchquerydb.Close(c)

		if err:= searchquerydb.Err(); err != nil {
			log.Println(err)
			ctx.IndentedJSON(400, "invalid request")
			return
		}

		defer cancel()
		ctx.IndentedJSON(200, searchProducts)

		
	}
}
