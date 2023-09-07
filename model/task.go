package model

type Task struct {
	TaskName string   `json:"task_name" bson:"task_name"`
	Status   string   `json:"status" bson:"status"`
	Action   []string `json:"action" bson:"actions"`
}
