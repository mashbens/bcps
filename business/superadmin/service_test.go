package superadmin_test

import (
	"fmt"
	"testing"

	mocks "github.com/mashbens/cps/business/mocks/superadmin"
	"github.com/mashbens/cps/business/superadmin"
	superadminEntity "github.com/mashbens/cps/business/superadmin/entity"
	"github.com/mashbens/cps/business/user"
	"github.com/stretchr/testify/mock"
)

func TestInsertSuperAdmin(t *testing.T) {
	mcokAdminRepo := new(mocks.SuperAdminRepo)
	mockAdmin := &superadminEntity.SuperAdmin{
		ID:       1,
		Name:     "admin1",
		Password: "admin123",
	}
	t.Run("Failed", func(t *testing.T) {
		mcokAdminRepo.On("FindSuperAdminByName", mock.Anything, mock.Anything).Return(mcokAdminRepo, nil)
		jwtService := user.NewJWTService()
		service := superadmin.NewSuperAdminService(mcokAdminRepo, jwtService)
		a, err := service.CreateSuperAdmin(*mockAdmin)
		fmt.Println(a)
		fmt.Println(err)
		// assert.NoError(t, err)
		// assert.NotNil(t, a)
		// mcokAdminRepo.AssertExpectations(t)
	})
}
