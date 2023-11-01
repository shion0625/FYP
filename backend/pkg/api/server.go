package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/shion0625/backend/cmd/api/docs"
	handlerInterface "github.com/shion0625/backend/pkg/api/handler/interfaces"
	"github.com/shion0625/backend/pkg/api/middleware"
	"github.com/shion0625/backend/pkg/api/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	Engine *gin.Engine
}

// @title						E-commerce Application Backend API
// @description				Backend API built with Golang using Clean Code architecture. \nGithub: [https://github.com/shion0625/backend].
//
// @contact.name				For API Support
// @contact.email				nikhilnarayanan623@gmail.com
//
// @license.name				MIT
// @license.url				https://opensource.org/licenses/MIT
//
// @BasePath					/api
// @SecurityDefinitions.apikey	BearerAuth
// @Name						Authorization
// @In							header
// @Description				Add prefix of Bearer before  token Ex: "Bearer token"
// @Query.collection.format	multi
func NewServerHTTP(authHandler handlerInterface.AuthHandler, middleware middleware.Middleware,
	adminHandler handlerInterface.AdminHandler, userHandler handlerInterface.UserHandler,
	cartHandler handlerInterface.CartHandler, paymentHandler handlerInterface.PaymentHandler,
	productHandler handlerInterface.ProductHandler, orderHandler handlerInterface.OrderHandler,
	couponHandler handlerInterface.CouponHandler, offerHandler handlerInterface.OfferHandler,
	stockHandler handlerInterface.StockHandler, branHandler handlerInterface.BrandHandler,
) *ServerHTTP {

	engine := gin.New()

	engine.LoadHTMLGlob("views/*.html")

	engine.Use(gin.Logger())

	// swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// set up routes
	routes.UserRoutes(engine.Group("/api"), authHandler, middleware, userHandler, cartHandler,
		productHandler, paymentHandler, orderHandler, couponHandler)
	routes.AdminRoutes(engine.Group("/api/admin"), authHandler, middleware, adminHandler,
		productHandler, paymentHandler, orderHandler, couponHandler, offerHandler, stockHandler, branHandler)

	// no handler
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "invalid url go to /swagger/index.html for api documentation",
		})
	})

	return &ServerHTTP{Engine: engine}
}

func (s *ServerHTTP) Start() error {

	return s.Engine.Run(":8000")
}
