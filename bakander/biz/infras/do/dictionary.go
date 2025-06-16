package do

import "kcers-survey/idl_gen/model/dictionary"

type Dictionary interface {
	Create(req *dictionary.DictionaryInfo) error
	Update(req *dictionary.DictionaryInfo) error
	Delete(id int64) error
	List(req *dictionary.DictListReq) (list []*dictionary.DictionaryInfo, total int, err error)
	CreateDetail(req *dictionary.DictionaryDetail) error
	UpdateDetail(req *dictionary.DictionaryDetail) error
	DeleteDetail(id int64) error
	DetailListByDict(req *dictionary.DetailListReq) (list []*dictionary.DictionaryDetail, total int64, err error)
	DetailByDictNameAndKey(dictName, key string) (detail *dictionary.DictionaryDetail, err error)
	GetDictionaryDetailTitle(id int64) string
}
