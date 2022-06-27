package entity

type Admin struct {
	ID           int
	Name         string
	Password     string
	Email        string
	Phone        string
	SuperAdminID int
	SuperAdmin   SuperAdmin
	Token        string
}

type SuperAdmin struct {
	ID       int
	Name     string
	Password string
	Token    string
}
