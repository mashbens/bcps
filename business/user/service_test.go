package user_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	mocks "github.com/mashbens/cps/business/mocks/user"
	"github.com/mashbens/cps/business/user"
	userEntity "github.com/mashbens/cps/business/user/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByEmail(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		ID:       1,
		Name:     "hello",
		Email:    "hello@mail.com",
		Phone:    "0898123747",
		Password: "hello123",
	}
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(mockUser, nil).Once()
		service := user.NewUserService(mockUserRepo)
		a, err := service.FindUserByEmail(mockUser.Email)

		assert.NoError(t, err)
		assert.NotNil(t, a)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("user not found")).Once()
		service := user.NewUserService(mockUserRepo)
		a, err := service.FindUserByEmail(mockUser.Email)

		assert.Error(t, err)
		assert.Nil(t, a)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		ID:       1,
		Name:     "hello",
		Email:    "hello@mail.com",
		Phone:    "0898123747",
		Password: "hello123",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByUserID", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		service := user.NewUserService(mockUserRepo)
		a, err := service.FindUserByID(mockUser.ID)
		assert.NoError(t, err)
		assert.NotNil(t, a)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("FindByUserID", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		service := user.NewUserService(mockUserRepo)
		a, err := service.FindUserByID(mockUser.ID)

		assert.Error(t, err)
		assert.Nil(t, a)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		ID:       2,
		Name:     "hello",
		Email:    "hello@mail.com",
		Phone:    "0898123747",
		Password: "hello123",
	}

	t.Run("Err no user exist", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Record not found"))
		service := user.NewUserService(mockUserRepo)
		a, err := service.FindUserByEmail(mockUser.Email)
		assert.Error(t, err)
		assert.Nil(t, a)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("InsertUser", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		service := user.NewUserService(mockUserRepo)
		a, err := service.CreateUser(*mockUser)
		assert.NoError(t, err)
		assert.NotNil(t, a)
		mockUserRepo.AssertExpectations(t)

	})

}

func TestUpdateUserExpiry(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		ID:             2,
		Name:           "hello",
		Email:          "hello@mail.com",
		Phone:          "0898123747",
		Password:       "hello123",
		Member_expired: "12",
		Member_type:    "1",
	}
	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("UpdateUserExpiry", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		service := user.NewUserService(mockUserRepo)
		err := service.UpdateUserExpiry(mockUser.ID, mockUser.Member_expired, mockUser.Member_type)
		assert.Equal(t, nil, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		ID:             2,
		Name:           "hello",
		Email:          "hello@mail.com",
		Phone:          "0898123747",
		Password:       "hello123",
		Member_expired: "12",
		Member_type:    "1",
	}
	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		service := user.NewUserService(mockUserRepo)
		a, err := service.UpdateUser(*mockUser)
		assert.NoError(t, err)
		assert.NotNil(t, a)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		service := user.NewUserService(mockUserRepo)
		a, err := service.UpdateUser(*mockUser)
		assert.Error(t, err)
		assert.Nil(t, a)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestResetPassword(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		ID:             2,
		Name:           "hello",
		Email:          "hello@mail.com",
		Phone:          "0898123747",
		Password:       "hello123",
		Member_expired: "12",
		Member_type:    "1",
	}
	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("ResetPassword", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		service := user.NewUserService(mockUserRepo)
		a, err := service.ResetPassword(*mockUser)
		assert.NoError(t, err)
		assert.NotNil(t, a)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("ResetPassword", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		service := user.NewUserService(mockUserRepo)
		a, err := service.ResetPassword(*mockUser)
		assert.Error(t, err)
		assert.Nil(t, a)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestGetSecreteKey(t *testing.T) {
	mockJwt := new(mocks.JWTService)
	mockUser := &userEntity.User{
		ID:             2,
		Name:           "hello",
		Email:          "hello@mail.com",
		Phone:          "0898123747",
		Password:       "hello123",
		Member_expired: "12",
		Member_type:    "1",
	}
	t.Run("Success", func(t *testing.T) {
		mockJwt.On("GetSecreteKey", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		service := user.NewJWTService()
		err := service.GenerateToken("1")
		assert.NoError(t, nil, err)

	})
	t.Run("Failed", func(t *testing.T) {
		mockJwt.On("GetSecreteKey", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		service := user.NewJWTService()
		err := service.GenerateToken("1")
		assert.NoError(t, nil, err)
	})
}
func TestValidateToken(t *testing.T) {
	type CustomContext struct {
		echo.Context
	}
	mockJwt := new(mocks.JWTService)
	mockUser := &userEntity.User{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsImV4cCI6MTY4OTUzNDIyNSwiaWF0IjoxNjU3OTk4MjI1LCJpc3MiOiJhZG1pbiJ9.MUeg5Vprk0STJqkrtbKNpmeCFabs8mB4NmGcWJR0Khk",
	}
	t.Run("Success", func(t *testing.T) {
		mockJwt.On("GetSecreteKey", mock.Anything, mock.Anything).Return(mockUser, nil).Once()

		service := user.NewJWTService()
		err := service.ValidateToken(mockUser.Token, CustomContext{})
		assert.NoError(t, nil, err)

	})
	t.Run("Failed", func(t *testing.T) {
		mockJwt.On("GetSecreteKey", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		service := user.NewJWTService()
		err := service.ValidateToken(mockUser.Token, CustomContext{})
		assert.NoError(t, nil, err)

	})
}

func TestRegister(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		Name:     "hello",
		Email:    "hello@mail.com",
		Phone:    "0898123747",
		Password: "hello123",
	}
	t.Run("Err Expected", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Record not found"))
		service := user.NewUserService(mockUserRepo)
		a, err := service.FindUserByEmail(mockUser.Email)
		assert.Error(t, err)
		assert.Nil(t, a)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("InsertUser", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		service := user.NewUserService(mockUserRepo)
		a, err := service.CreateUser(*mockUser)
		assert.NoError(t, err)
		assert.NotNil(t, a)
		mockUserRepo.AssertExpectations(t)

	})
	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("InsertUser", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		service := user.NewUserService(mockUserRepo)
		_, err := service.CreateUser(*mockUser)
		assert.Equal(t, nil, err)
	})

}

func TestLogin(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		Name:     "hello",
		Email:    "hello@mail.com",
		Phone:    "0898123747",
		Password: "hello123",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(mockUser, nil).Once()
		service := user.NewUserService(mockUserRepo)
		a, err := service.FindUserByEmail(mockUser.Email)

		assert.NoError(t, err)
		assert.NotNil(t, a)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		a, _ := service.Login(*mockUser)
		assert.Nil(t, a)

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		service := user.NewJWTService()
		err := service.GenerateToken("1")
		assert.NotNil(t, err)
	})
	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("Login", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		a, _ := service.Login(*mockUser)
		assert.Nil(t, a)
	})

}
func TestSendEmailForgotPassword(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		Name:     "hello",
		Email:    "hello@mail.com",
		Phone:    "0898123747",
		Password: "hello123",
		Totp:     "398123",
	}

	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		_, err := service.SendEmailForgotPassword(*mockUser)
		fmt.Print(err)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		_, err := service.SendEmailForgotPassword(*mockUser)
		assert.Error(t, err)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("SendOTPtoEmail", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		err := service.SendOTPtoEmail(mockUser.Totp, mockUser.Name, mockUser.Email)
		assert.NoError(t, err)
	})
}

