package member

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strconv"
	"time"

	imgBB "github.com/JohnNON/ImgBB"

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

	ImgUpload(b *multipart.FileHeader) string
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

	img := c.ImgUpload(member.ImgBB)

	member.Img = img

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

	img := c.ImgUpload(member.ImgBB)

	member.Img = img

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

func (c *memberService) ImgUpload(file *multipart.FileHeader) string {

	src, err := file.Open()
	if err != nil {
		return fmt.Sprintln(err)
	}

	b, err := io.ReadAll(src)
	if err != nil {
		log.Fatal(err)
	}
	key := "02406488a81ff26d2a22b6306b6b21f9"
	img := imgBB.NewImage(hashSum(b), "60", b)

	bb := imgBB.NewImgBB(key, 5*time.Second)

	r, e := bb.Upload(img)
	if e != nil {
		log.Fatal(e)
	}

	return r.Data.Url
}

func hashSum(b []byte) string {
	sum := md5.Sum(b)
	return hex.EncodeToString(sum[:])
}
