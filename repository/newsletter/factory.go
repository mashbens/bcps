package newsletter

import (
	"github.com/mashbens/cps/business/newsletter"
	"github.com/mashbens/cps/util"
)

func NewsRepoFactory(dbCon *util.DatabaseConnection) newsletter.NewsRepo {
	var newsRepo newsletter.NewsRepo

	if dbCon.Driver == util.PostgreSQL {
		newsRepo = NewNewsPosgresRepo(dbCon.PostgreSQL)
		dbCon.PostgreSQL.AutoMigrate(&News{})

	} else {
		panic("Database driver not supported")
	}

	return newsRepo
}
