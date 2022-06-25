package superadmin

import (
	"github.com/mashbens/cps/business/superadmin"
	"github.com/mashbens/cps/util"
)

func SuperAdminRepositoryFactory(dbCon *util.DatabaseConnection) superadmin.SuperAdminRepo {
	var userRepository superadmin.SuperAdminRepo

	if dbCon.Driver == util.PostgreSQL {
		userRepository = SuperAdminRepo(dbCon.PostgreSQL)
		dbCon.PostgreSQL.AutoMigrate(&SuperAdmin{})

	} else {
		panic("Database driver not supported")
	}

	return userRepository
}
