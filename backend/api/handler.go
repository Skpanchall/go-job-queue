package api

import (
	"net/http"
	"time"

	"github.com/Skpanchall/go-job-queue/backend/core"
	"github.com/gin-gonic/gin"
)

type JobHandler struct{ Queue chan core.Job }

func NewJobHandler(newjob chan core.Job) *JobHandler {
	return &JobHandler{
		Queue: newjob,
	}

}

func (h *JobHandler) SubmitJob(c *gin.Context) {
	var req core.Job
	// bind to json to our job struct
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	// make a new id
	req.Id = int(time.Now().UnixNano())

	// pass on que a job
	h.Queue <- req

	c.JSON(http.StatusOK, gin.H{
		"status": "queue",
		"job_id": req.Id,
	})

}
