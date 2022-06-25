package util

import (
	"fmt"

	"github.com/mashbens/cps/config"
	// "github.com/mashbens/cps/migrate"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	MySQL DatabaseDriver = "mysql"

	PostgreSQL DatabaseDriver = "postgres"

	Static DatabaseDriver = "static"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	MySQL *gorm.DB

	PostgreSQL *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch config.Database.Driver {
	case "MySQL":
		db.Driver = MySQL
		db.MySQL = newMySQL(config)
	case "PostgreSQL":
		db.Driver = PostgreSQL
		db.PostgreSQL = NewPostgreSQL(config)
	default:
		panic("Database driver not supported")
	}
	return &db
}
func NewPostgreSQL(config *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Database.DB_Host,
		config.Database.DB_User,
		config.Database.DB_Pass,
		config.Database.DB_Name,
		config.Database.DB_Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&migrate.User{}, &migrate.Membership{}, &migrate.Payment{}, &migrate.SuperAdmin{})

	// migrate.NewMigrate(db)
	return db
}
func newMySQL(config *config.AppConfig) *gorm.DB {
	return nil
}

func (db *DatabaseConnection) CloseConnection() {
	if db.MySQL != nil {
		db, _ := db.MySQL.DB()
		db.Close()
	}
}
