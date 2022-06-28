package entity

type Payment struct {
	ID           int
	UserID       int
	User         User
	MembershipID int
	Membership   Membership
	Amount       int
	SnapURL      string
}

type User struct {
	ID             int
	Name           string
	Email          string
	Phone          string
	Password       string
	Token          string
	Member_expired string
	Member_type    string
	Totp           string
}

type Membership struct {
	ID            int
	Type          string
	Price         int
	Duration      int
	Super_adminID int
}
