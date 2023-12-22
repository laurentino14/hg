package entities

import "time"

type OS struct {
	COD        int        `json:"cod_os"`
	ProjectCOD int        `json:"cod_project"`
	Price      float64    `json:"price"`
	Status     int        `json:"status"`
	WorkerID   string     `json:"worker_id"`
	ClientID   string     `json:"client_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type NewOSInput struct {
	ProjectCOD int     `json:"cod_project"`
	WorkerID   string  `json:"worker_id"`
	ClientID   string  `json:"client_id"`
	Price      float64 `json:"price"`
}

type UpdateOSInput struct {
	COD   int     `json:"cod_os"`
	Price float64 `json:"price"`
}

func NewOS(input *NewOSInput) (*OS, error) {
	os := &OS{}
	os.ProjectCOD = input.ProjectCOD
	os.WorkerID = input.WorkerID
	os.ClientID = input.ClientID
	os.Price = input.Price
	os.CreatedAt = time.Now()

	return os, nil
}
