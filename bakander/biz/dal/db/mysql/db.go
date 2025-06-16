package db

import (
	"kcers-survey/biz/dal/config"
	"kcers-survey/biz/dal/db/mysql/ent"
	"sync"
)

var onceClient sync.Once

var DB *ent.Client

func InitDB() {
	onceClient.Do(func() {
		DB = InItDB(config.GlobalServerConfig.MySQLInfo.Host, config.GlobalServerConfig.IsProd)
	})
}
