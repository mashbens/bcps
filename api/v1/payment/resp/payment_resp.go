package resp

import (
	payment "github.com/mashbens/cps/business/payment/entity"
	// member "github.com/mashbens/cps/business/member/entity"
	// user "github.com/mashbens/cps/business/user/entity"

	member "github.com/mashbens/cps/api/v1/member/resp"
	user "github.com/mashbens/cps/api/v1/user/resp"
)

type PaymentResp struct {
	ID      int               `json:"id"`
	User    user.UserResp     `json:"user"`
	Member  member.MemberResp `json:"membership"`
	Amount  int               `json:"amount"`
	SnapURL string            `json:"snap_url,omitempty"`
}

func FromService(payment payment.Payment) PaymentResp {
	return PaymentResp{
		ID: int(payment.ID),
		User: user.UserResp{
			ID:             payment.User.ID,
			Name:           payment.User.Name,
			Email:          payment.User.Email,
			Phone:          payment.User.Phone,
			Member_expired: payment.User.Member_expired,
			Member_type:    payment.User.Member_type,
		},
		Member: member.MemberResp{
			ID:       payment.Membership.ID,
			Type:     payment.Membership.Type,
			Price:    payment.Membership.Price,
			Duration: payment.Membership.Duration,
		},
		Amount:  payment.Amount,
		SnapURL: payment.SnapURL,
	}
}
