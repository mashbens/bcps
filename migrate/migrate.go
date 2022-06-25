package migrate

import (
	// member "github.com/mashbens/cps/repository/member"
	// payment "github.com/mashbens/cps/repository/payment"
	// superAdmin "github.com/mashbens/cps/repository/superadmin"
	user "github.com/mashbens/cps/repository/user"
	"gorm.io/gorm"
)

// type User struct {
// 	ID             int    `gorm:"primary_key:auto_increment" json:"-"`
// 	Name           string `gorm:"type:varchar(100)" json:"-"`
// 	Email          string `gorm:"type:varchar(100);unique;" json:"-"`
// 	Phone          string `gorm:"type:varchar(100);unique;" json:"-"`
// 	Password       string `gorm:"type:varchar(100)" json:"-"`
// 	Member_expired string `gorm:"type:varchar(100)" json:"-"`
// 	Member_type    string `gorm:"type:varchar(100)" json:"-"`
// 	*gorm.Model
// }

// type Membership struct {
// 	ID           int    `gorm:"primary_key:auto_increment" json:"-"`
// 	Type         string `gorm:"type:varchar(100)" `
// 	Price        int
// 	Duration     int
// 	SuperAdminID int
// 	Super_admin  SuperAdmin `gorm:"foreignkey:SuperAdminID" json:"-"`
// }

// type Payment struct {
// 	ID           int `gorm:"primary_key:auto_increment" json:"-"`
// 	UserID       int
// 	User         User `gorm:"foreignkey:UserID" json:"-"`
// 	MembershipID int
// 	Membership   Membership `gorm:"foreignkey:MembershipID" json:"-"`
// 	Amount       int
// }

// type SuperAdmin struct {
// 	ID       int    `gorm:"primary_key:auto_increment" json:"-"`
// 	Name     string `gorm:"type:varchar(100)" json:"-"`
// 	Password string `gorm:"type:varchar(100)" json:"-"`
// }

func NewMigrate(db *gorm.DB) {

	db.AutoMigrate(
		&user.User{},
		// &member.Membership{},
		// &superAdmin.SuperAdmin{},
		// &payment.Payment{},
	)
}
