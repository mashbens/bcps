package newsletter

import (
	"github.com/mashbens/cps/business/newsletter"
	"github.com/mashbens/cps/business/newsletter/entity"
	"gorm.io/gorm"
)

type NewsPosgresRepo struct {
	db *gorm.DB
}

func NewNewsPosgresRepo(db *gorm.DB) newsletter.NewsRepo {
	return &NewsPosgresRepo{
		db: db,
	}
}

func (c *NewsPosgresRepo) InsertNews(news entity.News) (entity.News, error) {
	record := fromService(news)
	res := c.db.Create(&record)

	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}
func (c *NewsPosgresRepo) UpdateNews(news entity.News) (entity.News, error) {
	record := fromService(news)
	res := c.db.Model(&record).Updates(map[string]interface{}{"title": news.Title, "content": news.Content, "date": news.Date, "img": news.Img})
	if res.Error != nil {
		return record.toService(), res.Error
	}

	return record.toService(), nil
}

func (c *NewsPosgresRepo) FindNewsByID(newsID string) (entity.News, error) {
	var record News
	res := c.db.Where("id = ?", newsID).Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}

	return record.toService(), nil
}

func (c *NewsPosgresRepo) FindAllNews(search string) (data []entity.News) {
	var record []News
	res := c.db.Find(&record)
	if res.Error != nil {
		return []entity.News{}
	}
	return toServiceList(record)
}
func (c *NewsPosgresRepo) DeleteNews(newsID string) error {
	record := []News{}
	res := c.db.Delete(&record, newsID)
	if res.Error != nil {
		return nil
	}
	return nil
}
