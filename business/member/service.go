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
	UpdateMemberType(member entity.Membership) (entity.Membership, error)
	FindAllMemberType(title string) (data []entity.Membership)
	DeleteMemberType(memberID string) error
}

type MemberService interface {
	FindMemberTypeByID(memberID string) (*entity.Membership, error)
	CreateMemberships(member entity.Membership) (*entity.Membership, error)
	UpdateMemberType(member entity.Membership) (*entity.Membership, error)
	FindAllMemberType(search string) (data []entity.Membership)
	DeleteMemberType(adminID string, memberID string) error
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

func (c *memberService) FindAllMemberType(search string) (data []entity.Membership) {
	data = c.memberRepo.FindAllMemberType(search)
	return
}

func (c *memberService) FindMemberTypeByID(memberID string) (*entity.Membership, error) {

	member, err := c.memberRepo.FindMemberByID(memberID)
	if err != nil {
		return nil, errors.New("Member type not found")
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

	m, err := c.memberRepo.InserMemberships(member)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (c *memberService) UpdateMemberType(member entity.Membership) (*entity.Membership, error) {

	sAdmin, err := c.superAdminSevice.FindSuperAdminByID(strconv.Itoa(member.Super_adminID))
	if err != nil {
		return nil, errors.New("Super admin not found")
	}
	_ = sAdmin

	m, err := c.FindMemberTypeByID(strconv.Itoa(member.ID))
	if err != nil {
		return nil, err
	}
	_ = m

	member, err = c.memberRepo.UpdateMemberType(member)
	if err != nil {
		return nil, err
	}

	return &member, nil
}

func (c *memberService) DeleteMemberType(adminID string, memberID string) error {
	sAdmin, err := c.superAdminSevice.FindSuperAdminByID(adminID)
	if err != nil {
		return nil
	}
	_ = sAdmin

	m := c.memberRepo.DeleteMemberType(memberID)
	if err != nil {
		return nil
	}
	log.Println(memberID)
	_ = m

	return nil
}
