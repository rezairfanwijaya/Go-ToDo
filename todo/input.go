package todo

type TodoCreateInput struct {
	Title      string `json:"title" binding:"required"`
	ActivityID int    `json:"activity_group_id" binding:"required"`
	IsActive   bool   `json:"is_active" binding:"required"`
}
