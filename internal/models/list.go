package models

type List struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description,omitempty"`
    GroupID     int    `json:"group_id,omitempty"`
    Tasks       []Task `json:"tasks,omitempty"`
}