func TestResetPasswordAuth(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		Name:     "hello",
		Email:    "hello@mail.com",
		Phone:    "0898123747",
		Password: "hello123",
		Totp:     "398123",
	}

	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("ResetPassword", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		_, err := service.ResetPassword(*mockUser)
		fmt.Print(err)
		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("ResetPassword", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		_, err := service.ResetPassword(*mockUser)
		fmt.Print(err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestSendEmailVerification(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		Name:     "hello",
		Email:    "hello@mail.com",
		Phone:    "0898123747",
		Password: "hello123",
		Totp:     "398123",
	}

	t.Run("Success", func(_ *testing.T) {
		mockUserRepo.On("SendEmailVerification", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		_, err := service.SendEmailVerification(mockUser.Email)
		fmt.Print(err)
	})
	t.Run("Failed", func(_ *testing.T) {
		mockUserRepo.On("SendEmailVerification", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		_, err := service.SendEmailVerification(mockUser.Email)
		fmt.Print(err)
	})
}

func TestRegisterAuth(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		Name:     "hello",
		Email:    "hello@mail.com",
		Phone:    "0898123747",
		Password: "hello123",
		Totp:     "398123",
	}

	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		userService := user.NewUserService(mockUserRepo)
		_, err := userService.CreateUser(*mockUser)
		fmt.Print(err)
		assert.Error(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("SendOTPtoEmail", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected")).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		err := service.SendOTPtoEmail(mockUser.Totp, mockUser.Name, mockUser.Email)
		assert.NoError(t, err)
	})
	t.Run("Err no user exist", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Record not found"))
		service := user.NewUserService(mockUserRepo)
		a, err := service.FindUserByEmail(mockUser.Email)
		assert.Error(t, err)
		assert.Nil(t, a)
	})
	t.Run("Succses", func(t *testing.T) {
		mockUserRepo.On("InsertUser", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		a, err := service.Register(*mockUser)
		assert.NoError(t, err)
		assert.NotNil(t, a)
	})
	t.Run("Failled", func(_ *testing.T) {
		mockUserRepo.On("InsertUser", mock.Anything, mock.Anything).Return(mockUser, errors.New("Unexpected")).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		_, err := service.Register(*mockUser)
		fmt.Print(err)
	})

}

func TestLoginAuth(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &userEntity.User{
		Name:     "hello",
		Email:    "hello@mail.com",
		Phone:    "0898123747",
		Password: "hello123",
	}
	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		userService := user.NewUserService(mockUserRepo)
		_, err := userService.CreateUser(*mockUser)
		fmt.Print(err)
		assert.Error(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(nil, errors.New("Record not found")).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		_, err := service.Login(*mockUser)
		fmt.Print(err)
		assert.Error(t, err)
	})
	t.Run("Filed", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(nil, errors.New("Record not found")).Once()
		jwtService := user.NewJWTService()
		userService := user.NewUserService(mockUserRepo)
		service := user.NewAuthService(userService, jwtService)
		_, err := service.Login(*mockUser)
		fmt.Print(err)
		assert.Error(t, err)
	})
}
