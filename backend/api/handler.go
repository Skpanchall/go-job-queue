package api

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/Skpanchall/go-job-queue/backend/core"
	"github.com/Skpanchall/go-job-queue/backend/store"
	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	Queue chan core.Job
	DB    *sql.DB
}

func NewJobHandler(newjob chan core.Job, db *sql.DB) *JobHandler {
	return &JobHandler{
		Queue: newjob,
		DB:    db,
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
	err = req.Insert(h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to save job"})
	}

	// store.SetJobStatus(req.Id, "pending")

	// pass on que a job
	h.Queue <- req

	c.JSON(http.StatusOK, gin.H{
		"status": "queue",
		"job_id": req.Id,
	})

}

func (h *JobHandler) GetJob(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	status := store.GetJobStatus(id)
	c.JSON(http.StatusOK, gin.H{
		"stauts": status,
		"job_id": id,
	})
}
