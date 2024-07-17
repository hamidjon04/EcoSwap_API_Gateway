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
// @Tags         challenges
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        challenge      body     challenges.Challenge  true  "Challenge payload"
// @Success      200            {object} challenges.Challenge  "Successful response"
// @Failure      400            {object} model.Error   "Bad request"
// @Router       challenge/createChallenge    [post]
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
// @Tags         challenges
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        attend          body     challenges.Attend  true  "Attend payload"
// @Success      200            {object} challenges.Attend  "Successful response"
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

	resp, err := h.ChallengeService.AttendChallenge(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("AttendChallenge request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      challenge natijasini yangilash
// @Description  Ushbu endpoint challenge natijasini yangilash uchun ishlatiladi.
// @Tags         challenges
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        update          body     challenges.ChallengeUpdate  true  "Update payload"
// @Success      200            {object} challenges.ChallengeUpdate  "Successful response"
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
// @Tags         eco_tips
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        create          body     challenges.EcoTip  true  "Create payload"
// @Success      200            {object} challenges.EcoTip  "Successful response"
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
// @Tags         eco_tips
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization  header   string   true  "Bearer token"
// @Param        title          query     string   false  "Filter by title"
// @Param        limit          path      int      true   "Limit"
// @Param        offset         path      int      true   "Offset"
// @Success      200            {object}  challenges.EcoTipList  "Successful response"
// @Failure      400            {object}  model.Error    "Bad request"
// @Router       /challenge/getAllEcoTips/{limit}/{offset}   [get]
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
