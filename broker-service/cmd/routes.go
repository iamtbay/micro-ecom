package main

import (
	"fmt"
	"net/http"
	handlersPackage "tyrping/broker-service/cmd/Handlers"

	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	//cors
	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		fmt.Println("origin", origin)
		if origin == "http://localhost:3000" || origin == "http://127.0.0.1:3000" || origin == "http://127.0.0.1:5173" || origin == "http://localhost:8080" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST,PATCH, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	route := r.Group("/api/v1")
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello!"})
	})
	//auth
	authHandlers := handlersPackage.InitAuthHandlers()
	route.GET("/auth/check", cookieRequired(), authHandlers.Check)
	route.POST("/auth/login", authHandlers.Login)
	route.POST("/auth/signup", authHandlers.Signup)
	route.PATCH("/auth/edit", cookieRequired(), authHandlers.Edit)
	route.PATCH("/auth/change-password", cookieRequired(), authHandlers.ChangePassword)
	route.DELETE("/auth/delete", cookieRequired(), authHandlers.Delete)
	route.POST("/auth/logout", cookieRequired(), authHandlers.Logout)

	//products
	productHandlers := handlersPackage.InitProductHandlers()
	route.GET("/products", productHandlers.GetAllProducts)
	route.GET("/product/:id", productHandlers.GetProductByID)
	route.POST("/product/add", cookieRequired(), productHandlers.AddProduct)
	route.PATCH("/product/:id", cookieRequired(), productHandlers.EditProduct)
	route.DELETE("/product/:id", cookieRequired(), productHandlers.DeleteProduct)

	//cart
	cartHandlers := handlersPackage.InitCartHandlers()
	route.GET("/cart", cookieRequired(), cartHandlers.GetCart)
	route.POST("/cart/checkout", cookieRequired(), cartHandlers.CheckOut)
	route.POST("/cart/new/:id", cookieRequired(), cartHandlers.AddToCart)
	route.PATCH("/cart/:id", cookieRequired(), cartHandlers.UpdateQuantityOfProduct)
	route.DELETE("/cart/:id", cookieRequired(), cartHandlers.DeleteProductOnCart)

	//orders
	orderHandlers := handlersPackage.InitOrderHandlers()
	route.GET("/orders/:id", cookieRequired(), orderHandlers.GetOrdersByUserID)
	route.GET("/order/:id", cookieRequired(), orderHandlers.GetSingleOrder)
	route.POST("/orders/newOrder", cookieRequired(), orderHandlers.NewOrder)
	route.DELETE("/order/:id", cookieRequired(), orderHandlers.DeleteOrder)

	//addresses
	addressHandlers := handlersPackage.InitAddressHandlers()
	route.GET("/address", cookieRequired(), addressHandlers.GetAddresses)
	route.POST("/address", cookieRequired(), addressHandlers.AddNewAddress)
	route.GET("/address/:id", cookieRequired(), addressHandlers.GetSingleAddressByID)
	route.PATCH("/address/:id", cookieRequired(), addressHandlers.EditAddressByID)
	route.DELETE("/address/:id", cookieRequired(), addressHandlers.DeleteAddressByID)

	//reviews
	reviewHandler := handlersPackage.InitReviewHandlers()
	//id=prod id
	route.GET("/reviews/:id", reviewHandler.GetProductReviewsByProductID)
	route.POST("/reviews/:id", cookieRequired(), reviewHandler.NewReview)

	route.GET("/review/:id", cookieRequired(), reviewHandler.GetReviewByID)
	route.PATCH("/review/:id", cookieRequired(), reviewHandler.EditReviewByReviewID)
	route.DELETE("/review/:id", cookieRequired(), reviewHandler.DeleteReviewByReviewID)

	//inventory
	inventoryHandler := handlersPackage.InitInventoryHandlers()
	inventoryRoute := route.Group("/inventory")
	inventoryRoute.Use(cookieRequired())
	inventoryRoute.GET("/:id", inventoryHandler.GetProductStock)
	inventoryRoute.POST("/new", inventoryHandler.AddNewProductStock)
	inventoryRoute.PATCH("/restock/:id", inventoryHandler.RestockProduct)
	inventoryRoute.PATCH("/cancel/:id", inventoryHandler.CancelStockReservation)
	inventoryRoute.PATCH("/reserved/:id", inventoryHandler.ConfirmStockReservation)
	inventoryRoute.PATCH("/sold/:id", inventoryHandler.UpdateStockAfterSale)
}
