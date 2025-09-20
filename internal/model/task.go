package model

type Status string

const (
	TODO        Status = "todo"
	IN_PROGRESS Status = "in-progress"
	DONE        Status = "done"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
