package resp

import (
	payment "github.com/mashbens/cps/business/payment/entity"
)

type PaymentResp struct {
	ID      int    `json:"id"`
	User    User   `json:"user"`
	Member  Member `json:"membership"`
	Amount  int    `json:"amount"`
	SnapURL string `json:"snap_url,omitempty"`
}

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Member_expired string `json:"member_expired"`
	Member_type    string `json:"member_type"`
}
type Member struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Price    int    `json:"price"`
	Duration int    `json:"duration"`
}

func FromService(payment payment.Payment) PaymentResp {
	return PaymentResp{
		ID: int(payment.ID),
		User: User{
			ID:             payment.User.ID,
			Name:           payment.User.Name,
			Email:          payment.User.Email,
			Phone:          payment.User.Phone,
			Member_expired: payment.User.Member_expired,
			Member_type:    payment.User.Member_type,
		},
		Member: Member{
			ID:       payment.Membership.ID,
			Type:     payment.Membership.Type,
			Price:    payment.Membership.Price,
			Duration: payment.Membership.Duration,
		},
		Amount:  payment.Amount,
		SnapURL: payment.SnapURL,
	}
}
