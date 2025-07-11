package routes

import (
	"github.com/Skpanchall/go-job-queue/backend/api"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler *api.JobHandler) {
	router.GET("/submit-job", handler.SubmitJob)
}
