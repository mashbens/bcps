package newsletter

import (
	"github.com/mashbens/cps/business/newsletter/entity"
)

type News struct {
	ID      int    `gorm:"primary_key:auto_increment" json:"-"`
	Title   string `gorm:"type:varchar(100)" `
	Date    string `gorm:"type:varchar(100)" `
	Content string `gorm:"type:text" `
	Img     string `gorm:"type:varchar(100)" `
}

func (n *News) toService() entity.News {
	return entity.News{
		ID:      n.ID,
		Title:   n.Title,
		Date:    n.Date,
		Content: n.Content,
		Img:     n.Img,
	}
}

func fromService(n entity.News) News {
	return News{
		ID:      n.ID,
		Title:   n.Title,
		Date:    n.Date,
		Content: n.Content,
		Img:     n.Img,
	}
}

func toServiceList(data []News) []entity.News {
	a := []entity.News{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
