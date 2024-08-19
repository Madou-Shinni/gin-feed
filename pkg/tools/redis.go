package tools

import (
	"github.com/go-redis/redis"
	"time"
)

func RedisHSet(rdb redis.Cmdable, key, filed string, val any) error {
	rdb.SAdd(key, filed, val)
	rdb.Expire(key, time.Second*5) // 延长缓存的时间
	//rdb.SMembers().ScanSlice()
	return nil
}
