package admin_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/mashbens/cps/business/admin"
	adminEntity "github.com/mashbens/cps/business/admin/entity"
	jwtService "github.com/mashbens/cps/business/user"
)

var service admin.AdminService
var admin1, admin2 adminEntity.Admin

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindAdminByID(t *testing.T) {
	t.Run("Expect to find admin by id", func(t *testing.T) {
		adminID := int(admin1.ID)
		adminIDs := strconv.Itoa(adminID)
		admin, err := service.FindAdminByID(adminIDs)
		if err != nil {
			t.Error("Expect to find admin by id", err)
		} else {
			if admin.ID != 0 {
				t.Errorf("Expected %d, got %d", 0, admin.ID)
			}
		}
	})
	t.Run("Expect not found the content", func(t *testing.T) {
		adminID := int(admin1.ID)
		adminIDs := strconv.Itoa(adminID)
		admin, err := service.FindAdminByID(adminIDs)

		if err != nil {
			t.Error("Expect error is nil. Error: ", err)
		} else if admin == nil {
			t.Error("Expect to find admin by id", err)

		}
	})
}

func TestFindAdminByEmail(t *testing.T) {
	t.Run("Expect to find admin by email", func(t *testing.T) {
		admin, err := service.FindAdminByEmail(admin1.Email)
		if err != nil {
			t.Error("Expect to find admin by email", err)
		} else {
			if admin.ID != 1 {
				t.Errorf("Expected %d, got %d", 0, admin.ID)
			}
		}
	})
	t.Run("Expect not found the content", func(t *testing.T) {
		admin, err := service.FindAdminByEmail("random@mail.com")

		if err != nil {
			t.Error("Expect error is nil. Error: ", err)
		} else if admin == nil {
			t.Error("Expect to find admin by id", err)

		}
	})
}

func TestInserAdmin(t *testing.T) {
	t.Run("Expect to Login admin ", func(t *testing.T) {
		_, err := service.InsertAdmin(admin1)
		if err != nil {
			t.Error("Expext error is invalid email or password. Error: ", err)
			t.FailNow()
		}
	})
}

func TestAdminLogin(t *testing.T) {
	t.Run("Expect to Login admin ", func(t *testing.T) {
		_, err := service.AdminLogin(admin1)
		if err != nil {
			t.Error("Expext error is invalid email or password. Error: ", err)
			t.FailNow()
		}
	})
}

func TestFindAllAdmins(t *testing.T) {
	t.Run("Expect to Login admin ", func(t *testing.T) {
		sadminID := int(admin1.SuperAdminID)
		sadminIDs := strconv.Itoa(sadminID)
		err := service.FindAllAdmins(sadminIDs, "")
		if err == nil {
			t.Error("Expext error is invalid email or password. Error: ", err)
		}
	})
}

func TestUpdateAdmin(t *testing.T) {
	t.Run("Expect to Login admin ", func(t *testing.T) {
		_, err := service.UpdateAdmin(admin1)
		if err != nil {
			t.Error("Expext error is invalid email or password. Error: ", err)
			t.FailNow()
		}
	})
}
func TestDeleteAdmin(t *testing.T) {
	t.Run("Expect to Login admin ", func(t *testing.T) {
		adminID := int(admin1.ID)
		adminIDs := strconv.Itoa(adminID)
		err := service.DeleteAdmin(adminIDs, adminIDs)
		if err != nil {
			t.Error("Expext error is invalid email or password. Error: ", err)
			t.FailNow()
		}
	})
}

func setup() {
	admin1.ID = 1
	admin1.Name = "John"
	admin1.Password = "test123"
	admin1.Email = "admin@mail.com"
	admin1.SuperAdminID = 1

	admin2.ID = 2
	admin2.Name = "John"
	admin2.Password = "test123"

	repo := newInMemoryRepository()
	jwtService := jwtService.NewJWTService()
	service = admin.NewAdminService(repo, jwtService)
}

type inMemoryRepository struct {
	adminB       map[string]adminEntity.Admin
	adminByEmail map[string]adminEntity.Admin
}

func newInMemoryRepository() *inMemoryRepository {
	var repo inMemoryRepository
	repo.adminB = make(map[string]adminEntity.Admin)
	repo.adminByEmail = make(map[string]adminEntity.Admin)

	userID := int64(admin1.ID)
	userIDs := strconv.FormatInt(userID, 10)
	repo.adminB[userIDs] = admin1
	repo.adminByEmail[admin1.Email] = admin1

	return &repo
}
func (r *inMemoryRepository) InsertAdmin(admin adminEntity.Admin) (adminEntity.Admin, error) {
	return adminEntity.Admin{}, nil
}

func (r *inMemoryRepository) FindAdminByEmail(email string) (adminEntity.Admin, error) {
	return r.adminByEmail[email], nil
}

func (r *inMemoryRepository) FindAdminByID(id string) (adminEntity.Admin, error) {
	return adminEntity.Admin{}, nil
}

func (r *inMemoryRepository) FindAllAdmins(search string) (data []adminEntity.Admin) {
	return []adminEntity.Admin{}
}

func (r *inMemoryRepository) UpdateAdmin(admin adminEntity.Admin) (adminEntity.Admin, error) {
	return adminEntity.Admin{}, nil
}
func (r *inMemoryRepository) DeleteAdmin(adminID string) error {
	return nil
}
