package response

import "time"

type NewsRes struct {
	ID        uint      `json:"id"`
	Image     string    `json:"image"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
