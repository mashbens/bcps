package classoff

import (
	"github.com/mashbens/cps/business/classoff"
	"github.com/mashbens/cps/util"
)

func ClassOffRepoFactory(dbCon *util.DatabaseConnection) classoff.ClassOffRepo {
	var classOffRepository classoff.ClassOffRepo

	if dbCon.Driver == util.PostgreSQL {
		classOffRepository = NewClassOffPostgresRepo(dbCon.PostgreSQL)
		dbCon.PostgreSQL.AutoMigrate(&Offlineclass{})
	} else {
		panic("Database driver not supported")
	}

	return classOffRepository
}
