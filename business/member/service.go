package member

import (
	"errors"
	"log"
	"strconv"

	"github.com/mashbens/cps/business/member/entity"
	"github.com/mashbens/cps/business/superadmin"
)

type MemberRepo interface {
	FindMemberByID(memberID string) (entity.Membership, error)
	InserMemberships(member entity.Membership) (entity.Membership, error)
}

type MemberService interface {
	FindMemberTypeByID(memberID string) (*entity.Membership, error)
	CreateMemberships(member entity.Membership) (*entity.Membership, error)
}

type memberService struct {
	memberRepo       MemberRepo
	superAdminSevice superadmin.SuperAdminService
}

func NewMemberService(
	MemberRepo MemberRepo,
	superAdminSevice superadmin.SuperAdminService,
) MemberService {
	return &memberService{
		memberRepo:       MemberRepo,
		superAdminSevice: superAdminSevice,
	}
}

func (c *memberService) FindMemberTypeByID(memberID string) (*entity.Membership, error) {

	member, err := c.memberRepo.FindMemberByID(memberID)
	if err != nil {
		return nil, errors.New("Member type not found------")
	}

	return &member, nil

}

func (c *memberService) CreateMemberships(member entity.Membership) (*entity.Membership, error) {
	strID := strconv.Itoa(member.Super_adminID)
	sAdmin, err := c.superAdminSevice.FindSuperAdminByID(strID)
	if err != nil {
		return nil, err
	}
	_ = sAdmin
	// member.Super_admin = *sAdmin

	log.Println(member)
	m, err := c.memberRepo.InserMemberships(member)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
