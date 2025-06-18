package do

import "kcers-survey/idl_gen/model/logs"

type Logs interface {
	Create(logsReq *logs.LogsInfo) error
	List(req *logs.LogsListReq) (list []*logs.LogsInfo, total int, err error)
	DeleteAll() error
}
