package user

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/mashbens/cps/business/user/entity"
	"github.com/mashbens/cps/config"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/xlzd/gotp"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user entity.User) (*entity.User, error)
	Login(user entity.User) (*entity.User, error)
	SendEmailForgotPassword(user entity.User) (*entity.User, error)
	ResetPassword(user entity.User) (*entity.User, error)
	VerifyCredential(email string, password string) error
	GenerateTOTP(email string) string
	SendOTPtoEmail(otp string, name string, email string) error
	SendEmailVerification(email string) (*entity.User, error)
}

type authService struct {
	userService UserService
	jwtService  JWTService
}

func NewAuthService(
	userService UserService,
	jwtService JWTService,
) AuthService {
	return &authService{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *authService) Register(user entity.User) (*entity.User, error) {
	u, err := c.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}
	totp := c.GenerateTOTP(user.Email)
	u.Totp = totp

	sendEmail := c.SendOTPtoEmail(totp, user.Name, user.Email)
	if sendEmail != nil {
		return nil, sendEmail
	}

	token := c.jwtService.GenerateToken((strconv.Itoa(user.ID)))
	u.Token = token

	return u, nil
}

func (c *authService) Login(user entity.User) (*entity.User, error) {
	// verify credential

	err := c.VerifyCredential(user.Email, user.Password)
	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	usr, _ := c.userService.FindUserByEmail(user.Email)

	token := c.jwtService.GenerateToken((strconv.Itoa(usr.ID)))
	usr.Token = token

	return usr, nil
}

func (c *authService) SendEmailForgotPassword(user entity.User) (*entity.User, error) {

	usr, err := c.userService.FindUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	totp := c.GenerateTOTP(usr.Email)
	usr.Totp = totp

	sendEmail := c.SendOTPtoEmail(totp, usr.Name, usr.Email)
	if sendEmail != nil {
		return nil, sendEmail
	}

	return usr, nil
}

func (c *authService) ResetPassword(user entity.User) (*entity.User, error) {
	u, err := c.userService.ResetPassword(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (c *authService) VerifyCredential(email string, password string) error {
	user, err := c.userService.FindUserByEmail(email)
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

func (c *authService) GenerateTOTP(email string) string {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "golang",
		AccountName: email,
		Period:      180,
		SecretSize:  20,
		Digits:      otp.DigitsEight,
	})
	if err != nil {
		log.Println(err)
	}

	totp := gotp.NewDefaultTOTP(key.Secret()).Now()
	fmt.Println("TOTP:", totp)

	return totp
}

func (c *authService) SendOTPtoEmail(otp string, name string, email string) error {

	// config := config.GetConfig()

	// companyEmail := config.Mailjet.Email
	// publicKey := config.Mailjet.PublicKey
	// privateKey := config.Mailjet.PrivateKey

	// mailjetClient := mailjet.NewMailjetClient(publicKey, privateKey)

	// htmlpart := fmt.Sprintf(`<H3>Hai %s, Terimakasih sudah daftarin dirimu di GYM30. Dibawah ini adalah kode verifikasi kamu<H3/> <h1>%s</h1>`, name, otp)
	// messagesInfo := []mailjet.InfoMessagesV31{
	// 	{
	// 		From: &mailjet.RecipientV31{
	// 			Email: companyEmail,
	// 			Name:  "GYM30",
	// 		},
	// 		To: &mailjet.RecipientsV31{
	// 			mailjet.RecipientV31{
	// 				Email: email,
	// 				Name:  name,
	// 			},
	// 		},
	// 		Subject:  "OTP Key GYM30",
	// 		TextPart: otp,
	// 		HTMLPart: htmlpart,
	// 		CustomID: "AppGettingStartedTest",
	// 	},
	// }
	// messages := mailjet.MessagesV31{Info: messagesInfo}
	// res, err := mailjetClient.SendMailV31(&messages)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// log.Println(res)
	return nil
}

func (c *authService) SendEmailVerification(email string) (*entity.User, error) {
	config := config.GetConfig()

	companyEmail := config.Mailjet.Email
	publicKey := config.Mailjet.PublicKey
	privateKey := config.Mailjet.PrivateKey
	otp := c.GenerateTOTP(email)

	mailjetClient := mailjet.NewMailjetClient(publicKey, privateKey)
	htmlpart := fmt.Sprintf(`<H3> Dibawah ini adalah kode verifikasi Forgot Password kamu<H3/><h1>%s</h1>`, otp)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: companyEmail,
				Name:  "GYM30",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  email,
				},
			},
			Subject:  "Forgot Password OTP Key GYM30",
			TextPart: otp,
			HTMLPart: htmlpart,
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(res)

	user, err := c.userService.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	user.Totp = otp

	return user, nil
}
