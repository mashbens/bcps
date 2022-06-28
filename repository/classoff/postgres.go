package classoff

import (
	"github.com/mashbens/cps/business/classoff"
	"github.com/mashbens/cps/business/classoff/entity"
	"gorm.io/gorm"
)

type ClassOffPostgresRepo struct {
	db *gorm.DB
}

func NewClassOffPostgresRepo(db *gorm.DB) classoff.ClassOffRepo {
	return &ClassOffPostgresRepo{
		db: db,
	}
}

func (c *ClassOffPostgresRepo) FindClassOffByID(classID string) (entity.Classoff, error) {
	var record Offlineclass
	res := c.db.Where("id = ?", classID).Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}

func (c *ClassOffPostgresRepo) InserClassOff(class entity.Classoff) (entity.Classoff, error) {

	record := fromService(class)
	res := c.db.Create(&record)

	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}

func (c *ClassOffPostgresRepo) FindAllClassOff(search string) (data []entity.Classoff) {
	var record []Offlineclass
	res := c.db.Find(&record)
	if res.Error != nil {
		return []entity.Classoff{}
	}
	return toServiceList(record)
}
func (c *ClassOffPostgresRepo) UpdateClassOff(class entity.Classoff) (entity.Classoff, error) {
	record := fromService(class)
	res := c.db.Model(&record).Updates(map[string]interface{}{"classname": class.Classname, "trainer": class.Trainer, "date": class.Date, "clock": class.Clock, "description": class.Description})
	if res.Error != nil {
		return record.toService(), res.Error
	}

	return record.toService(), nil
}
func (c *ClassOffPostgresRepo) DeleteClassOff(classID string) error {
	record := []Offlineclass{}
	res := c.db.Delete(&record, classID)
	if res.Error != nil {
		return nil
	}
	return nil
}
