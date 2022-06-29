package class

import (
	"errors"
	"strconv"

	"github.com/mashbens/cps/business/admin"
	"github.com/mashbens/cps/business/class/entity"
)

type ClassRepo interface {
	FindAllClass(search string) (data []entity.Class)
	FindAllClassOn(search string) (data []entity.Class)
	FindAllClassOff(search string) (data []entity.Class)
	FindClassOnByID(classID string) (entity.Class, error)
	FindClassOffByID(classID string) (entity.Class, error)

	InsertClass(class entity.Class) (entity.Class, error)
	UpdateClass(class entity.Class) (entity.Class, error)
	DeleteClass(classID string) error
}

type ClassService interface {
	FindAllClass(search string) (data []entity.Class)
	FindAllClassOn(search string) (data []entity.Class)
	FindAllClassOff(search string) (data []entity.Class)
	FindClassOnByID(classID string) (*entity.Class, error)
	FindClassOffByID(classID string) (*entity.Class, error)

	InsertClass(class entity.Class) (*entity.Class, error)
	UpdateClass(class entity.Class) (*entity.Class, error)
	DeleteClass(adminId string, classID string) error
}

type clasService struct {
	classRepo    ClassRepo
	adminService admin.AdminService
}

func NewClassService(
	ClassRepo ClassRepo,
	adminService admin.AdminService,
) ClassService {
	return &clasService{
		classRepo:    ClassRepo,
		adminService: adminService,
	}
}
func (c *clasService) InsertClass(class entity.Class) (*entity.Class, error) {
	adminID := strconv.Itoa(class.AdminID)
	admin, err := c.adminService.FindAdminByID(adminID)
	if err != nil {
		return nil, err
	}
	_ = admin
	clas, err := c.classRepo.InsertClass(class)
	if err != nil {
		return nil, err
	}
	return &clas, nil
}

func (c *clasService) UpdateClass(class entity.Class) (*entity.Class, error) {
	admin, err := c.adminService.FindAdminByID(strconv.Itoa(class.AdminID))
	if err != nil {
		return nil, errors.New("Admin not found")
	}
	_ = admin
	class, err = c.classRepo.UpdateClass(class)
	if err != nil {
		return nil, err
	}
	return &class, nil

}

func (c *clasService) DeleteClass(adminID string, classID string) error {
	admin, err := c.adminService.FindAdminByID(adminID)
	if err != nil {
		return nil
	}
	_ = admin
	cls := c.classRepo.DeleteClass(classID)
	_ = cls
	return nil
}

func (c *clasService) FindAllClass(search string) (data []entity.Class) {
	data = c.classRepo.FindAllClass(search)
	return
}

func (c *clasService) FindAllClassOn(search string) (data []entity.Class) {
	data = c.classRepo.FindAllClassOn(search)
	return
}
func (c *clasService) FindAllClassOff(search string) (data []entity.Class) {
	data = c.classRepo.FindAllClassOff(search)
	return
}

func (c *clasService) FindClassOnByID(classID string) (*entity.Class, error) {
	class, err := c.classRepo.FindClassOnByID(classID)
	if err != nil {
		return nil, errors.New("Class not found")
	}

	return &class, nil
}
func (c *clasService) FindClassOffByID(classID string) (*entity.Class, error) {
	class, err := c.classRepo.FindClassOffByID(classID)
	if err != nil {
		return nil, errors.New("Class not found")
	}

	return &class, nil
}
