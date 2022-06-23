package request

import "github.com/mashbens/cps/business/payment/entity"

type CreatePaymentRequest struct {
	UserID       int `json:"user_id"`
	MembershipID int `json:"membership_id"`
}

func NewCreatePaymentReq(req CreatePaymentRequest) entity.Payment {
	return entity.Payment{
		UserID:       req.UserID,
		MembershipID: req.MembershipID,
	}
}
