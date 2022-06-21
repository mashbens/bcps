package request

type EmailVerificationRequest struct {
	Email string `json:"email"`
}

func NewEmailVerificationRequest(req EmailVerificationRequest) string {
	return req.Email
}
