package admin

import (
	"github.com/mashbens/cps/business/admin"
	"github.com/mashbens/cps/business/admin/entity"
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
