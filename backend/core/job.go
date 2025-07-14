package core

type Job struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Payload string `json:"payload"`
}
