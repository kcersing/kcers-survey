package do

import "kcers-survey/idl_gen/model/token"

type Token interface {
	Create(req *token.TokenInfo) error
	Update(req *token.TokenInfo) error
	IsExistByUserId(userId int64) bool
	Delete(userId int64) error
	List(req *token.TokenListReq) (res []*token.TokenInfo, total int, err error)

	CreateMemberToken(req *token.MemberTokenInfo) error
	UpdateMemberToken(req *token.MemberTokenInfo) error
	IsExistByMemberIdToken(memberId int64) bool
	DeleteMemberToken(MemberId int64) error
	MemberTokenList(req *token.TokenListReq) (res []*token.MemberTokenInfo, total int, err error)
}
