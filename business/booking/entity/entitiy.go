package entity

type Booking struct {
	ID         int
	UserID     int
	User       User
	ClassID    int
	Class      Class
	ClassSlice []Class
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
}

type Class struct {
	ID          int
	Classname   string
	Trainer     string
	Date        string
	Clock       string
	Description string
	ClassType   string
	Capacity    int
	UserBooked  int
	Status      string
}
