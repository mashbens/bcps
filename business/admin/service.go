package admin

import (
	// "errors"
	"errors"
	"log"
	"strconv"

	"github.com/mashbens/cps/business/admin/entity"
	jwtService "github.com/mashbens/cps/business/user"

	"golang.org/x/crypto/bcrypt"
)

type AdminRepo interface {
	InsertAdmin(admin entity.Admin) (entity.Admin, error)
	FindAdminByEmail(email string) (entity.Admin, error)
	FindAdminByID(id string) (entity.Admin, error)
	FindAllAdmins(search string) (data []entity.Admin)
	UpdateAdmin(admin entity.Admin) (entity.Admin, error)
	DeleteAdmin(adminID string) error
}

type AdminService interface {
	InsertAdmin(admin entity.Admin) (*entity.Admin, error)
	FindAdminByEmail(email string) (*entity.Admin, error)
	AdminLogin(admin entity.Admin) (*entity.Admin, error)
	FindAdminByID(adminID string) (*entity.Admin, error)
	FindAllAdmins(sAdminID string, search string) (data []entity.Admin)
	UpdateAdmin(admin entity.Admin) (*entity.Admin, error)
	DeleteAdmin(sAdminID string, adminID string) error
}

type adminService struct {
	adminRepo  AdminRepo
	jwtService jwtService.JWTService
}

func NewAdminService(
	adminRepo AdminRepo,
	jwtService jwtService.JWTService,

) AdminService {
	return &adminService{
		adminRepo:  adminRepo,
		jwtService: jwtService,
	}
}

func (c *adminService) InsertAdmin(admin entity.Admin) (*entity.Admin, error) {
	adm, err := c.adminRepo.FindAdminByEmail(admin.Email)
	if err == nil {
		return nil, errors.New("admin already exists")
	}
	admin.Password = hashAndSalt([]byte(admin.Password))
	adm, err = c.adminRepo.InsertAdmin(admin)
	if err != nil {
		return nil, err
	}
	return &adm, nil
}

func (c *adminService) AdminLogin(admin entity.Admin) (*entity.Admin, error) {
	err := c.VerifyCredential(admin.Email, admin.Password)
	if err != nil {
		return nil, errors.New("Invalid email or password")
	}
	adm, _ := c.adminRepo.FindAdminByEmail(admin.Email)
	token := c.jwtService.GenerateToken(strconv.Itoa(adm.ID))
	adm.Token = token
	return &adm, nil
}

func (c *adminService) VerifyCredential(email string, password string) error {
	user, err := c.FindAdminByEmail(email)
	if err != nil {
		println(err.Error())
		return err
	}
	isValidPassword := comparePassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}
	return nil
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (c *adminService) FindAdminByEmail(email string) (*entity.Admin, error) {
	admin, err := c.adminRepo.FindAdminByEmail(email)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (c *adminService) FindAdminByID(adminID string) (*entity.Admin, error) {
	admin, err := c.adminRepo.FindAdminByID(adminID)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func (c *adminService) FindAllAdmins(sAdminID string, search string) (data []entity.Admin) {
	data = c.adminRepo.FindAllAdmins(search)
	return
}
func (c *adminService) UpdateAdmin(admin entity.Admin) (*entity.Admin, error) {
	a, err := c.FindAdminByID(strconv.Itoa(admin.ID))
	if err != nil {
		return nil, err
	}
	adm, err := c.adminRepo.UpdateAdmin(*a)
	if err != nil {
		return nil, err
	}
	return &adm, nil
}

func (c *adminService) DeleteAdmin(sAdminID string, adminID string) error {
	admin, err := c.FindAdminByID(adminID)
	if err != nil {
		return err
	}
	_ = c.adminRepo.DeleteAdmin(strconv.Itoa(admin.ID))
	return err
}
