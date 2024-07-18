package handler

import (
	pb "api_gateway/genproto/rating"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Foydalanuvchi reytingini yaratish
// @Description  Ushbu endpoint foydalanuvchi reytingini yaratish uchun ishlatiladi.
// @Tags         rating
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        rating          body     rating.RatingReq  true  "Rating request object"
// @Success      200            {object}  rating.RatingResp  "Successful response"
// @Failure      400            {object}  model.Error    "Bad request"
// @Router       /rating/createRating      [post]
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

// @Summary      Foydalanuvchi faoliyatini olish
// @Description  Ushbu endpoint foydalanuvchi faoliyatini olish uchun ishlatiladi.
// @Tags         rating
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id        path     string            true  "User ID"
// @Param        start_date     query    string            true "Start date (YYYY-MM-DD)"
// @Param        end_date       query    string            true "End date (YYYY-MM-DD)"
// @Success      200            {object} rating.Activity  "Successful response"
// @Failure      400            {object} model.Error      "Bad request"
// @Router       /rating/userActivity/{id} [get]
func(h *Handler) GetUserActivity(c *gin.Context){
	req := pb.FilterActivity{UserId: c.Param("id")}

	req.StartDate = c.Query("start_date")
	req.EndDate = c.Query("end_date")

	resp, err := h.RatingService.GetUserActivity(c, &req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetUserActivity request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return 
	}
	
	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchi reytingini olish
// @Description  Ushbu endpoint foydalanuvchi reytingini olish uchun ishlatiladi.
// @Tags         rating
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        user_id        path     string            true  "User ID"
// @Param        limit          query     string           false  "Limit"
// @Param        offset         query     string           false  "Offset"
// @Success      200            {object} rating.UserRating  "Successful response"
// @Failure      400            {object} model.Error        "Bad request"
// @Router       /rating/getUserRating/{user_id} [get]
func(h *Handler) GetUserRating(c *gin.Context){
	req := pb.FilterField{UserId: c.Param("user_id")}

	limit, err := strconv.Atoi(c.Param("limit"))
	if err == nil{
		req.Limit = int32(limit)
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err == nil{
		req.Offset = int32(offset)
	}else{
		req.Offset = 0
	} 


	resp, err := h.RatingService.GetUserRating(c, &req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetUserRating request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}