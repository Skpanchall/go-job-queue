package store

import "sync"

var (
	statusMap  = make(map[int]string)
	statusLock = sync.RWMutex{}
)

func SetJobStatus(id int, status string) {
	statusLock.Lock()
	defer statusLock.Unlock()
	statusMap[id] = status
}

func GetJobStatus(id int) string {
	statusLock.RLock()
	defer statusLock.RUnlock()
	return statusMap[id]

}
