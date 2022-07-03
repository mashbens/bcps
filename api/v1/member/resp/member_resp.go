package resp

import (
	member "github.com/mashbens/cps/business/member/entity"
)

type MemberResp struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Price       int    `json:"price"`
	Duration    int    `json:"duration"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

func FromService(member member.Membership) MemberResp {
	return MemberResp{
		ID:          member.ID,
		Type:        member.Type,
		Price:       member.Price,
		Duration:    member.Duration,
		Description: member.Description,
		Img:         member.Img,
	}
}

func FromServiceSlice(data []member.Membership) []MemberResp {
	var memberAray []MemberResp
	for key := range data {
		memberAray = append(memberAray, FromService(data[key]))

	}
	return memberAray
}
