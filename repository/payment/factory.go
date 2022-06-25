package payment

import (
	"github.com/mashbens/cps/business/payment"
	"github.com/mashbens/cps/util"
)

func PaymentRepositoryFactory(dbCon *util.DatabaseConnection) payment.PaymentRepository {
	var paymentRepository payment.PaymentRepository

	if dbCon.Driver == util.PostgreSQL {
		paymentRepository = NewPaymentPostgresRepo(dbCon.PostgreSQL)
		dbCon.PostgreSQL.AutoMigrate(&Payment{})

	} else {
		panic("Database driver not supported")
	}

	return paymentRepository
}
