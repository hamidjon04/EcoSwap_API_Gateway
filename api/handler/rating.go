package handler

import (
	pb "api_gateway/genproto/rating"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func(h *Handler) CreateUserRating(c *gin.Context){
	req := pb.RatingReq{}

	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("CreateUserRating error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return 
	}

	resp, err := h.RatingService.CreateUserRating(c, &req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("CreateUserRating request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return 
	}

	c.JSON(http.StatusOK, resp)
}


func(h *Handler) GetUserActivity(c *gin.Context){
	req := pb.FilterActivity{UserId: c.Param("user_id")}

	req.StartDate = c.Param("start_date")
	req.EndDate = c.Param("end_date")

	resp, err := h.RatingService.GetUserActivity(c, &req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetUserActivity request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return 
	}
	
	c.JSON(http.StatusOK, resp)
}


func(h *Handler) GetUserRating(c *gin.Context){
	req := pb.FilterField{UserId: c.Param("user_id")}

	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Limit kiritilmadi yoki xato kiritildi: %v", err))
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Offset kiritilmadi yoki xato kiritildi: %v", err))
	} 

	req.Limit = int32(limit)
	req.Offset = int32(offset)

	resp, err := h.RatingService.GetUserRating(c, &req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetUserRating request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}