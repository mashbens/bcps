package entity

import (
	superAdmin "github.com/mashbens/cps/business/superadmin/entity"
)

type Membership struct {
	ID            int
	Type          string
	Price         int
	Duration      int
	Super_adminID int
	Super_admin   superAdmin.SuperAdmin
}
