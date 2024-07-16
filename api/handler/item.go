package handler

import (
	pb "api_gateway/genproto/items"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateItem(c *gin.Context) {
	req := pb.Item{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateItem error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.ItemService.CreateItem(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateItem request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateItem(c *gin.Context) {
	req := pb.ItemUpdate{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateItem error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.ItemService.UpdateItem(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateItem request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteItem(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.ItemService.DeleteItem(c, &pb.ItemId{Id: id})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteItem request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetItem(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.ItemService.GetItem(c, &pb.ItemId{Id: id})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetItem request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) SearchItems(c *gin.Context) {
	req := pb.FilterField{}

	req.CategoryId = c.Query("category_id")
	req.Condition = c.Query("condition")
	req.Name = c.Query("name")
	
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

	resp, err := h.ItemService.SearchItems(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("SearchItems request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllItems(c *gin.Context) {
	req := pb.Limit{}

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

	resp, err := h.ItemService.GetAllItems(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetAllItems request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateCategory(c *gin.Context) {
	req := pb.Category{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateCategory error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.ItemService.CreateCategory(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateCategory request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
