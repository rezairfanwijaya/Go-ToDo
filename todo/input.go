package todo

type TodoCreateInput struct {
	Title      string `json:"title" binding:"required"`
	ActivityID int    `json:"activity_group_id" binding:"required"`
	IsActive   bool   `json:"is_active" binding:"required"`
}

type TodoUpdateInput struct {
	Title    string `json:"title" binding:"required"`
	Priority string `json:"priority" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
	Status   string `json:"status" binding:"required"`
}
