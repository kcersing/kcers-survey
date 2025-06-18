package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"kcers-survey/biz/dal/config"
	"time"
)

var (
	RDUser      *redis.Client
	RDMember    *redis.Client
	RDSmsUser   *redis.Client
	RDSmsMember *redis.Client
	ctx         = context.Background()
)

func InitRedis() {
	RDUser = redis.NewClient(&redis.Options{
		Addr:     config.GlobalServerConfig.Redis.Host,
		Password: config.GlobalServerConfig.Redis.Password,
		DB:       1,
	})
	RDMember = redis.NewClient(&redis.Options{
		Addr:     config.GlobalServerConfig.Redis.Host,
		Password: config.GlobalServerConfig.Redis.Password,
		DB:       2,
	})
	RDSmsUser = redis.NewClient(&redis.Options{
		Addr:     config.GlobalServerConfig.Redis.Host,
		Password: config.GlobalServerConfig.Redis.Password,
		DB:       3,
	})
	RDSmsMember = redis.NewClient(&redis.Options{
		Addr:     config.GlobalServerConfig.Redis.Host,
		Password: config.GlobalServerConfig.Redis.Password,
		DB:       4,
	})

}

const (
	UserSmsKey     = ":userSms"
	MemberSmsKey   = ":memberSms"
	UserTokenKey   = ":userToken"
	MemberTokenKey = ":memberToken"
)

type (
	Rds struct {
		Key         string
		RedisClient *redis.Client
		Expiration  time.Duration
	}
)

func (r *Rds) Add(id, str string) {
	add(r.RedisClient, id+r.Key, str, r.Expiration)
}

func (r *Rds) Del(id, str string) {
	del(r.RedisClient, id+r.Key, str, r.Expiration)
}

func (r *Rds) Check(id string) bool {
	return check(r.RedisClient, id+r.Key)
}

func (r *Rds) Exist(id, str string) bool {
	return exist(r.RedisClient, id+r.Key, str, r.Expiration)
}

func (r *Rds) Get(id string) []interface{} {
	return get(r.RedisClient, id+r.Key, r.Expiration)
}
