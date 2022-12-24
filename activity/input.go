package activity

type ActivityCreateInput struct {
	Title string `json:"title" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type ActivityUpdateInput struct {
	Title string `json:"title" binding:"required"`
}
