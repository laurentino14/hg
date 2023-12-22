package entities

import "time"

type Task struct {
	COD        int        `json:"cod_task"`
	Name       string     `json:"name"`
	Status     int        `json:"status"`
	OwnerID    string     `json:"owner_id"`
	ProjectCOD string     `json:"project_cod"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type NewTaskInput struct {
	Name       string `json:"name"`
	OwnerID    string `json:"owner_id"`
	ProjectCOD string `json:"project_cod"`
}

type UpdateTaskInput struct {
	COD    int    `json:"cod_task"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func NewTask(input *NewTaskInput) (*Task, error) {
	task := &Task{}

	task.Name = input.Name
	task.Status = 0
	task.OwnerID = input.OwnerID
	task.ProjectCOD = input.ProjectCOD
	task.CreatedAt = time.Now()

	return task, nil
}
