package handler

import (
	pb "api_gateway/genproto/challenges"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateChallenge(c *gin.Context) {
	req := pb.Challenge{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateChallenge error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.ChallengeService.CreateChallenge(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateChallenge request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) AttendChallenge(c *gin.Context) {
	req := pb.Attend{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("AttendChallenge error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.ChallengeService.AttendChallenge(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("AttendChallenge request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateChallengeResult(c *gin.Context) {
	req := pb.ChallengeUpdate{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateChallengeResult error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.ChallengeService.UpdateChallengeResult(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateChallengeResult request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateEcoTips(c *gin.Context) {
	req := pb.EcoTip{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateEcoTips error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.ChallengeService.CreateEcoTips(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateEcoTips request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllEcoTips(c *gin.Context) {
	req := pb.FilterTip{}

	req.Title = c.Query("title")
	
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

	resp, err := h.ChallengeService.GetAllEcoTips(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetAllEcoTips request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
