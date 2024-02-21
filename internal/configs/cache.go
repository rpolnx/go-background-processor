package configs

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func NewCachePool() *redis.Pool {
	return &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				fmt.Sprintf("%s:%d", GlobalAppConfig.Redis.Host, GlobalAppConfig.Redis.Port),
				redis.DialUsername(GlobalAppConfig.Redis.User),
				redis.DialPassword(GlobalAppConfig.Redis.Password),
				redis.DialDatabase(GlobalAppConfig.Redis.Db),
			)
		},
	}
}
