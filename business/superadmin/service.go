package superadmin

import (
	"errors"
	"log"
	"strconv"

	"github.com/mashbens/cps/business/superadmin/entity"
	jwtService "github.com/mashbens/cps/business/user"
	"golang.org/x/crypto/bcrypt"
)

type SuperAdminRepo interface {
	InsertSuperAdmin(sAdmin entity.SuperAdmin) (entity.SuperAdmin, error)
	FindSuperAdminByName(name string) (entity.SuperAdmin, error)
	FindSuperAdminByID(id string) (entity.SuperAdmin, error)
}

type SuperAdminService interface {
	CreateSuperAdmin(sAdmin entity.SuperAdmin) (*entity.SuperAdmin, error)
	FindSuperAdminByName(name string) (*entity.SuperAdmin, error)
	FindSuperAdminByID(id string) (*entity.SuperAdmin, error)
	Login(sAdmin entity.SuperAdmin) (*entity.SuperAdmin, error)
}

type superAdminSevice struct {
	superAdminRepo SuperAdminRepo
	jwtService     jwtService.JWTService
}

func NewSuperAdminService(
	superAdminRepo SuperAdminRepo,
	jwtService jwtService.JWTService,

) SuperAdminService {
	return &superAdminSevice{
		superAdminRepo: superAdminRepo,
		jwtService:     jwtService,
	}
}

func (c *superAdminSevice) CreateSuperAdmin(sAdmin entity.SuperAdmin) (*entity.SuperAdmin, error) {

	admin, err := c.superAdminRepo.FindSuperAdminByName(sAdmin.Name)
	if err == nil {
		return nil, errors.New("Super admin already exists")
	}

	sAdmin.Password = hashAndSalt([]byte(sAdmin.Password))

	admin, err = c.superAdminRepo.InsertSuperAdmin(sAdmin)
	if err != nil {
		return nil, err
	}

	token := c.jwtService.GenerateToken((strconv.Itoa(sAdmin.ID)))
	admin.Token = token

	return &admin, nil

}

func (c *superAdminSevice) Login(admin entity.SuperAdmin) (*entity.SuperAdmin, error) {
	// verify credential

	err := c.VerifyCredential(admin.Name, admin.Password)
	if err != nil {
		return nil, errors.New("Invalid username or password")
	}

	usr, _ := c.superAdminRepo.FindSuperAdminByName(admin.Name)

	token := c.jwtService.GenerateToken((strconv.Itoa(usr.ID)))
	usr.Token = token
	return &usr, nil
}

func (c *superAdminSevice) FindSuperAdminByName(name string) (*entity.SuperAdmin, error) {

	admin, err := c.superAdminRepo.FindSuperAdminByName(name)
	if err != nil {
		return nil, errors.New("Super admin not found")
	}

	return &admin, nil
}

func (c *superAdminSevice) FindSuperAdminByID(id string) (*entity.SuperAdmin, error) {

	admin, err := c.superAdminRepo.FindSuperAdminByID(id)
	if err != nil {
		return nil, errors.New("Super admin not found")
	}

	return &admin, nil
}

func (c *superAdminSevice) VerifyCredential(name string, password string) error {
	user, err := c.FindSuperAdminByName(name)
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

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
