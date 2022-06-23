package resp

import "github.com/mashbens/cps/business/payment/entity"

type PaymentResp struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	MembeshipID int    `json:"membership_id"`
	Amount      int    `json:"amount"`
	SnapURL     string `json:"snap_url,omitempty"`
}

func FromService(payment entity.Payment) PaymentResp {
	return PaymentResp{
		ID:          int(payment.ID),
		UserID:      payment.UserID,
		MembeshipID: payment.MembershipID,
		Amount:      payment.Amount,
		SnapURL:     payment.SnapURL,
	}
}
