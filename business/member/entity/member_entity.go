package entity

import "mime/multipart"

type Membership struct {
	ID            int
	Type          string
	Price         int
	Duration      int
	Description   string
	Super_adminID int
	Img           string
	Super_admin   SuperAdmin
	ImgBB         *multipart.FileHeader
}

type SuperAdmin struct {
	ID       int
	Name     string
	Password string
	Token    string
}
