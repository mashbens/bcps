package classoff

import (
	"errors"
	"strconv"

	"github.com/mashbens/cps/business/admin"
	"github.com/mashbens/cps/business/classoff/entity"
)

type ClassOffRepo interface {
	FindClassOffByID(classID string) (entity.Classoff, error)
	InserClassOff(class entity.Classoff) (entity.Classoff, error)
	FindAllClassOff(search string) (data []entity.Classoff)
	UpdateClassOff(class entity.Classoff) (entity.Classoff, error)
	DeleteClassOff(classID string) error
}

type ClassOffService interface {
	FindClassOffByID(classID string) (*entity.Classoff, error)
	InserClassOff(class entity.Classoff) (*entity.Classoff, error)
	FindAllClassOff(search string) (data []entity.Classoff)
	UpdateClassOff(class entity.Classoff) (*entity.Classoff, error)
	DeleteClassOff(adminId string, classID string) error
}

type clasOffService struct {
	classOffRepo ClassOffRepo
	adminService admin.AdminService
}

func NewClassOffService(
	ClassOffRepo ClassOffRepo,
	adminService admin.AdminService,
) ClassOffService {
	return &clasOffService{
		classOffRepo: ClassOffRepo,
		adminService: adminService,
	}
}

func (c *clasOffService) FindClassOffByID(classID string) (*entity.Classoff, error) {
	class, err := c.classOffRepo.FindClassOffByID(classID)
	if err != nil {
		return nil, errors.New("Class not found")
	}

	return &class, nil
}
func (c *clasOffService) InserClassOff(class entity.Classoff) (*entity.Classoff, error) {
	adminID := strconv.Itoa(class.AdminID)
	admin, err := c.adminService.FindAdminByID(adminID)
	if err != nil {
		return nil, err
	}
	_ = admin

	clas, err := c.classOffRepo.InserClassOff(class)
	if err != nil {
		return nil, err
	}

	return &clas, nil
}

func (c *clasOffService) FindAllClassOff(search string) (data []entity.Classoff) {
	data = c.classOffRepo.FindAllClassOff(search)
	return
}

func (c *clasOffService) UpdateClassOff(class entity.Classoff) (*entity.Classoff, error) {
	admin, err := c.adminService.FindAdminByID(strconv.Itoa(class.AdminID))
	if err != nil {
		return nil, errors.New("Admin not found")
	}
	_ = admin

	class, err = c.classOffRepo.UpdateClassOff(class)
	if err != nil {
		return nil, err
	}

	return &class, nil

}
func (c *clasOffService) DeleteClassOff(adminID string, classID string) error {
	admin, err := c.adminService.FindAdminByID(adminID)
	if err != nil {
		return nil
	}
	_ = admin

	cls := c.classOffRepo.DeleteClassOff(classID)

	_ = cls

	return nil
}
