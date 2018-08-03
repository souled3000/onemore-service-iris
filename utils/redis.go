package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func init() {
	fmt.Println("Redis Init:", InitRedix())
}

//redis连接池
var (
	Redix *redis.Pool
)

/*
*addr: ip:port
*pwd: 密码，没有填""
*maxActive: 最大连接数
*idle: 连接空闲时长
 */
func InitRedix() error {
	Redix = newPool()
	r := Redix.Get()
	_, err := r.Do("PING")
	if err != nil {
		return err
	}
	r.Close()
	return err
}
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5,
		MaxActive:   Conf.Redis.MaxActive,
		IdleTimeout: time.Second * time.Duration(Conf.Redis.Idle),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", Conf.Redis.Addr)
			if err != nil {
				return nil, err
			}
			if Conf.Redis.Pwd != "" {
				if _, err := c.Do("AUTH", Conf.Redis.Pwd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return err
			}
			return err
		},
	}
}
