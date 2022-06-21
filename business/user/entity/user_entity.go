package entity

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
