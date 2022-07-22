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

func (r *PostgresRepository) InsertUser(user entity.User) (*entity.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	record := fromService(user)
	res := r.db.Create(&record)
	if res.Error != nil {
		return nil, res.Error
	}
	data := record.toService()
	return &data, nil

}

func (r *PostgresRepository) FindByEmail(email string) (*entity.User, error) {
	var record User
	res := r.db.Where("email = ?", email).Take(&record)
	if res.Error != nil {
		return nil, res.Error
	}
	data := record.toService()
	return &data, nil

}
func (r *PostgresRepository) ResetPassword(user entity.User) (*entity.User, error) {

	record := fromService(user)
	record.Password = hashAndSalt([]byte(user.Password))
	res := r.db.Model(&user).Where("email = ?", record.Email).Update("password", record.Password)
	if res.Error != nil {
		return nil, res.Error
	}

	res = r.db.Where("email = ?", record.Email).Take(&record)
	data := record.toService()
	return &data, nil
}

func (r *PostgresRepository) FindByUserID(userID int) (*entity.User, error) {
	var record User
	res := r.db.Where("id = ?", userID).Take(&record)
	if res.Error != nil {
		return nil, res.Error
	}
	data := record.toService()
	return &data, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func (r *PostgresRepository) UpdateUser(user entity.User) (*entity.User, error) {

	record := fromService(user)

	var tempRecord User
	r.db.Find(&tempRecord, user.ID)

	if record.Password != "" {
		record.Password = hashAndSalt([]byte(user.Password))
	} else {
		r.db.Find(&tempRecord, user.ID)
		record.Password = tempRecord.Password
	}

	if record.Email != "" {
		record.Email = user.Email
	} else {
		record.Email = tempRecord.Email
	}

	if record.Phone != "" {
		record.Phone = user.Phone
	} else {
		record.Phone = tempRecord.Phone
	}

	if record.Phone == tempRecord.Phone {
		record.Phone = tempRecord.Phone
	}

	if record.Name != "" {
		record.Name = user.Name
	} else {
		record.Name = tempRecord.Name
	}

	if record.Member_expired != "" {
		record.Member_expired = user.Member_expired
	} else {
		record.Member_expired = tempRecord.Member_expired
	}

	if record.Member_type != "" {
		record.Member_type = user.Member_type
	} else {
		record.Member_type = tempRecord.Member_type
	}

	r.db.Save(&record)
	data := record.toService()
	return &data, nil
}

func (r *PostgresRepository) UpdateUserExpiry(userID int, expiry string, memberType string) error {
	var record User
	res := r.db.Model(&record).Where("id = ?", userID).Updates(map[string]interface{}{"member_expired": expiry, "member_type": memberType})
	if res != nil {
		return nil
	}
	return nil
}
