package handler

import (
	pb "api_gateway/genproto/challenges"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Yangi challenge yaratish
// @Description  Ushbu endpoint yangi challenge yaratish uchun ishlatiladi.
// @Tags         challenge
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        challenge      body     challenges.Challenge  true  "Challenge payload"
// @Success      200            {object} challenges.RespChallenge  "Successful response"
// @Failure      400            {object} model.Error   "Bad request"
// @Router       /challenge/createChallenge [post]
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

// @Summary      challengening ishtirokchisiga bo'lish
// @Description  Ushbu endpoint challengening ishtirokchisiga bo'lish uchun ishlatiladi.
// @Tags         challenge
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        attend          body     challenges.Attend  true  "Attend payload"
// @Success      200            {object} challenges.AttendResp  "Successful response"
// @Failure      400            {object} model.Error   "Bad request"
// @Router       /challenge/attend   [post]
func (h *Handler) AttendChallenge(c *gin.Context) {
	req := pb.Attend{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("AttendChallenge error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(req.ChallengeId)

	resp, err := h.ChallengeService.AttendChallenge(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("AttendChallenge request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// @Summary      challenge natijasini yangilash
// @Description  Ushbu endpoint challenge natijasini yangilash uchun ishlatiladi.
// @Tags         challenge
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        update          body     challenges.ChallengeUpdate  true  "Update payload"
// @Success      200            {object} challenges.RespUpdate  "Successful response"
// @Failure      400            {object} model.Error   "Bad request"
// @Router       /challenge/updateChallenge   [put]
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

// @Summary      Ekologik maslahat yaratish
// @Description  Ushbu endpoint ekologik maslahat yaratish uchun ishlatiladi.
// @Tags         ico_tip
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        create          body     challenges.EcoTip  true  "Create payload"
// @Success      200            {object} challenges.RespEcoTip  "Successful response"
// @Failure      400            {object} model.Error   "Bad request"
// @Router       /challenge/createEcoTip   [post]
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

// @Summary      Ekologik maslahatlar ro'yxati
// @Description  Ushbu endpoint ekologik maslahatlar ro'yxatini olish uchun ishlatiladi.
// @Tags         ico_tip
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        title          query     string   false  "Filter by title"
// @Param        limit          query      string      false   "Limit"
// @Param        offset         query      string      false   "Offset"
// @Success      200            {object}  challenges.Tips  "Successful response"
// @Failure      400            {object}  model.Error    "Bad request"
// @Router       /challenge/getAllEcoTips/   [get]
func (h *Handler) GetAllEcoTips(c *gin.Context) {
	req := pb.FilterTip{}

	req.Title = c.Query("title")

	limit, err := strconv.Atoi(c.Param("limit"))
	if err == nil {
		req.Limit = int32(limit)
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err == nil {
		req.Offset = int32(offset)
	}else{
		req.Offset = 0
	}


	resp, err := h.ChallengeService.GetAllEcoTips(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetAllEcoTips request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
