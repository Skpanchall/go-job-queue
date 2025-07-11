package main

import (
	"fmt"

	"github.com/Skpanchall/go-job-queue/backend/core"
)

func main() {

	ch := make(chan core.Job, 10)
	oneJob := core.Job{
		Id:      1,
		Type:    "send email",
		Payload: "&to=sachin@gmail.com",
	}
	// pass onejonb into buffer channel
	ch <- oneJob
	// recieve channe throug value
	oneJobrecieved := <-ch
	fmt.Printf("job received %+v", oneJobrecieved)
}
