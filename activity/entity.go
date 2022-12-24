package activity

import "time"

type Activity struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title"`
	Email      string    `json:"email"`
	CreatredAt time.Time `json:"updatedAt"`
	UpdatedAt  time.Time `json:"createdAt"`
}
