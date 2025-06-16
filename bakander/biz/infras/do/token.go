package do

import "kcers-survey/idl_gen/model/token"

type Token interface {
	Create(req *token.TokenInfo) error
	Update(req *token.TokenInfo) error
	IsExistByUserId(userId int64) bool
	Delete(userId int64) error
	List(req *token.TokenListReq) (res []*token.TokenInfo, total int, err error)
}
