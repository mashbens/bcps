package resp

import "github.com/mashbens/cps/business/user/entity"

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Member_expired string `json:"member_expired"`
	Member_type    string `json:"member_type"`
	Token          string `json:"token,omitempty"`
	Totp           string `json:"totp,omitempty"`
}

func FromService(user entity.User) User {
	return User{
		ID:             int(user.ID),
		Name:           user.Name,
		Email:          user.Email,
		Phone:          user.Phone,
		Member_expired: user.Member_expired,
		Member_type:    user.Member_type,
		Token:          user.Token,
		Totp:           user.Totp,
	}
}
