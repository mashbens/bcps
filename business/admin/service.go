package admin

import (
	"errors"
	"log"
	"strconv"

	"github.com/mashbens/cps/business/admin/entity"
	"github.com/mashbens/cps/business/superadmin"
	jwtService "github.com/mashbens/cps/business/user"

	"golang.org/x/crypto/bcrypt"
)

type AdminRepo interface {
	InsertAdmin(admin entity.Admin) (entity.Admin, error)
	FindAdminByEmail(email string) (entity.Admin, error)
}

type AdminService interface {
	InsertAdmin(admin entity.Admin) (*entity.Admin, error)
	FindAdminByEmail(email string) (*entity.Admin, error)
	AdminLogin(admin entity.Admin) (*entity.Admin, error)
}

type adminService struct {
	adminRepo        AdminRepo
	jwtService       jwtService.JWTService
	superAdminSevice superadmin.SuperAdminService
}

func NewAdminService(
	adminRepo AdminRepo,
	jwtService jwtService.JWTService,
	superAdminSevice superadmin.SuperAdminService,

) AdminService {
	return &adminService{
		adminRepo:        adminRepo,
		jwtService:       jwtService,
		superAdminSevice: superAdminSevice,
	}
}

func (c *adminService) InsertAdmin(admin entity.Admin) (*entity.Admin, error) {
	log.Println(admin, "asdmingh")
	findSA, err := c.superAdminSevice.FindSuperAdminByID(strconv.Itoa(admin.SuperAdminID))
	if err != nil {
		return nil, errors.New("Invalid Credential")
	}
	_ = findSA

	adm, err := c.adminRepo.FindAdminByEmail(admin.Email)
	if err == nil {
		return nil, errors.New("Admin already exists")
	}
	admin.Password = hashAndSalt([]byte(admin.Password))

	adm, err = c.adminRepo.InsertAdmin(admin)
	if err != nil {
		return nil, err
	}

	token := c.jwtService.GenerateToken((strconv.Itoa(admin.ID)))
	adm.Token = token

	return &adm, nil
}

func (c *adminService) AdminLogin(admin entity.Admin) (*entity.Admin, error) {
	err := c.VerifyCredential(admin.Email, admin.Password)
	if err != nil {
		return nil, errors.New("Invalid username or password")
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
		return nil, errors.New("admin not found")
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
