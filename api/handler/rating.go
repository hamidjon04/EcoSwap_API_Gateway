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
// @Tags         ratings
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        Authorization  header   string        true  "Bearer token"
// @Param        rating          body     rating.RatingReq  true  "Rating request object"
// @Success      200            {object}  rating.RatingResp  "Successful response"
// @Failure      400            {object}  model.Error    "Bad request"
// @Router       /ratings/createRating      [post]
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
// @Tags         activity
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        Authorization  header   string            true  "Bearer token"
// @Param        user_id        path     string            true  "User ID"
// @Param        start_date     query    string            false "Start date (YYYY-MM-DD)"
// @Param        end_date       query    string            false "End date (YYYY-MM-DD)"
// @Success      200            {object} rating.UserActivity  "Successful response"
// @Failure      400            {object} model.Error      "Bad request"
// @Router       /rating/userActivity/{id} [get]
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

// @Summary      Foydalanuvchi reytingini olish
// @Description  Ushbu endpoint foydalanuvchi reytingini olish uchun ishlatiladi.
// @Tags         rating
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        Authorization  header   string            true  "Bearer token"
// @Param        user_id        path     string            true  "User ID"
// @Param        limit          path     integer           true  "Limit"
// @Param        offset         path     integer           true  "Offset"
// @Success      200            {object} rating.RatingResponse  "Successful response"
// @Failure      400            {object} model.Error        "Bad request"
// @Router       /rating/getUserRating/{user_id}/{limit}/{offset} [get]
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