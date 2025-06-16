package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/member"
	"kcers-survey/biz/dal/db/mysql/ent/user"
	"strconv"
)

func GetTokenUser(ctx context.Context, c *app.RequestContext, db *ent.Client) (userOne *ent.User, err error) {
	id, exist := c.Get("userId")
	if exist || id != nil {
		uId, ok := id.(string)
		if ok {
			uid, err := strconv.ParseInt(uId, 10, 64)
			if err != nil {
				return nil, err
			}
			create, err := db.User.Query().Where(user.IDEQ(uid)).First(ctx)
			if err != nil {
				return nil, err
			}
			return create, nil
		}
		return nil, err
	}
	return nil, err
}

func GetTokenMember(ctx context.Context, c *app.RequestContext, db *ent.Client) (userOne *ent.Member, err error) {
	id, exist := c.Get("memberId")
	if exist || id != nil {
		mId, ok := id.(string)
		if ok {
			mid, err := strconv.ParseInt(mId, 10, 64)
			if err != nil {
				return nil, err
			}
			userOne, err = db.Member.Query().Where(member.IDEQ(mid)).First(ctx)
			if err != nil {
				return nil, err
			}
			return userOne, nil
		}
		return nil, err
	}
	return nil, err
}
