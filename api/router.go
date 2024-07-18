package api

import (
	_ "api_gateway/api/docs"
	"api_gateway/api/handler"
	"api_gateway/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth Service API3
// @version 1.0
// @description This is a sample server for Auth Service.
// @host localhost:4444
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
func Router() *gin.Engine {
	router := gin.Default()
	h := handler.NewHandler()

	// Swagger endpointini sozlash
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.JWTMiddleware())

	center := router.Group("/center")
	center.POST("/create", h.CreateRecyclingCenter)
	center.GET("/search", h.SearchRecyclingCenter)
	center.POST("/productDelivery", h.ProductDelivery)
	center.GET("/statistics", h.GetStatistics)

	challenge := router.Group("/challenge")
	challenge.POST("/createChallenge", h.CreateChallenge)
	challenge.POST("/attend", h.AttendChallenge)
	challenge.PUT("/updateChallenge", h.UpdateChallengeResult)
	challenge.POST("/createEcoTip", h.CreateEcoTips)
	challenge.GET("/getAllEcoTips", h.GetAllEcoTips)

	item := router.Group("/item")
	item.POST("/createItem", h.CreateItem)
	item.PUT("/updateItem", h.UpdateItem)
	item.DELETE("/deleteItem/:id", h.DeleteItem)
	item.GET("/getItem/:id", h.GetItem)
	item.GET("/searchItem", h.SearchItems)
	item.GET("/getAllItems", h.GetAllItems)
	item.POST("/createCategory", h.CreateCategory)

	rating := router.Group("/rating")
	rating.POST("/createRating", h.CreateUserRating)
	rating.GET("/userActivity/:id", h.GetUserActivity)
	rating.GET("/getUserRating/:user_id", h.GetUserRating)

	swap := router.Group("/swap")
	swap.POST("/sendRequest", h.SendSwapRequest)
	swap.POST("/adoptionSwap", h.AdoptionSwapRequest)
	swap.POST("/rejectRequest", h.RejectionSwapRequest)
	swap.GET("/allRequests", h.GetAllSwapRequests)

	user := router.Group("/user")
	user.GET("/getProfile/:id", h.GetProfile)
	user.PUT("/updateProfile", h.UpdateProfile)
	user.DELETE("/deleteProfile/:id", h.DeleteProfile)
	user.GET("/getAllUsers", h.GetAllUsers)
	user.GET("/userEcopoints/:id", h.GetEcoPointsByUser)
	user.POST("/createEcopoints", h.CreateEcoPointsByUser)
	user.GET("/historyUserPoint/:id", h.HistoryEcoPointsByUser)

	return router
}
