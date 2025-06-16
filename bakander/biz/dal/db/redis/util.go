package redis

import (
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

// add k & v to redis
func add(c *redis.Client, k string, v interface{}, expiration time.Duration) {
	tx := c.TxPipeline()
	tx.SAdd(ctx, k, v)
	tx.Expire(ctx, k, expiration)
	tx.Exec(ctx)
}

// del k & v
func del(c *redis.Client, k string, v interface{}, expiration time.Duration) {
	tx := c.TxPipeline()
	tx.SRem(ctx, k, v)
	tx.Expire(ctx, k, expiration)
	tx.Exec(ctx)
}

// check the set of k if exist
func check(c *redis.Client, k string) bool {
	if e, _ := c.Exists(ctx, k).Result(); e > 0 {
		return true
	}
	return false
}

// exist check the relation k and v if exist
func exist(c *redis.Client, k string, v interface{}, expiration time.Duration) bool {
	if e, _ := c.SIsMember(ctx, k, v).Result(); e {
		c.Expire(ctx, k, expiration)
		return true
	}
	return false
}

// count get the size of the set of key
func count(c *redis.Client, k string, expiration time.Duration) (sum int64, err error) {
	if sum, err = c.SCard(ctx, k).Result(); err == nil {
		c.Expire(ctx, k, expiration)
		return sum, err
	}
	return sum, err
}

func get(c *redis.Client, k string, expiration time.Duration) (vt []interface{}) {
	v, _ := c.SMembers(ctx, k).Result()
	c.Expire(ctx, k, expiration)
	for _, vs := range v {
		vI64, _ := strconv.ParseInt(vs, 10, 64)
		vt = append(vt, vI64)
	}
	return vt
}
