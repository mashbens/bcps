package member

import (
	"github.com/mashbens/cps/business/member"
	"github.com/mashbens/cps/util"
)

func MemberRepoFactory(dbCon *util.DatabaseConnection) member.MemberRepo {
	var memberRepository member.MemberRepo

	if dbCon.Driver == util.PostgreSQL {
		memberRepository = NewMemberPostgresRepository(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}
	return memberRepository
}
