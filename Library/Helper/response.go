package Helper

import (
	"github.com/Revelts/module-errors"
	"github.com/gin-gonic/gin"
	"mobile-banking-service/Constants"
	"net/http"
	"time"
)

type ResponsePublisher struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
}

type setResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Code       int         `json:"code"`
	AccessTime string      `json:"accessTime"`
}

func HttpResponseSuccess(g *gin.Context, data interface{}) {
	location, _ := time.LoadLocation(Constants.TimeLocation)
	responseData := setResponse{
		Code:       http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Data:       data,
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05")}

	g.Header("Content-Type", Constants.ContentTypeJSON)
	g.JSON(http.StatusOK, responseData)
}

func HttpResponseError(g *gin.Context, data interface{}, errModule module_errors.Errors) {
	location, _ := time.LoadLocation(Constants.TimeLocation)
	responseData := setResponse{
		Code:       errModule.SubCode,
		Status:     errModule.Description,
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05"),
		Data:       data,
	}
	g.Header("Content-Type", Constants.ContentTypeJSON)
	g.JSON(errModule.Code, responseData)
}
