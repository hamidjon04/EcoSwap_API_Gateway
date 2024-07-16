package handler

import (
	pb "api_gateway/genproto/recycling_center"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRecyclingCenter(c *gin.Context) {
	req := pb.ResCenter{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateRecyclingCenter error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.CenterService.CreateRecyclingCenter(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateRecyclingCenter request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) SearchRecyclingCenter(c *gin.Context) {
	req := pb.FilterField{}

	req.Material = c.Query("material")

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

	resp, err := h.CenterService.SearchRecyclingCenter(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("SearchRecyclingCenter request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) ProductDelivery(c *gin.Context) {
	req := pb.Submission{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("ProductDelivery error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.CenterService.ProductDelivery(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("ProductDelivery request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetStatistics(c *gin.Context) {
	req := pb.StatisticField{}

	req.StartDate = c.Query("start_date")
	req.EndDate = c.Query("end_date")

	resp, err := h.CenterService.GetStatistics(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetStatistics request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
