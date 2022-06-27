package admin

import (
	"github.com/mashbens/cps/business/admin"
	"github.com/mashbens/cps/util"
)

func AdminRepositoryFactory(dbCon *util.DatabaseConnection) admin.AdminRepo {
	var userRepository admin.AdminRepo

	if dbCon.Driver == util.PostgreSQL {
		userRepository = AdminRepo(dbCon.PostgreSQL)
		dbCon.PostgreSQL.AutoMigrate(&Admin{})

	} else {
		panic("Database driver not supported")
	}

	return userRepository
}
