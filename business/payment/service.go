package payment

import (
	"errors"
	"log"
	"strconv"
	"time"

	memberService "github.com/mashbens/cps/business/member"
	"github.com/mashbens/cps/business/payment/entity"
	userService "github.com/mashbens/cps/business/user"
	"github.com/mashbens/cps/config"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentRepository interface {
	InsertPayment(payment entity.Payment) (entity.Payment, error)
	GetPaymentDetails(userID string) (entity.Payment, error)
	// GetPaymentByID(paymentID int) (entity.Payment, error)
}

type PaymentService interface {
	CreatePayment(payment entity.Payment) (*entity.Payment, error)
	FindPaymentDetails(userID string) (*entity.Payment, error)
	// FindPaymentByID(paymentID int) (*entity.Payment, error)
	PaymentMidtrans(paymentID int, memberType string, amount int) string
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

	// create payment
	payment, err := c.paymentRepo.InsertPayment(paymentReq)
	if err != nil {
		return nil, err
	}

	// add duration
	duration := member.Duration
	userExpiry := time.Now().AddDate(0, duration, 0).Format("2006-01-02")

	// find update user expired
	_ = c.userService.UpdateUserExpiry(paymentReq.UserID, userExpiry, member.Type)
	user, err := c.userService.FindUserByID(paymentReq.UserID)
	if err != nil {
		return nil, err
	}

	// snap midtrans
	snap := c.PaymentMidtrans(payment.ID, member.Type, payment.Amount)
	payment.SnapURL = snap

	membertoPayment := MemberToPayment(*member)
	payment.Membership = membertoPayment

	usertoPayment := UserToPayment(*user)
	payment.User = usertoPayment

	return &payment, nil
}

func (c *paymentService) FindPaymentDetails(userID string) (*entity.Payment, error) {

	payment, err := c.paymentRepo.GetPaymentDetails(userID)
	if err != nil {
		return nil, err
	}
	user, err := c.userService.FindUserByID(payment.UserID)
	if err != nil {
		return nil, err
	}
	strMemberID := strconv.Itoa(payment.MembershipID)

	member, err := c.memberService.FindMemberTypeByID(strMemberID)
	if err != nil {
		return nil, err
	}

	membertoPayment := MemberToPayment(*member)
	payment.Membership = membertoPayment

	usertoPayment := UserToPayment(*user)
	payment.User = usertoPayment

	return &payment, nil

}

func (c *paymentService) PaymentMidtrans(paymentID int, memberType string, amount int) string {
	config := config.GetConfig()

	// var SandboxClientKey = config.Midtrans.ClientKey
	var SandboxServerKey = config.Midtrans.ServerKey

	var s snap.Client
	s.New(SandboxServerKey, midtrans.Sandbox)

	int64Amount := int64(amount)
	strPaymentID := strconv.Itoa(paymentID)
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MEMBER" + "-" + strPaymentID + "-" + memberType + "-" + Random(),
			GrossAmt: int64Amount,
		},
	}
	resp, _ := s.CreateTransaction(req)
	log.Print(resp)
	strRiderectURL := resp.RedirectURL
	return strRiderectURL
}

func Random() string {
	time.Sleep(500 * time.Millisecond)
	return strconv.FormatInt(time.Now().Unix(), 10)
}
