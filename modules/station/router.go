package station

import (
	"net/http"

	"github.com/HaikalRFadhilahh/api-mrt-go-bwa/common/response"
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	// Create Station Service Struct Impl Interface
	stationService := NewService()

	// Create Sub Router Group
	station := router.Group("/stations")

	// Routing
	station.GET("/", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})

	station.GET("/:id", func(c *gin.Context) {
		CheckShedulesByStation(c, stationService)
	})
}

func GetAllStation(ctx *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		// Handle Error
		ctx.JSON(http.StatusInternalServerError, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// Return Response HTTP
	ctx.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Data MRT Train",
		Data:    datas,
	})
}

func CheckShedulesByStation(ctx *gin.Context, service Service) {
	// Take ID From URL
	id := ctx.Param("id")

	// Execute Service Function
	datas, err := service.CheckSheduleByStation(id)
	if err != nil {
		// Handle Error
		ctx.JSON(http.StatusInternalServerError, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// Return Response HTTP
	ctx.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Data Schedule By Station",
		Data:    datas,
	})
}
