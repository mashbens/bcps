package superadmin

import (
	"github.com/mashbens/cps/business/superadmin"
	"github.com/mashbens/cps/business/superadmin/entity"

	"gorm.io/gorm"
)

type SuperAdminPosgresRepo struct {
	db *gorm.DB
}

func SuperAdminRepo(db *gorm.DB) superadmin.SuperAdminRepo {
	return &SuperAdminPosgresRepo{
		db: db,
	}
}

func (r *SuperAdminPosgresRepo) InsertSuperAdmin(sAdmin entity.SuperAdmin) (entity.SuperAdmin, error) {
	record := fromService(sAdmin)
	res := r.db.Create(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil

}

func (r *SuperAdminPosgresRepo) FindSuperAdminByName(name string) (entity.SuperAdmin, error) {
	var record SuperAdmin
	res := r.db.Where("name = ?", name).Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}
func (r *SuperAdminPosgresRepo) FindSuperAdminByID(id string) (entity.SuperAdmin, error) {
	var record SuperAdmin
	res := r.db.Where("id = ?", id).Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}
