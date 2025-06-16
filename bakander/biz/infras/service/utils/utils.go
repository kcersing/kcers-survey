package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"

	"strconv"
)

func GetTokenUser(ctx context.Context, c *app.RequestContext, db *ent.Client) (one *ent.User, err error) {
	id, exist := c.Get("userId")
	hlog.Info("userId:", id)
	if exist || id != nil {
		uId, ok := id.(string)
		if ok {
			uid, err := strconv.ParseInt(uId, 10, 64)
			if err != nil {
				return nil, err
			}
			one, err = GetUser(db, uid)
			if err != nil {
				return nil, err
			}
			return one, nil
		}
		return nil, err
	}
	return nil, err
}

func GetTokenMember(ctx context.Context, c *app.RequestContext, db *ent.Client) (one *ent.Member, err error) {
	id, exist := c.Get("memberId")
	if exist || id != nil {
		mId, ok := id.(string)
		if ok {
			mid, err := strconv.ParseInt(mId, 10, 64)
			if err != nil {
				return nil, err
			}
			one, err = GetMember(db, mid)
			if err != nil {
				return nil, err
			}
			return one, nil
		}
		return nil, err
	}
	return nil, err
}

func GetMember(db *ent.Client, id int64) (one *ent.Member, err error) {
	one, err = db.Member.Query().Where(member.IDEQ(id)).First(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "未查询到该会员")
	}
	return one, nil

}
func GetMemberProduct(db *ent.Client, id int64) (one *ent.MemberProduct, err error) {
	one, err = db.MemberProduct.Query().Where(memberproduct.ID(id)).First(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "未查询到该会员产品")
	}
	return one, err
}

func GetUser(db *ent.Client, id int64) (one *ent.User, err error) {
	one, err = db.User.Query().Where(user.ID(id)).First(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "未查询到该员工")
	}
	return one, err
}

func GetVenue(db *ent.Client, id int64) (one *ent.Venue, err error) {
	one, err = db.Venue.Query().Where(venue.ID(id)).First(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "未查询到该场馆")
	}
	return one, err
}

func GetProduct(db *ent.Client, id int64) (one *ent.Product, err error) {
	one, err = db.Product.Query().Where(product2.ID(id)).First(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "未查询到该产品")
	}
	return one, err
}
func DeductNumberSurplus(db *ent.Client, id int64) (err error) {
	_, err = db.MemberProduct.UpdateOneID(id).
		AddNumberSurplus(-1).
		Save(context.Background())
	return err
}
