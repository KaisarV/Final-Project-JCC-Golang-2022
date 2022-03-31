package routes

import (
	controller "Final-Project-JCC-Golang-2022/controller"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//1. User Biasa 2. Memiliki Toko 3. Admin
	router.POST("/login", controller.UserLogin)
	router.GET("/logout", controller.Logout)
	router.POST("/register", controller.InsertUser)

	admin := router.Group("/")
	admin.Use(controller.Authenticate(3))
	{
		admin.DELETE("/users/:id", controller.DeleteUser)
	}

	basicUser := router.Group("/")
	basicUser.Use(controller.Authenticate(1))
	{
		basicUser.PUT("/user", controller.UpdateUsers)
		basicUser.POST("/store", controller.InsertMyStore)

		//Cart
		basicUser.GET("/carts", controller.GetAllMyCart)
		basicUser.DELETE("/cart/:cartId", controller.DeleteMyCart)
		basicUser.POST("/cart", controller.InsertMyCart)
		basicUser.PUT("/cart/:cartId", controller.UpdateMyCart)

		//Transaction
		basicUser.GET("/transactions", controller.GetAllMyTransaction)
		basicUser.POST("/transaction", controller.InsertMyTransactions)

		//Product Review
		basicUser.GET("/reviews", controller.GetAllMyProductReviews)
		basicUser.DELETE("/review/:productid", controller.DeleteMyProductReview)
		basicUser.POST("/review/:productid", controller.InsertMyProductReview)
	}

	storeOwner := router.Group("/")
	storeOwner.Use(controller.Authenticate(2))
	{
		//Store
		storeOwner.DELETE("/store", controller.DeleteMyStore)
		storeOwner.PUT("/store", controller.UpdateMyStore)

		//Products
		storeOwner.DELETE("/product/:productId", controller.DeleteMyProduct)
		storeOwner.POST("/product", controller.InsertMyProduct)
		storeOwner.PUT("/product/:productId", controller.UpdateMyProduct)
	}

	router.GET("/users", controller.GetAllUsers)
	router.GET("/stores", controller.GetAllStores)
	router.GET("/products", controller.GetAllProducts)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
