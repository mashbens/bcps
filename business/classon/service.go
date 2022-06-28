package classon

import (
	"errors"
	"strconv"

	"github.com/mashbens/cps/business/admin"
	"github.com/mashbens/cps/business/classon/entity"
)

type ClassOnRepo interface {
	FindClassOnByID(classID string) (entity.Classon, error)
	InserClassOn(class entity.Classon) (entity.Classon, error)
	FindAllClassOn(search string) (data []entity.Classon)
	UpdateClassOn(class entity.Classon) (entity.Classon, error)
	DeleteClassOn(classID string) error
}

type ClassOnService interface {
	FindClassOnByID(classID string) (*entity.Classon, error)
	// FindClassOnByAdmin(adminID string, classID string) (*entity.Classon, error)
	InserClassOn(class entity.Classon) (*entity.Classon, error)
	FindAllClassOn(search string) (data []entity.Classon)
	UpdateClassOn(class entity.Classon) (*entity.Classon, error)
	DeleteClassOn(adminId string, classID string) error
}

type clasOnService struct {
	classOnRepo  ClassOnRepo
	adminService admin.AdminService
}

func NewClassOnService(
	ClassOnRepo ClassOnRepo,
	adminService admin.AdminService,
) ClassOnService {
	return &clasOnService{
		classOnRepo:  ClassOnRepo,
		adminService: adminService,
	}
}

func (c *clasOnService) FindClassOnByID(classID string) (*entity.Classon, error) {
	class, err := c.classOnRepo.FindClassOnByID(classID)
	if err != nil {
		return nil, errors.New("Class not found")
	}

	return &class, nil
}
func (c *clasOnService) InserClassOn(class entity.Classon) (*entity.Classon, error) {
	adminID := strconv.Itoa(class.AdminID)
	admin, err := c.adminService.FindAdminByID(adminID)
	if err != nil {
		return nil, err
	}
	_ = admin

	clas, err := c.classOnRepo.InserClassOn(class)
	if err != nil {
		return nil, err
	}

	return &clas, nil
}

func (c *clasOnService) FindAllClassOn(search string) (data []entity.Classon) {
	data = c.classOnRepo.FindAllClassOn(search)
	return
}

func (c *clasOnService) UpdateClassOn(class entity.Classon) (*entity.Classon, error) {
	admin, err := c.adminService.FindAdminByID(strconv.Itoa(class.AdminID))
	if err != nil {
		return nil, errors.New("Admin not found")
	}
	_ = admin

	class, err = c.classOnRepo.UpdateClassOn(class)
	if err != nil {
		return nil, err
	}

	return &class, nil

}
func (c *clasOnService) DeleteClassOn(adminID string, classID string) error {
	admin, err := c.adminService.FindAdminByID(adminID)
	if err != nil {
		return nil
	}
	_ = admin

	cls := c.classOnRepo.DeleteClassOn(classID)

	_ = cls

	return nil
}

// func (c *clasOnService) FindClassOnByAdmin(adminID string, classID string) (*entity.Classon, error) {
// 	admin, err := c.adminService.FindAdminByID(adminID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	_ = admin

// 	class, err := c.classOnRepo.FindClassOnByID(classID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &class, nil

// }
