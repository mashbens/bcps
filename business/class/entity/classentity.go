package entity

type Class struct {
	ID          int
	Classname   string
	Trainer     string
	Date        string
	Clock       string
	Description string
	ClassType   string
	AdminID     int
	Admin       Admin
}

type Admin struct {
	ID       int
	Name     string
	Password string
	Email    string
	Phone    string
}
