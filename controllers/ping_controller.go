package controllers

import (
	"net/http"

	"github.com/drpaneas/drew/services"
	"github.com/gin-gonic/gin"
)

// PingController is sending 'pong' if return code is 200
func PingController(c *gin.Context) {
	result, err := services.PingServiceVar.PingService()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.String(http.StatusOK, result)
	}
}
