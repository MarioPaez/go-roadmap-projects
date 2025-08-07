package main

import (
	"fmt"
	"net/http"
	"unit-tracker-gin/length"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
}

type Conversion struct {
	To    string
	Value float64
	From  string
}

func main() {
	router := gin.Default()
	router.POST("/length", lengthConversionHandler)
	router.Run()

}

func lengthConversionHandler(g *gin.Context) {
	conversion := Conversion{}
	if err := g.ShouldBindBodyWithJSON(&conversion); err != nil {
		g.JSON(http.StatusBadRequest, Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}
	if err := length.ValidateUnits(conversion.To, conversion.From); err != nil {
		g.JSON(http.StatusBadRequest, Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}
	msg := fmt.Sprintf("To %s from: %s y el valor es: %f", conversion.To, conversion.From, conversion.Value)
	g.JSON(200, Response{
		Success: true,
		Message: msg,
	})
}
