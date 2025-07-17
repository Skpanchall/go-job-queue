package core

import "database/sql"

type Job struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func (j *Job) Insert(db *sql.DB) error {

	_, err := db.Exec("insert into jobs (id , payload  , status) values (? ,? ,?)", j.Id, j.Payload, j.Type)
	return err
}
