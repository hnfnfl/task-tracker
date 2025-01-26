package cmd

type StatusEnum int

const (
	Todo StatusEnum = iota
	InProgress
	Done
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      StatusEnum `json:"status"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}
