package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/user"
	"strconv"
)

func GetTokenUserID(c *app.RequestContext) int64 {
	id, exist := c.Get("userId")
	if exist || id != nil {
		uId, ok := id.(string)
		if ok {
			uid, err := strconv.ParseInt(uId, 10, 64)
			if err != nil {
				return 0
			}
			return uid
		}
		return 0
	}
	return 0
}

func GetTokenUser(ctx context.Context, c *app.RequestContext, db *ent.Client) (one *ent.User, err error) {
	id := GetTokenUserID(c)
	if id >= 0 {
		one, err = GetUser(db, id)
		if err != nil {
			return nil, err
		}
		return one, nil
	}
	return nil, err
}
func GetTokenUserName(ctx context.Context, c *app.RequestContext, db *ent.Client) string {
	userEnt, _ := GetTokenUser(ctx, c, db)
	if userEnt != nil {
		return userEnt.Name
	}
	return ""
}

func GetUser(db *ent.Client, id int64) (one *ent.User, err error) {
	one, err = db.User.Query().Where(user.ID(id)).First(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "未查询到该账号")
	}
	return one, err
}
