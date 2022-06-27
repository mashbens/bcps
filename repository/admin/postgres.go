package admin

import (
	"log"

	"github.com/mashbens/cps/business/admin"
	"github.com/mashbens/cps/business/admin/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminPostgresRepo struct {
	db *gorm.DB
}

func AdminRepo(db *gorm.DB) admin.AdminRepo {
	return &AdminPostgresRepo{
		db: db,
	}
}

func (r *AdminPostgresRepo) InsertAdmin(admin entity.Admin) (entity.Admin, error) {
	record := fromService(admin)
	res := r.db.Create(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}
func (r *AdminPostgresRepo) FindAdminByEmail(email string) (entity.Admin, error) {
	var record Admin
	res := r.db.Where("email = ?", email).Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}

func (r *AdminPostgresRepo) FindAdminByID(id string) (entity.Admin, error) {
	var record Admin
	res := r.db.Where("id = ?", id).Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}

func (r *AdminPostgresRepo) FindAllAdmins(search string) (data []entity.Admin) {
	var record []Admin
	res := r.db.Find(&record)
	if res.Error != nil {
		return []entity.Admin{}
	}
	return toServiceList(record)
}

func (c *AdminPostgresRepo) UpdateAdmin(admin entity.Admin) (entity.Admin, error) {
	record := fromService(admin)
	var tempRecord Admin
	c.db.Find(&tempRecord, admin.ID)
	if record.Password != "" {
		record.Password = hashAndSalt([]byte(admin.Password))
	} else {
		c.db.Find(&tempRecord, admin.ID)
		record.Password = tempRecord.Password
	}
	c.db.Save(&record)
	return record.toService(), nil

}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func (c *AdminPostgresRepo) DeleteAdmin(adminId string) error {
	record := []Admin{}
	res := c.db.Delete(&record, adminId)
	if res.Error != nil {
		return nil
	}
	return nil
}
