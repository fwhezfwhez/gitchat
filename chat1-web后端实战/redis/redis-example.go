package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
    var pool = GetRedis("redis://localhost:6379")
    conn := pool.Get()
    defer conn.Close()
    key:= "test_redis_key"
    value := "test_redis_value"
    _,e:=conn.Do("SETEX", key, 60*60*24 ,value)
    if e!=nil {
    	panic(e)
	}

    getValue,e:=redis.String(conn.Do("GET", key))
	if e!=nil {
		panic(e)
	}
    fmt.Println("receive from redis key 'test_redis_key':" ,getValue)
}

func GetRedis(url string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 200,
		//MaxActive:   0,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(url)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
