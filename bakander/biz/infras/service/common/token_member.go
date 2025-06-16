package common

import (
	"fmt"
	"github.com/pkg/errors"
	"kcers-survey/biz/dal/db/mysql/ent"
	member2 "kcers-survey/biz/dal/db/mysql/ent/member"
	"kcers-survey/biz/dal/db/mysql/ent/membertoken"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	"kcers-survey/idl_gen/model/token"
	"time"
)

func (t *Token) CreateMemberToken(req *token.MemberTokenInfo) error {
	expiredAt, _ := time.ParseInLocation(time.DateTime, req.ExpiredAt, time.Local)
	if expiredAt.Sub(time.Now()).Seconds() < 5 {
		return errors.New("expired time must be greater than now, more than 5s")
	}
	tokenExist, err := t.db.MemberToken.Query().Where(membertoken.MemberID(req.MemberId)).Only(t.ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return errors.Wrap(err, "create Token failed")
		}
	}
	if tokenExist != nil {
		req.ID = tokenExist.ID
		return t.UpdateMemberToken(req)
	}
	memberInfo, err := t.db.Member.Query().Where(member2.ID(req.MemberId)).Only(t.ctx)
	if err != nil {
		return err
	}
	_, err = t.db.MemberToken.Create().
		SetOwner(memberInfo).
		SetMemberID(req.MemberId).
		SetToken(req.Token).
		SetSource(req.Source).
		SetExpiredAt(expiredAt).
		Save(t.ctx)

	return nil
}

func (t *Token) UpdateMemberToken(req *token.MemberTokenInfo) error {
	expiredAt, _ := time.ParseInLocation(time.DateTime, req.ExpiredAt, time.Local)
	if expiredAt.Sub(time.Now()).Seconds() < 5 {
		return errors.New("expired time must be greater than now, more than 5s")
	}
	_, err := t.db.MemberToken.UpdateOneID(req.ID).
		SetMemberID(req.MemberId).
		SetToken(req.Token).
		SetSource(req.Source).
		SetUpdatedAt(time.Now()).
		SetExpiredAt(expiredAt).
		Save(t.ctx)
	if err != nil {
		return errors.Wrap(err, "update Token failed")
	}

	t.cache.SetWithTTL(fmt.Sprintf("member_token_%d", req.MemberId), req.MemberId, 1, expiredAt.Sub(time.Now()))
	return nil
}

func (t *Token) IsExistByMemberIdToken(memberId int64) bool {
	_, exist := t.cache.Get(fmt.Sprintf("member_token_%d", memberId))
	if exist {
		return true
	}
	exist, _ = t.db.MemberToken.Query().Where(membertoken.MemberIDEQ(memberId)).Exist(t.ctx)
	return exist
}

func (t *Token) DeleteMemberToken(MemberId int64) error {
	_, err := t.db.MemberToken.Delete().Where(membertoken.MemberID(MemberId)).Exec(t.ctx)
	if err != nil {
		return errors.Wrap(err, "delete Token failed")
	}
	t.cache.Del(fmt.Sprintf("member_token_%d", MemberId))
	return nil
}
func (t *Token) MemberTokenList(req *token.TokenListReq) (res []*token.MemberTokenInfo, total int, err error) {

	var predicates = []predicate.Member{member2.HasToken()}
	if req.Username != "" {
		predicates = append(predicates, member2.UsernameContainsFold(req.Username))
	}
	if req.UserId != 0 {
		predicates = append(predicates, member2.IDEQ(req.UserId))
	}
	predicates = append(predicates, member2.Delete(0))
	MemberTokens, err := t.db.Member.Query().Where(predicates...).
		WithToken(func(q *ent.MemberTokenQuery) {
			// get token all fields default, or use q.Select() to get some fields
		}).Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(member2.FieldID)).
		Limit(int(req.PageSize)).All(t.ctx)
	if err != nil {
		return res, total, errors.Wrap(err, "get User - Token list failed")
	}

	for _, memberEnt := range MemberTokens {
		res = append(res, &token.MemberTokenInfo{
			ID:        memberEnt.Edges.Token.ID,
			MemberId:  memberEnt.ID,
			Name:      memberEnt.Username,
			Token:     memberEnt.Edges.Token.Token,
			Source:    memberEnt.Edges.Token.Source,
			CreatedAt: memberEnt.CreatedAt.Format(time.DateTime),
			UpdatedAt: memberEnt.UpdatedAt.Format(time.DateTime),
			ExpiredAt: memberEnt.Edges.Token.ExpiredAt.Format(time.DateTime),
		})

		// delete expired token from db
		if memberEnt.Edges.Token.ExpiredAt.Sub(time.Now()).Seconds() < 0 {
			_ = t.Delete(memberEnt.ID)
		}
	}
	total, _ = t.db.Member.Query().Where(predicates...).Count(t.ctx)
	return
}
