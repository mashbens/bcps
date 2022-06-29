package class

import (
	"github.com/mashbens/cps/business/class"
	"github.com/mashbens/cps/util"
)

func ClassRepoFactory(dbCon *util.DatabaseConnection) class.ClassRepo {
	var memberRepository class.ClassRepo

	if dbCon.Driver == util.PostgreSQL {
		memberRepository = NewClassPostgresRepo(dbCon.PostgreSQL)
		dbCon.PostgreSQL.AutoMigrate(&Class{})

	} else {
		panic("Database driver not supported")
	}

	return memberRepository
}
