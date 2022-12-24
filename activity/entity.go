package activity

import "time"

type Activity struct {
	ID         int       `json:"id,omitempty" gorm:"primaryKey" `
	Title      string    `json:"title,omitempty"`
	Email      string    `json:"email,omitempty"`
	CreatredAt time.Time `json:"updatedAt,omitempty"`
	UpdatedAt  time.Time `json:"createdAt,omitempty"`
}
