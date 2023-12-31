package router

import (
	"github.com/labstack/echo/v4"
	handlerInterfaces "github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
	"github.com/shion0625/FYP/backend/pkg/api/middleware"
)

func UserRoutes(api *echo.Group, middleware middleware.Middleware, authHandler handlerInterfaces.AuthHandler,
	userHandler handlerInterfaces.UserHandler,
	productHandler handlerInterfaces.ProductHandler,
	orderHandler handlerInterfaces.OrderHandler,
) {
	// product
	product := api.Group("/products")
	{
		product.GET("/", productHandler.GetAllProductsUser())
		product.GET("/:product_id", productHandler.GetProduct)

		productItem := product.Group("/:product_id/items")
		{
			productItem.GET("/", productHandler.GetAllProductItemsUser())
		}
	}

	// category
	category := api.Group("/categories")
	{
		category.GET("/", productHandler.GetAllCategories)
	}

	auth := api.Group("/auth")
	{
		signup := auth.Group("/sign-in")
		{
			signup.POST("/", authHandler.UserSignUp)
			signup.GET("/", productHandler.GetAllCategories)
		}

		login := auth.Group("/login")
		{
			login.POST("/", authHandler.UserLogin)
			// login.POST("/otp/send", authHandler.UserLoginOtpSend)
			// login.POST("/otp/verify", authHandler.UserLoginOtpVerify)
		}
		auth.POST("/renew-access-token", authHandler.UserRenewAccessToken())
	}

	api.Use(middleware.AuthenticateUser)
	{
		// profile
		account := api.Group("/account")
		{
			account.GET("/", userHandler.GetProfile)
			account.PUT("/", userHandler.UpdateProfile)

			account.GET("/addresses", userHandler.GetAllAddresses)
			account.GET("/address/:address_id", userHandler.GetAddressById)
			account.POST("/address", userHandler.SaveAddress)  // to add a new address
			account.PUT("/address", userHandler.UpdateAddress) // to edit address
			account.POST("/payment-method",
				userHandler.SavePaymentMethod)
			account.GET("/payment-method",
				userHandler.GetAllPaymentMethods)
			account.PUT("/payment-method",
				userHandler.UpdatePaymentMethods)
		}

		// order
		order := api.Group("/order")
		{
			order.GET("/", orderHandler.GetOrderHistory)
			order.POST("/purchase", orderHandler.PayOrder)
		}
	}
}
