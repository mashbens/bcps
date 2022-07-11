package member_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/mashbens/cps/business/member"
	"github.com/mashbens/cps/business/member/entity"
	memberEntity "github.com/mashbens/cps/business/member/entity"
)

var service member.MemberService
var member1, member2 memberEntity.Membership

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindMemberByID(t *testing.T) {
	t.Run("Expect to find member by id", func(t *testing.T) {
		memberID := int(member1.ID)
		memberIDs := strconv.Itoa(memberID)
		member, err := service.FindMemberTypeByID(memberIDs)
		if err != nil {
			t.Error("Expect to find member by id", err)
		} else {
			if member.ID != 0 {
				t.Errorf("Expected %d, got %d", 0, member.ID)
			}
		}
	})
	t.Run("Expect not found the content", func(t *testing.T) {
		memberID := int(member1.ID)
		memberIDs := strconv.Itoa(memberID)
		member, err := service.FindMemberTypeByID(memberIDs)

		if err != nil {
			t.Error("Expect error is nil. Error: ", err)
		} else if member == nil {
			t.Error("Expect to find user by id", err)

		}
	})
}
func TestFindMemberAll(t *testing.T) {
	t.Run("Expect to find member by email", func(t *testing.T) {
		err := service.FindAllMemberType("")
		if err == nil {
			t.Error("Expext error is invalid email or password. Error: ", err)
			t.FailNow()
		}
	})
}

func TestDeletMember(t *testing.T) {
	t.Run("Expect to find member by email", func(t *testing.T) {
		memberID := int(member1.ID)
		memberIDs := strconv.Itoa(memberID)
		err := service.DeleteMemberType(memberIDs, memberIDs)
		if err != nil {
			t.Error("Expext error is invalid email or password. Error: ", err)
			t.FailNow()
		}
	})
}

func setup() {
	member1.ID = 1
	member1.Type = "John"
	member1.Duration = 6
	member1.Description = "test123"

	member2.ID = 2
	member2.Type = "John"
	member2.Duration = 12
	member2.Description = "test123"

	repo := newInMemoryRepository()

	service = member.NewMemberService(repo)
}

type inMemoryRepository struct {
	memberByID    map[string]memberEntity.Membership
	memberByTitle map[string]memberEntity.Membership
}

func newInMemoryRepository() *inMemoryRepository {
	var repo inMemoryRepository
	repo.memberByID = make(map[string]memberEntity.Membership)
	repo.memberByTitle = make(map[string]memberEntity.Membership)

	userID := int64(member1.ID)
	userIDs := strconv.FormatInt(userID, 10)
	repo.memberByID[userIDs] = member1
	repo.memberByTitle[member1.Type] = member1

	return &repo
}
func (r *inMemoryRepository) InserMemberships(member entity.Membership) (entity.Membership, error) {
	return memberEntity.Membership{}, nil
}

func (r *inMemoryRepository) FindMemberByID(memberID string) (entity.Membership, error) {
	return r.memberByTitle[memberID], nil
}

func (r *inMemoryRepository) FindAllMemberType(title string) (data []entity.Membership) {
	return []memberEntity.Membership{}
}

func (r *inMemoryRepository) UpdateMemberType(member entity.Membership) (entity.Membership, error) {
	return member, nil
}
func (r *inMemoryRepository) DeleteMemberType(memberID string) error {
	return nil
}
