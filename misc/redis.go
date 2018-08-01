package misc

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"onemore-service-iris/config"
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
		MaxActive:   config.Conf.Redis.MaxActive,
		IdleTimeout: time.Second * time.Duration(config.Conf.Redis.Idle),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Conf.Redis.Addr)
			if err != nil {
				return nil, err
			}
			if config.Conf.Redis.Pwd != "" {
				if _, err := c.Do("AUTH", config.Conf.Redis.Pwd); err != nil {
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

func usecase() {
	//从redis连接池中获取一个连接
	r := Redix.Get()
	//把连接还给连接池
	defer r.Close()
	//TODO do your business
}
