package payment

import (
	"github.com/mashbens/cps/business/payment"
	"github.com/mashbens/cps/business/payment/entity"
	"gorm.io/gorm"
)

type PaymentPostgresRepo struct {
	db *gorm.DB
}

func NewPaymentPostgresRepo(db *gorm.DB) payment.PaymentRepository {
	return &PaymentPostgresRepo{
		db: db,
	}
}

func (c *PaymentPostgresRepo) InsertPayment(payment entity.Payment) (entity.Payment, error) {
	record := fromService(payment)
	res := c.db.Create(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}
