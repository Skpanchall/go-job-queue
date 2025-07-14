package worker

import (
	"fmt"
	"time"

	"github.com/Skpanchall/go-job-queue/backend/core"
)

func StartWorkerPool(queue chan core.Job, numOfWorker int) {
	i := 0
	for i < numOfWorker {
		go func() {
			for {
				job := <-queue
				fmt.Printf("[ Worker %d] Picked job id , %d", i, job.Id)

				time.Sleep(2 * time.Second)
				fmt.Printf("[ Worker %d] Done Processing  job id , %d", i, job.Id)
			}

		}()
		i++
	}

}
