package worker

import (
	"fmt"
	"time"

	"github.com/Skpanchall/go-job-queue/backend/core"
	"github.com/Skpanchall/go-job-queue/backend/store"
)

func StartWorkerPool(queue chan core.Job, numOfWorker int) {

	for i := 0; i < numOfWorker; i++ {
		workerId := i
		go func() {
			for {
				job := <-queue
				store.SetJobStatus(job.Id, "processing")
				fmt.Printf("[ Worker %d] Picked job id , %d", workerId, job.Id)

				time.Sleep(3 * time.Second)
				store.SetJobStatus(job.Id, "done")
				fmt.Printf("[ Worker %d] Done Processing  job id , %d", workerId, job.Id)
			}

		}()
		// i++
	}

}
