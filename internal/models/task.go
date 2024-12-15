package models

import "time"

type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description,omitempty"`
    Status      string    `json:"status"`
    ListID      int       `json:"list_id"`
    CreatedAt   time.Time `json:"created_at"`
    DueDate     time.Time `json:"due_date,omitempty"`
}
