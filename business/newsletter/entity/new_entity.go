package entity

import "mime/multipart"

type News struct {
	ID      int
	Title   string
	Date    string
	Content string
	Img     string
	AdminID int

	ImgBB *multipart.FileHeader
}
