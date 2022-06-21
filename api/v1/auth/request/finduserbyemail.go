package request

type FindUserByEmailRequest struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}

func NewFindUserByEmailRequest(req string) FindUserByEmailRequest {
	return FindUserByEmailRequest{
		Email: req,
	}
}
