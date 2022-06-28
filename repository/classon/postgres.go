package classon

import (
	"github.com/mashbens/cps/business/classon"
	"github.com/mashbens/cps/business/classon/entity"
	"gorm.io/gorm"
)

type ClassOnPostgresRepo struct {
	db *gorm.DB
}

func NewClassOnPostgresRepo(db *gorm.DB) classon.ClassOnRepo {
	return &ClassOnPostgresRepo{
		db: db,
	}
}

func (c *ClassOnPostgresRepo) FindClassOnByID(classID string) (entity.Classon, error) {
	var record ClassOnline
	res := c.db.Where("id = ?", classID).Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}

func (c *ClassOnPostgresRepo) InserClassOn(class entity.Classon) (entity.Classon, error) {

	record := fromService(class)
	res := c.db.Create(&record)

	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}

func (c *ClassOnPostgresRepo) FindAllClassOn(search string) (data []entity.Classon) {
	var record []ClassOnline
	res := c.db.Find(&record)
	if res.Error != nil {
		return []entity.Classon{}
	}
	return toServiceList(record)
}
func (c *ClassOnPostgresRepo) UpdateClassOn(class entity.Classon) (entity.Classon, error) {
	record := fromService(class)
	res := c.db.Model(&record).Updates(map[string]interface{}{"classname": class.Classname, "trainer": class.Trainer, "date": class.Date, "clock": class.Clock, "description": class.Description})
	if res.Error != nil {
		return record.toService(), res.Error
	}

	return record.toService(), nil
}
func (c *ClassOnPostgresRepo) DeleteClassOn(classID string) error {
	record := []ClassOnline{}
	res := c.db.Delete(&record, classID)
	if res.Error != nil {
		return nil
	}
	return nil
}
