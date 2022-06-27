package entity

type Membership struct {
	ID            int
	Type          string
	Price         int
	Duration      int
	Description   string
	Super_adminID int
	Super_admin   SuperAdmin
}

type SuperAdmin struct {
	ID       int
	Name     string
	Password string
	Token    string
}
