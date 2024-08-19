package tools

import (
	"github.com/go-redis/redis"
	"testing"
)

func TestZAdd(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 8,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		t.Fatalf("redis ping err: %v", err)
		return
	}

	list := []redis.Z{
		{Score: 0, Member: 1},
		{Score: 0, Member: 2},
	}

	rdb.ZAdd("moving", list...)

	result, err := rdb.ZRangeWithScores("moving", 0, 10).Result()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(result)
}

func TestHSet(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 8,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		t.Fatalf("redis ping err: %v", err)
		return
	}

	rdb.HSet("followed", "uid:1", 10)

	result, err := rdb.HGet("followed", "uid:1").Result()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(result)
}
