package todo

import (
	"gotodo/activity"
	"time"
)

type Todo struct {
	ID         int               `json:"id" gorm:"primaryKey"`
	Title      string            `json:"title"`
	ActivityID int               `json:"activity_group_id" gorm:"column:activity_group_id"`
	IsActive   bool              `json:"is_active"`
	Priority   string            `json:"priority"`
	Activity   activity.Activity `json:"activity"`
	CreatredAt time.Time         `json:"updatedAt"`
	UpdatedAt  time.Time         `json:"createdAt"`
}
