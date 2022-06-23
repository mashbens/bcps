package user

import (
	"log"

	"github.com/mashbens/cps/business/user"
	"github.com/mashbens/cps/business/user/entity"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) user.UserRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) InsertUser(user entity.User) (entity.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	record := fromService(user)
	res := r.db.Create(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil

}

func (r *PostgresRepository) FindByEmail(email string) (entity.User, error) {
	var record User
	res := r.db.Where("email = ?", email).Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil

}
func (r *PostgresRepository) ResetPassword(user entity.User) (entity.User, error) {

	record := fromService(user)
	record.Password = hashAndSalt([]byte(user.Password))
	res := r.db.Model(&user).Where("email = ?", record.Email).Update("password", record.Password)
	if res.Error != nil {
		return record.toService(), res.Error
	}

	res = r.db.Where("email = ?", record.Email).Take(&record)
	return record.toService(), nil
}

func (r *PostgresRepository) FindByUserID(userID string) (entity.User, error) {
	var record User
	res := r.db.Where("id = ?", userID).Take(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
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

func (r *PostgresRepository) UpdateUser(user entity.User) (entity.User, error) {

	record := fromService(user)
	if record.Password != "" {
		record.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempRecord User
		r.db.Find(&tempRecord, user.ID)
		record.Password = tempRecord.Password
	}
	r.db.Save(&record)
	return record.toService(), nil
}

func (r *PostgresRepository) UpdateUserExpiry(userID string, expiry string, memberType string) error {
	var record User
	res := r.db.Model(&record).Where("id = ?", userID).Updates(map[string]interface{}{"member_expired": expiry, "member_type": memberType})
	if res != nil {
		return nil
	}
	return nil
}
