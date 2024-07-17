package handler

import (
	pb "api_gateway/genproto/items"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Mahsulot qo'shish
// @Description  Ushbu endpoint mahsulot qo'shish uchun ishlatiladi.
// @Tags         items
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        item            body     items.Item  true  "Item object"
// @Success      200            {object}  items.Item  "Successful response"
// @Failure      400            {object}  model.Error  "Bad request"
// @Router       /item/createItem  [post]
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

// @Summary      Mahsulotni yangilash
// @Description  Ushbu endpoint mahsulotni yangilash uchun ishlatiladi.
// @Tags         items
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        item_update    body     items.ItemUpdate  true  "Item Update object"
// @Success      200            {object}  items.Item  "Successful response"
// @Failure      400            {object}  model.Error  "Bad request"
// @Router       /item/updateItem  [put]
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

// @Summary      Mahsulotni o'chirish
// @Description  Ushbu endpoint mahsulotni o'chirish uchun ishlatiladi.
// @Tags         items
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        id             path     string   true  "Item ID"
// @Success      200            {object}  items.DeleteResponse  "Successful response"
// @Failure      400            {object}  model.Error        "Bad request"
// @Router       /item/deleteItem{id}    [delete]
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

// @Summary      Mahsulotni olish
// @Description  Ushbu endpoint mahsulotni olish uchun ishlatiladi.
// @Tags         items
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        id             path     string   true  "Item ID"
// @Success      200            {object}  items.Item  "Successful response"
// @Failure      400            {object}  model.Error  "Bad request"
// @Router       /item/getItem{id}    [get]
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

// @Summary      Mahsulotlarni qidirish
// @Description  Ushbu endpoint mahsulotlarni qidirish uchun ishlatiladi.
// @Tags         items
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        category_id    query    string   false  "Category ID"
// @Param        condition      query    string   false  "Condition"
// @Param        name           query    string   false  "Name"
// @Param        limit          path     integer  false  "Limit"
// @Param        offset         path     integer  false  "Offset"
// @Success      200            {object}  items.ItemList  "Successful response"
// @Failure      400            {object}  model.Error  "Bad request"
// @Router       /item/searchItem  [get]
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

// @Summary      Barcha mahsulotlarni olish
// @Description  Ushbu endpoint barcha mahsulotlarni olish uchun ishlatiladi.
// @Tags         items
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        limit          path     integer  false  "Limit"
// @Param        offset         path     integer  false  "Offset"
// @Success      200            {object}  items.ItemList  "Successful response"
// @Failure      400            {object}  model.Error  "Bad request"
// @Router       /item/getAllItems     [get]
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

// @Summary      Kategoriya yaratish
// @Description  Ushbu endpoint kategoriya yaratish uchun ishlatiladi.
// @Tags         categories
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        category       body     items.Category  true  "Category object"
// @Success      200            {object}  items.Category  "Successful response"
// @Failure      400            {object}  model.Error  "Bad request"
// @Router       /item/createCategory    [post]
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
