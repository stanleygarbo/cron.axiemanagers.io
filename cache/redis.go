package cache

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	Pool *redis.Pool
)

func InitRedis() {
	Pool = &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}

	RemoveScholarChartFromCache("000000000000000000000000")
}

func RemoveScholarChartFromCache(ronin string) {
	rd := Pool.Get()
	_, err := rd.Do("DEL", ronin + "_chart")
	if err != nil{
		fmt.Printf("@RemoveScholarChart Err: %v", err)
	}
}
