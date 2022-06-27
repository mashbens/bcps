package entity

import "github.com/mashbens/cps/business/superadmin/entity"

type Admin struct {
	ID           int
	Name         string
	Password     string
	Email        string
	Phone        string
	SuperAdminID int
	SuperAdmin   entity.SuperAdmin
	Token        string
}
