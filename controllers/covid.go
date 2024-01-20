package controllers

import (
	"fmt"
	"net/http"

	"github.com/IamMrTeapot/LWMNAssignment/services"
	"github.com/gin-gonic/gin"
)

func ReportSummary(c *gin.Context) {
	summary, err := services.GetSummary()
	if err != nil {
		fmt.Print(err.Error())
	}
	c.JSON(http.StatusOK, summary)
}
