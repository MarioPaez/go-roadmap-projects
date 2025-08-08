package main

import (
	"fmt"
	"net/http"
	"unit-tracker-gin/model"
	"unit-tracker-gin/service"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
}

func main() {
	router := gin.Default()
	router.POST("", conversionHandler)
	router.Run()
}

func conversionHandler(g *gin.Context) {
	conversion := model.Conversion{}
	if err := g.ShouldBindBodyWithJSON(&conversion); err != nil {
		g.JSON(http.StatusBadRequest, Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}

	result, err := service.DoConversion(&conversion)

	if err != nil {
		g.JSON(http.StatusBadRequest, Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}

	msg := fmt.Sprintf("Result of your calculation: %.3f %s = %.3f %s", conversion.Value, conversion.From, result, conversion.To)
	g.JSON(200, Response{
		Success: true,
		Message: msg,
	})

}
