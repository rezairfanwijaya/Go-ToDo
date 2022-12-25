package activity

type ActivityCreateInput struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type ActivityUpdateInput struct {
	Title string `json:"title" binding:"required"`
}
