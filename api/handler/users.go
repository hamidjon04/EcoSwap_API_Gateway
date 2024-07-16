package handler

import (
	pb "api_gateway/genproto/users"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfile(c *gin.Context) {
	resp, err := h.UsersService.GetProfile(c, &pb.UserId{Id: c.Param("id")})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetProfile request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

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

func (h *Handler) DeleteProfile(c *gin.Context) {
	resp, err := h.UsersService.DeleteProfile(c, &pb.UserId{Id: c.Param("id")})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteProfile request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	req := pb.FilterField{}

	req.FullName = c.Query("full_name")

	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Limit kiritilmadi yoki xato kiritildi: %v", err))
	}
	offset, err := strconv.Atoi(c.Param("offset"))
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
