package handler

import (
	pb "api_gateway/genproto/recycling_center"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Yangi qayta ishlash markazi yaratish
// @Description  Ushbu endpoint yangi qayta ishlash markazini yaratish uchun ishlatiladi.
// @Tags         recycling_center
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization   header    string       true  "Bearer token"
// @Param        recycling_center  body      recycling_center.ResCenter  true  "Recycling Center Payload"
// @Success      200               {object}  recycling_center.ResCenter  "Successful Response"
// @Failure      400               {object}  model.Error   "Bad Request"
// @Router       /center/create [post]
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

// @Summary      Qayta ishlash markazini qidirish
// @Description  Ushbu endpoint qayta ishlash markazini qidirish uchun ishlatiladi.
// @Tags         recycling_center
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization   header    string       true  "Bearer token"
// @Param        limit   query      integer     true        "Limit"
// @Param        offset  query      integer     true        "Offset"
// @Param        material query     string      false       "Material type"
// @Success      200               {object}  recycling_center.ResCenter  "Successful Response"
// @Failure      400               {object}  model.Error   "Bad Request"
// @Router       /center/search/{limit}/{offset} [get]
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

// @Summary      Maxsulot yetkazib berish
// @Description  Ushbu endpoint maxsulot yetkazib berish uchun ishlatiladi.
// @Tags         delivery
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization   header    string       true  "Bearer token"
// @Param        delivery_body  body      recycling_center.Submission  true  "Delivery Payload"
// @Success      200               {object}  recycling_center.Status  "Successful Response"
// @Failure      400               {object}  model.Error   "Bad Request"
// @Router       /center/productDelivery [post]
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

// @Summary      Statistika olish
// @Description  Ushbu endpoint statistika ma'lumotlarini olish uchun ishlatiladi.
// @Tags         statistics
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Authorization   header    string       true  "Bearer token"
// @Param        start_date      query     string       true  "Start date in format yyyy-mm-dd"
// @Param        end_date        query     string       true  "End date in format yyyy-mm-dd"
// @Success      200             {object}  recycling_center.StatisticField  "Successful Response"
// @Failure      400             {object}  model.Error       "Bad Request"
// @Router       /center/statistics     [get]
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
