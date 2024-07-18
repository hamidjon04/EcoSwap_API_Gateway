package handler

import (
	pb "api_gateway/genproto/users"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Foydalanuvchi profilini olish
// @Description  Ushbu endpoint asosida foydalanuvchi profilini olish uchun ishlatiladi.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id             path     string            true  "Foydalanuvchi identifikatori"
// @Success      200            {object} users.UserInfo   "Successful response"
// @Failure      400            {object} model.Error      "Bad request"
// @Router       /user/getProfile/{id} [get]
func (h *Handler) GetProfile(c *gin.Context) {
	resp, err := h.UsersService.GetProfile(c, &pb.UserId{Id: c.Param("id")})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetProfile request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchi profilini yangilash
// @Description  Ushbu endpoint foydalanuvchi profilini yangilash uchun ishlatiladi.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body           body     users.ProfileUpdate    true  "Profil yangilash so'rovi"
// @Success      200            {object} users.UpdateResponse          "Yangilash muvaffaqiyatli"
// @Failure      400            {object} model.Error        "Xato so'rov"
// @Router       /user/updateProfile [put]
func (h Handler) UpdateProfile(c *gin.Context) {
	req := pb.ProfileUpdate{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateProfile error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.UsersService.UpdateProfile(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateProfile request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchi profilini o'chirish
// @Description  Ushbu endpoint foydalanuvchi profilini o'chirish uchun ishlatiladi.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id             path     string       true  "Foydalanuvchi ID"
// @Success      200            {object} users.Status   "O'chirish muvaffaqiyatli"
// @Failure      400            {object} model.Error "Xato so'rov"
// @Router       /user/deleteProfile/{id} [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	resp, err := h.UsersService.DeleteProfile(c, &pb.UserId{Id: c.Param("id")})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteProfile request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Barcha foydalanuvchilarni olish
// @Description  Ushbu endpoint barcha foydalanuvchilarni olish uchun ishlatiladi.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        full_name      query    string       false "Foydalanuvchi to'liq ismi"
// @Param        limit          query     string      false  "Limit"
// @Param        offset         query     string      false  "Offset"
// @Success      200            {object} users.Users "Foydalanuvchilar ro'yxati"
// @Failure      400            {object} model.Error  "Xato so'rov"
// @Router       /user/getAllUsers [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	req := pb.FilterField{}

	req.FullName = c.Query("full_name")

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Limit kiritilmadi yoki xato kiritildi: %v", err))
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Offset kiritilmadi yoki xato kiritildi: %v", err))
	}

	req.Limit = int32(limit)
	req.Offset = int32(offset)

	resp, err := h.UsersService.GetAllUsers(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetAllUsers request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchi eko ballarini olish
// @Description  Ushbu endpoint foydalanuvchi eko ballarini olish uchun ishlatiladi.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id              path     string       true  "Foydalanuvchi ID"
// @Success      200            {object} users.UserEcoPoints  "Foydalanuvchi eko ballari"
// @Failure      400            {object} model.Error       "Xato so'rov"
// @Router       /user/userEcopoints/{id} [get]
func(h *Handler) GetEcoPointsByUser(c *gin.Context){
	resp, err := h.UsersService.GetEcoPointsByUser(c, &pb.UserId{Id: c.Param("id")})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetEcoPointsByUser request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchi eko ballarini yaratish
// @Description  Ushbu endpoint foydalanuvchi uchun eko ballarini yaratish uchun ishlatiladi.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body           body     users.CreateEcoPoints true  "Foydalanuvchi eko ballari ma'lumotlari"
// @Success      200            {object} users.InfoUserEcoPoints         "Yaratish muvaffaqiyatli"
// @Failure      400            {object} model.Error       "Xato so'rov"
// @Router       /user/createEcopoints [post]
func(h *Handler) CreateEcoPointsByUser(c *gin.Context){
	req := pb.CreateEcoPoints{}

	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("CreateEcoPointsByUser error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.UsersService.CreateEcoPointsByUser(c, &req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("CreateEcoPointsByUser request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchi eko ballari tarixini olish
// @Description  Ushbu endpoint foydalanuvchi uchun eko ballari tarixini olish uchun ishlatiladi.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id             path     string      true   "Id"
// @Param        limit          query     string      false  "Limit"
// @Param        offset         query     string      false  "Offset"
// @Success      200            {object} users.Histories  "Tarix ro'yxati"
// @Failure      400            {object} model.Error       "Xato so'rov"
// @Router       /user/historyUserPoint/{id} [get]
func(h *Handler) HistoryEcoPointsByUser(c *gin.Context){
	req := pb.HistoryReq{UserId: c.Param("id")}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err == nil {
		req.Limit = int32(limit)
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err == nil {
		req.Offset = int32(offset)
	}else{
		req.Offset = 0
	}
	

	resp, err := h.UsersService.HistoryEcoPointsByUser(c, &req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("HistoryEcoPointsByUser query error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
