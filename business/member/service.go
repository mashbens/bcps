package member

import (
	"errors"

	"github.com/mashbens/cps/business/member/entity"
)

type MemberRepo interface {
	FindMemberByID(memberID string) (entity.Membership, error)
}

type MemberService interface {
	FindMemberTypeByID(memberID string) (*entity.Membership, error)
}

type memberService struct {
	memberRepo MemberRepo
}

func NewMemberService(MemberRepo MemberRepo) MemberService {
	return &memberService{
		memberRepo: MemberRepo,
	}
}

func (c *memberService) FindMemberTypeByID(memberID string) (*entity.Membership, error) {

	member, err := c.memberRepo.FindMemberByID(memberID)
	if err != nil {
		return nil, errors.New("Member type not found------")
	}

	return &member, nil

}
