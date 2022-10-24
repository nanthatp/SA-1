package main

import (
	"github.com/Nanthat/sa-65-example/controller"
	"github.com/Nanthat/sa-65-example/entity"
	"github.com/Nanthat/sa-65-example/middlewares"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{

		router.Use(middlewares.Authorizes())
		{
			// Employee Routes
			router.GET("/employees", controller.ListEmployees)
			//router.GET("/employee/:id", controller.ListEmployees)
			router.GET("/employee/:id", controller.GetEmployee)
			//router.POST("/employees", controller.CreateEmployee)
			router.PATCH("/employees", controller.UpdateEmployee)
			router.DELETE("/employees/:id", controller.DeleteEmployee)

			// // gender Routes
			// router.GET("/genders", controller.ListGenders)
			// router.GET("/gender/:id", controller.GetGender)
			// router.POST("/genders", controller.CreateGender)
			// router.PATCH("/genders", controller.UpdateGender)
			// router.DELETE("/genders/:id", controller.DeleteGender)

			// // provinces Routes
			// router.GET("/provinces", controller.ListProvinces)
			// router.GET("/province/:id", controller.GetProvince)
			// router.POST("/provinces", controller.CreateProvince)
			// router.PATCH("/provinces", controller.UpdateProvince)
			// router.DELETE("/provinces/:id", controller.DeleteProvince)

			// Member Routes
			router.GET("/members", controller.ListMembers)
			router.GET("/member/:id", controller.GetMember)
			router.POST("/members", controller.CreateMember)
			router.PATCH("/members", controller.UpdateMember)
			router.DELETE("/members/:id", controller.DeleteMember)

			// // Typeproduct Routes
			// router.GET("/typeproducts", controller.ListTypeproducts)
			// router.GET("/typeproduct/:id", controller.GetTypeproduct)
			// router.POST("/typeproducts", controller.CreateTypeproduct)
			// router.PATCH("/typeproducts", controller.UpdateTypeproduct)
			// router.DELETE("/typeproducts/:id", controller.DeleteTypeproduct)

			// // Manufacturer Routes
			// router.GET("/manufacturers", controller.ListManufacturers)
			// router.GET("/manufacturer/:id", controller.GetManufacturer)
			// router.POST("/manufacturers", controller.CreateManufacturer)
			// router.PATCH("/manufacturers", controller.UpdateManufacturer)
			// router.DELETE("/manufacturers/:id", controller.DeleteManufacturer)

			// Product Routes
			router.GET("/products", controller.ListProducts)
			router.GET("/product/:id", controller.GetProduct)
			router.POST("products", controller.CreateProduct)
			router.PATCH("/products", controller.UpdateProduct)
			router.DELETE("/products/:id", controller.DeleteProduct)

			// // Shelf Routes
			// router.GET("/shelfproducts", controller.ListShelfproducts)
			// router.GET("/shelfproduct/:id", controller.GetShelfproduct)
			// router.POST("/shelfproducts", controller.CreateShelfproduct)
			// router.PATCH("/shelfproducts", controller.UpdateShelfproduct)
			// router.DELETE("/shelfproducts/:id", controller.DeleteShelfproduct)

			// // Lot Routes
			// router.GET("/lots", controller.ListLots)
			// router.GET("/lot/:id", controller.GetLot)
			// router.POST("/lots", controller.CreateLot)
			// router.PATCH("/lots", controller.UpdateLot)
			// router.DELETE("/lots/:id", controller.DeleteLot)

			// // Stock Routes
			// router.GET("/stocks", controller.ListStocks)
			// router.GET("/stock/:id", controller.GetStock)
			// router.POST("/stocks", controller.CreateStock)
			// router.PATCH("/stocks", controller.UpdateStock)
			// router.DELETE("/stocks/:id", controller.DeleteStock)

			// Cart Routes
			router.GET("/carts", controller.ListCarts)
			router.GET("/cart/:id", controller.GetCart)
			router.POST("carts", controller.CreateCart)
			router.PATCH("/carts", controller.UpdateCart)
			router.DELETE("/carts/:id", controller.DeleteCart)

			// Order Routes
			router.GET("/orders", controller.ListOrders)
			router.GET("/order/:id", controller.GetOrder)
			router.POST("orders", controller.CreateOrder)
			router.PATCH("/orders", controller.UpdateOrder)
			router.DELETE("/orders/:id", controller.DeleteOrder)

			// // Paymenttype Routes
			// router.GET("/paymenttypes", controller.ListPaymenttypes)
			// router.GET("/paymenttype/:id", controller.GetPaymenttype)
			// router.POST("/paymenttypes", controller.CreatePaymenttype)
			// router.PATCH("/paymenttypes", controller.UpdatePaymenttype)
			// router.DELETE("/paymenttypes/:id", controller.DeletePaymenttype)

			// // Receipt Routes
			// router.GET("/receipts", controller.ListReceipts)
			// router.GET("/receipt/:id", controller.GetReceipt)
			// router.POST("/receipts", controller.CreateReceipt)
			// router.PATCH("/receipts", controller.UpdateReceipt)
			// router.DELETE("/receipts/:id", controller.DeleteReceipt)

		}
	}
	// Signup User Route
	r.POST("/signup", controller.CreateEmployee)
	// login User Route
	r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
