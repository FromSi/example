package responses

import "time"

type PostResponse struct {
	ID        string     `json:"id"`
	Text      string     `json:"text"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
