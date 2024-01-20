package server

import (
	"github.com/IamMrTeapot/LWMNAssignment/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/covid/summary", controllers.ReportSummary)

	return router
}
