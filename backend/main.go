package main

import (
	"github.com/Skpanchall/go-job-queue/backend/api"
	"github.com/Skpanchall/go-job-queue/backend/core"
	"github.com/Skpanchall/go-job-queue/backend/db"
	"github.com/Skpanchall/go-job-queue/backend/routes"
	"github.com/Skpanchall/go-job-queue/backend/worker"
	"github.com/gin-gonic/gin"
)

func main() {
	jobQueue := make(chan core.Job, 10)

	gin := gin.Default()

	db := db.Init()

	handler := api.NewJobHandler(jobQueue, db)

	worker.StartWorkerPool(jobQueue, 3)

	// setup routes here
	routes.SetupRoutes(gin, handler)
	gin.Run()
	// ch := make(chan core.Job, 10)
	// oneJob := core.Job{
	// 	Id:      1,
	// 	Type:    "send email",
	// 	Payload: "&to=sachin@gmail.com",
	// }
	// // pass onejonb into buffer channel
	// ch <- oneJob
	// // recieve channe throug value
	// oneJobrecieved := <-ch
	// fmt.Printf("job received %+v", oneJobrecieved)
}
