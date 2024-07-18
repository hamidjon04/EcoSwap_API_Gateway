package handler

import (
	pb "api_gateway/genproto/swaps"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Swap so'rovini jo'natish
// @Description  Ushbu endpoint avval kiritilgan ma'lumotlar bilan swap so'rovini jo'natish uchun ishlatiladi.
// @Tags         swap
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body           body     swaps.SwapRequest    true  "Swap Request Payload"
// @Success      200            {object} swaps.SwapResponce  "Successful response"
// @Failure      400            {object} model.Error      "Bad request"
// @Router       /swap/sendRequest [post]
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

// @Summary      Swap so'rovini qabul qilish
// @Description  Ushbu endpoint asosida swap so'rovi uchun sababni qabul qilish uchun ishlatiladi.
// @Tags         swap
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body           body     swaps.Reason       true  "Reason Payload"
// @Success      200            {object} swaps.Responce  "Successful response"
// @Failure      400            {object} model.Error    "Bad request"
// @Router       /swap/adoptionSwap [post]
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

// @Summary      Swap so'rovini rad qilish
// @Description  Ushbu endpoint asosida swap so'rovini rad qilish uchun sababni qabul qilish uchun ishlatiladi.
// @Tags         swap
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body           body     swaps.Reason       true  "Reason Payload"
// @Success      200            {object} swaps.Responce  "Successful response"
// @Failure      400            {object} model.Error    "Bad request"
// @Router       /swap/rejectRequest [post]
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

// @Summary      Barcha swap so'rovlarni olish
// @Description  Ushbu endpoint asosida barcha swap so'rovlarni olish uchun ishlatiladi.
// @Tags         swap
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        status         query     string            false  "Status of swap requests (e.g., pending, accepted, rejected)"
// @Param        limit          query     string           false  "Limit of items to return"
// @Param        offset         query     string           false  "Offset for pagination"
// @Success      200            {object} swaps.AllSwaps  "Successful response"
// @Failure      400            {object} model.Error      "Bad request"
// @Router       /swap/allRequests/ [get]
func (h *Handler) GetAllSwapRequests(c *gin.Context) {
	req := pb.FilterField{}

	req.Status = c.Query("status")

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


	resp, err := h.SwapService.GetAllSwapRequests(c, &req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("RejectionSwapRequest request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
