package resp

import (
	"github.com/mashbens/cps/business/newsletter/entity"
)

type NewsResp struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func FromService(news entity.News) NewsResp {
	return NewsResp{
		ID:      news.ID,
		Title:   news.Title,
		Date:    news.Date,
		Content: news.Content,
		Image:   news.Img,
	}
}

func FromServiceSlice(data []entity.News) []NewsResp {
	var memberAray []NewsResp
	for key := range data {
		memberAray = append(memberAray, FromService(data[key]))

	}
	return memberAray
}
