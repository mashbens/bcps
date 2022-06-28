package payment

import (
	_member "github.com/mashbens/cps/business/member/entity"
	payment "github.com/mashbens/cps/business/payment/entity"
	_user "github.com/mashbens/cps/business/user/entity"
)

func MemberToPayment(data _member.Membership) payment.Membership {
	return payment.Membership{
		ID:       data.ID,
		Type:     data.Type,
		Price:    data.Price,
		Duration: data.Duration,
	}
}

func UserToPayment(data _user.User) payment.User {
	return payment.User{
		ID:             data.ID,
		Name:           data.Name,
		Email:          data.Email,
		Phone:          data.Phone,
		Member_expired: data.Member_expired,
		Member_type:    data.Member_type,
	}
}
