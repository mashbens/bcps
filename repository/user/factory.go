package user

import (
	"github.com/mashbens/cps/business/user"
	util "github.com/mashbens/cps/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) user.UserRepository {
	var userRepository user.UserRepository

	if dbCon.Driver == util.PostgreSQL {
		userRepository = NewPostgresRepository(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}

	return userRepository
}
