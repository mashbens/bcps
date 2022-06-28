package classon

import (
	"github.com/mashbens/cps/business/classon"
	"github.com/mashbens/cps/util"
)

func ClassOnRepoFactory(dbCon *util.DatabaseConnection) classon.ClassOnRepo {
	var classOnRepository classon.ClassOnRepo

	if dbCon.Driver == util.PostgreSQL {
		classOnRepository = NewClassOnPostgresRepo(dbCon.PostgreSQL)
		dbCon.PostgreSQL.AutoMigrate(&Onlineclass{})

	} else {
		panic("Database driver not supported")
	}

	return classOnRepository
}
