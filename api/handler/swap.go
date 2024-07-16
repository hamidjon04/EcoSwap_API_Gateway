package handler

import (
	pb "api_gateway/genproto/swaps"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendSwapRequest(c *gin.Context) {
	req := pb.SwapRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("SendSwapRequest error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.SwapService.SendSwapRequest(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("SendSwapRequest request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) AdoptionSwapRequest(c *gin.Context) {
	req := pb.Reason{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("AdoptionSwapRequest error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.SwapService.AdoptionSwapRequest(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("AdoptionSwapRequest request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) RejectionSwapRequest(c *gin.Context) {
	req := pb.Reason{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("RejectionSwapRequest error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.SwapService.RejectionSwapRequest(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("RejectionSwapRequest request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllSwapRequests(c *gin.Context) {
	req := pb.FilterField{}

	req.Status = c.Param("status")

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

	resp, err := h.SwapService.GetAllSwapRequests(c, &req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("RejectionSwapRequest request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
