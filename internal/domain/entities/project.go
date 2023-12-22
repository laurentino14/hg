package entities

import "time"

type Project struct {
	COD         int        `json:"cod_project"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Type        int16      `json:"type"`
	StartDate   string     `json:"start_date"`
	EndDate     string     `json:"end_date"`
	OwnerID     string     `json:"owner_id"`
	Status      int16      `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type NewProjectInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int16  `json:"type"`
	OwnerID     string `json:"owner_id"`
}

type UpdateProjectInput struct {
	COD         int        `json:"cod_project"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Type        int16      `json:"type"`
	StartDate   string     `json:"start_date"`
	EndDate     string     `json:"end_date"`
	Status      string     `json:"status"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func NewProject(input *NewProjectInput) (*Project, error) {
	project := &Project{}

	project.Name = input.Name
	project.Description = input.Description
	project.Type = input.Type
	project.OwnerID = input.OwnerID
	project.CreatedAt = time.Now()

	return project, nil
}
