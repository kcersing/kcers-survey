package common

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"kcers-survey/biz/dal/cache"
	"kcers-survey/biz/dal/config"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/contract"
	"kcers-survey/biz/dal/db/mysql/ent/dictionary"
	"kcers-survey/biz/dal/db/mysql/ent/dictionarydetail"
	"kcers-survey/biz/dal/db/mysql/ent/member"
	"kcers-survey/biz/dal/db/mysql/ent/memberproductproperty"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	"kcers-survey/biz/dal/db/mysql/ent/product"
	"kcers-survey/biz/dal/db/mysql/ent/productproperty"
	"kcers-survey/biz/dal/db/mysql/ent/user"
	venue2 "kcers-survey/biz/dal/db/mysql/ent/venue"
	"kcers-survey/biz/dal/db/mysql/ent/venueplace"
	"kcers-survey/biz/infras/do"

	"kcers-survey/idl_gen/model/sys"
	"strconv"
)

type Sys struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (s *Sys) RoleList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error) {
	var predicates []predicate.Role

	lists, err := s.db.Role.Query().Where(predicates...).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Role list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, &sys.SysList{
			ID:   v.ID,
			Name: v.Remark,
		})
	}
	total = int64(len(list))

	return
}

func (s *Sys) PlaceList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error) {
	var predicates []predicate.VenuePlace

	if req.VenueId > 0 {
		predicates = append(predicates, venueplace.VenueID(req.VenueId))
	}

	lists, err := s.db.VenuePlace.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get VenuePlace list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, &sys.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))

	return
}

func (s *Sys) ProductList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error) {
	var predicates []predicate.Product

	if req.Name != "" {
		predicates = append(predicates, product.NameEQ(req.Name))
	}
	lists, err := s.db.Product.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, &sys.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))
	return
}

func (s *Sys) PropertyList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error) {
	var predicates []predicate.ProductProperty

	if req.Name != "" {
		predicates = append(predicates, productproperty.NameEQ(req.Name))
	}
	if req.Type != "" {
		predicates = append(predicates, productproperty.TypeEQ(req.Type))
	}
	if req.ProductId > 0 {
		predicates = append(predicates, productproperty.HasProductWith(product.IDEQ(req.ProductId)))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, productproperty.HasVenuesWith(venue2.IDEQ(req.VenueId)))
	}

	lists, err := s.db.ProductProperty.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		price := strconv.FormatFloat(v.Price, 'f', -1, 64)
		list = append(list, &sys.SysList{
			ID:   v.ID,
			Name: v.Name,
			Key:  price,
		})
	}
	total = int64(len(list))
	return
}

func (s *Sys) PropertyType(req *sys.SysListReq) (list []*sys.SysList, total int64, err error) {
	var predicates []predicate.DictionaryDetail

	if req.Name != "" {
		predicates = append(predicates, dictionarydetail.ValueEQ(req.Name))
	}
	if req.DictionaryId != 0 {
		predicates = append(predicates, dictionarydetail.HasDictionaryWith(dictionary.IDEQ(req.DictionaryId)))
	}

	lists, err := s.db.DictionaryDetail.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, &sys.SysList{
			ID:   v.ID,
			Name: v.Value,
			Key:  v.Key,
		})
	}
	total = int64(len(list))
	return
}

func (s *Sys) VenueList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error) {
	var predicates []predicate.Venue

	if req.Name != "" {
		predicates = append(predicates, venue2.Name(req.Name))
	}

	lists, err := s.db.Venue.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, &sys.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))

	return
}

func (s *Sys) MemberList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error) {
	var predicates []predicate.Member

	if req.Name != "" {
		predicates = append(predicates, func(s *sql.Selector) {
			s.Where(sql.Like(member.FieldName, "%"+req.Name+"%"))
		})
	}
	if req.Mobile != "" {
		predicates = append(predicates, member.Mobile(req.Mobile))
	}

	if req.ProductId > 0 {
		ints, err := s.db.Debug().MemberProductProperty.Query().Where(memberproductproperty.MemberProductID(req.ProductId)).GroupBy(member.FieldID).Ints(s.ctx)
		if err != nil {
			return nil, 0, err
		}

		var mids []int64
		for _, v := range ints {
			mids = append(mids, int64(v))
		}
		predicates = append(predicates, member.IDIn(mids...))
	}
	lists, err := s.db.Member.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {

		list = append(list, &sys.SysList{
			ID:   v.ID,
			Name: v.Name,
			Key:  v.Mobile,
		})
	}
	total = int64(len(list))
	return
}

func (s *Sys) ContractList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error) {
	var predicates []predicate.Contract

	if req.Name != "" {
		predicates = append(predicates, contract.Name(req.Name))
	}

	lists, err := s.db.Contract.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, &sys.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))
	return
}

func (s *Sys) StaffList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error) {
	var predicates []predicate.User

	if req.Name != "" {
		predicates = append(predicates, user.NameEQ(req.Name))
	}

	lists, err := s.db.User.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, &sys.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))
	return
}

func NewSys(ctx context.Context, c *app.RequestContext) do.Sys {
	return &Sys{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
