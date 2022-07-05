package request

import (
	"mime/multipart"

	"github.com/mashbens/cps/business/newsletter/entity"
)

type CreatNewsReq struct {
	ID      int    `json:"id" form:"id" param:"id"`
	Title   string `json:"title" form:"title"`
	Date    string `json:"date" form:"date"`
	Content string `json:"content" form:"content"`
	Image   string `json:"image" form:"image"`
	AdminID int    `json:"admin_id" form:"admin_id"`
	ImgBB   *multipart.FileHeader
}

func NewCreateNewsReq(req CreatNewsReq) entity.News {
	return entity.News{
		ID:      req.ID,
		Date:    req.Date,
		Title:   req.Title,
		Content: req.Content,
		AdminID: req.AdminID,
		ImgBB:   req.ImgBB,
	}
}
