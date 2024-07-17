package api

import (
	"api_gateway/api/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	h := handler.NewHandler()

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
	item.DELETE("/deleteItem", h.DeleteItem)
	item.GET("/getItem", h.GetItem)
	item.GET("/searchItem", h.SearchItems)
	item.GET("/getAllItems", h.GetAllItems)
	item.POST("/createCategory", h.CreateCategory)

	rating := router.Group("/rating")
	rating.POST("/createRating", h.CreateUserRating)
	rating.GET("/userActivity/:id", h.GetUserActivity)
	rating.GET("/getUserRating", h.GetUserRating)

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
	user.GET("/historyUserRcopoints", h.HistoryEcoPointsByUser)

	return router
}
