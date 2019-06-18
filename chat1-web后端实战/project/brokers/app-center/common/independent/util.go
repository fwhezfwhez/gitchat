package independent

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

type SyncRedisI interface {
	RedisKey(mode string) string
}

// it can sync data from model to redis, requiring model realized RedisKey(string)string
func SyncRedis(o SyncRedisI, mode string, conn redis.Conn) error {
	if mode == "" {
		mode = "dev"
	}

	buf, e := json.Marshal(o)
	if e != nil {
		return e
	}
	_, e = conn.Do("SETEX", o.RedisKey(mode), 60*60*24, buf)
	if e != nil {
		return  e
	}
	return nil
}

// var obj Obj
// BindFromRedis(&obj, "dev", conn)
func BindFromRedis(o SyncRedisI, mode string,conn redis.Conn) (error){
    buf, e:=redis.Bytes(conn.Do("GET", o.RedisKey(mode)))
    if e!=nil {
    	return e
	}
    e = json.Unmarshal(buf, o)
	if e!=nil {
		return e
	}
    return nil
}
