package todo

import (
	"gotodo/activity"
	"time"
)

type Todo struct {
	ID              int                    `json:"id" gorm:"primaryKey"`
	ActivityGroupID int                    `json:"activity_group_id"`
	Title           string                 `json:"title"`
	IsActive        bool                   `json:"is_active"`
	Priority        string                 `json:"priority"`
	ActivityGroup   activity.ActivityGroup `json:"activity"`
	CreatredAt      time.Time              `json:"updatedAt"`
	UpdatedAt       time.Time              `json:"createdAt"`
}
