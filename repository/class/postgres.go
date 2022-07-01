package class

import (
	"github.com/mashbens/cps/business/class"
	"github.com/mashbens/cps/business/class/entity"
	"gorm.io/gorm"
)

type ClassPostgresRepo struct {
	db *gorm.DB
}

func NewClassPostgresRepo(db *gorm.DB) class.ClassRepo {
	return &ClassPostgresRepo{
		db: db,
	}
}

func (c *ClassPostgresRepo) InsertClass(class entity.Class) (entity.Class, error) {

	record := fromService(class)
	res := c.db.Create(&record)

	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}

func (c *ClassPostgresRepo) UpdateClass(class entity.Class) (entity.Class, error) {
	record := fromService(class)
	res := c.db.Model(&record).Updates(map[string]interface{}{"classname": class.Classname, "trainer": class.Trainer, "date": class.Date, "clock": class.Clock, "description": class.Description, "class_type": class.ClassType, "capacity": class.Capacity})
	if res.Error != nil {
		return record.toService(), res.Error
	}

	return record.toService(), nil
}

func (c *ClassPostgresRepo) DeleteClass(classID string) error {
	record := []Class{}
	res := c.db.Delete(&record, classID)
	if res.Error != nil {
		return nil
	}
	return nil

}

func (c *ClassPostgresRepo) FindClassByID(classID string) (entity.Class, error) {
	var record Class
	res := c.db.Where("id = ?", classID).Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}

	return record.toService(), nil
}
func (c *ClassPostgresRepo) FindAllClass(classID string) (data []entity.Class) {
	var record []Class
	res := c.db.Find(&record)
	if res.Error != nil {
		return []entity.Class{}
	}
	return toServiceList(record)
}
func (c *ClassPostgresRepo) FindClassOnByID(classID string) (entity.Class, error) {
	var record Class
	res := c.db.Where("id = ? AND class_type = ? ", classID, "online").Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}

func (c *ClassPostgresRepo) FindAllClassOn(search string) (data []entity.Class) {
	var record []Class
	res := c.db.Where("class_type = ?", "online").Take(&record)
	if res.Error != nil {
		return []entity.Class{}
	}
	return toServiceList(record)
}
func (c *ClassPostgresRepo) FindClassOffByID(classID string) (entity.Class, error) {
	var record Class
	res := c.db.Where("id = ? AND class_type = ? ", classID, "offline").Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}

func (c *ClassPostgresRepo) FindAllClassOff(search string) (data []entity.Class) {
	var record []Class
	res := c.db.Where("class_type = ?", "offline").Find(&record)
	if res.Error != nil {
		return []entity.Class{}
	}
	return toServiceList(record)
}

func (c *ClassPostgresRepo) UpdateClassStatus(classID string, status string) error {
	var record Class
	res := c.db.Model(&record).Where("id = ?", classID).Update("status", status)

	if res != nil {
		return nil
	}
	return nil
}

func (c *ClassPostgresRepo) UpdateUserBooked(classID string, userBooked int) error {
	var record Class
	res := c.db.Model(&record).Where("id = ?", classID).Update("user_booked", userBooked)
	if res != nil {
		return nil
	}
	return nil
}
