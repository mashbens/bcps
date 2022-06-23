package payment

import (
	"errors"
	"log"
	"strconv"
	"time"

	memberService "github.com/mashbens/cps/business/member"
	"github.com/mashbens/cps/business/payment/entity"
	userService "github.com/mashbens/cps/business/user"
)

type PaymentRepository interface {
	InsertPayment(payment entity.Payment) (entity.Payment, error)
	// GetPaymentDetails(userID string) (entity.Payment, error)
	// GetPaymentByID(paymentID int) (entity.Payment, error)
}

type PaymentService interface {
	CreatePayment(payment entity.Payment) (*entity.Payment, error)
	// FindPaymentDetails(userID string) (*entity.Payment, error)
	// FindPaymentByID(paymentID int) (*entity.Payment, error)
	// PaymentMidtrans(paymentID int, memberType string, amount int) (string, error)
}

type paymentService struct {
	paymentRepo   PaymentRepository
	memberService memberService.MemberService
	userService   userService.UserService
}

func NewPaymentService(
	paymentRepo PaymentRepository,
	memberService memberService.MemberService,
	userService userService.UserService,
) PaymentService {
	return &paymentService{
		paymentRepo:   paymentRepo,
		memberService: memberService,
		userService:   userService,
	}
}

func (c *paymentService) CreatePayment(paymentReq entity.Payment) (*entity.Payment, error) {
	// find Member type
	strMemberID := strconv.Itoa(paymentReq.MembershipID)

	member, err := c.memberService.FindMemberTypeByID(strMemberID)
	if err != nil {
		return nil, errors.New("member type not found")
	}

	paymentReq.Amount = member.Price

	log.Println(paymentReq.ID)
	log.Println(paymentReq.UserID)

	// create payment
	payment, err := c.paymentRepo.InsertPayment(paymentReq)
	if err != nil {
		return nil, err
	}

	// add duration
	duration := member.Duration
	userExpiry := time.Now().AddDate(0, duration, 0).Format("2006-01-02")

	// update user expired
	strUserID := strconv.Itoa(paymentReq.UserID)
	user := c.userService.UpdateUserExpiry(strUserID, userExpiry, member.Type)
	_ = user

	log.Println("userid  ->", strUserID)
	log.Println("member type ->", member.Type)
	log.Println("user exx ->", userExpiry)

	// snap midtrans
	// snap, _ := c.PaymentMidtrans(payment.ID, member.Type, payment.Amount)

	return &payment, nil

}
