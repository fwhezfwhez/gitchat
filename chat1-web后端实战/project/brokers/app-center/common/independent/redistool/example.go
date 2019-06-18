package redistool
// package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {

	//1.连接redis
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("1.连接redis失败", err)
		return
	}
	fmt.Println("1.连接成功", c)
	defer c.Close()

	//2.数据读写 : SET命令和GET命令
	// 2.1 创建key=user_name,value="ft"的redis数据对象，写入
	_, err = c.Do("MSET", "user_name", "ft")
	if err != nil {
		fmt.Println("数据设置失败:", err)
	}

	username, err := redis.String(c.Do("GET", "user_name"))

	if err != nil {
		fmt.Println("数据获取失败:", err)
	} else {
		fmt.Println("2.1.获取user_name", username)
	}
	// 2.2写入一段时限为5秒过期的内容: EX命令
	_, err = c.Do("SET", "user_name2", "ft2", "EX", "5")
	if err != nil {
		fmt.Println("数据设置失败:", err)
	}

	//未过期
	username2, err := redis.String(c.Do("GET", "user_name2"))
	if err != nil {
		fmt.Println("数据获取失败:", err)
	} else {
		fmt.Printf("2.2直接获取未过期的 user_name2: %v \n", username2)
	}
	//延迟8秒，过期
	time.Sleep(8 * time.Second)
	username2, err = redis.String(c.Do("GET", "user_name2"))
	if err != nil {
		fmt.Println("2.2过期后数据获取失败:", err)
	} else {
		fmt.Printf("2.2延迟后获取不到过期的 user_name2: %v \n", username2)
	}

	//2.3 批量写入和批量写出:MSET,MGET命令
	_, err = c.Do("MSET", "user_name", "ft", "class_name", "UD01")
	if err != nil {
		fmt.Println("批量数据设置失败:", err)
	}

	results, err := redis.Strings(c.Do("MGET", "user_name", "class_name"))

	if err != nil {
		fmt.Println("数据获取失败:", err)
	} else {
		fmt.Println("2.3批量获取成功", results)
	}

	//2.4 判断是否存在某键值对
	If_Exit, err := redis.Bool(c.Do("EXISTS", "class_name"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("2.4 class_name是否存在: %v \n", If_Exit)
	}

	//3 删除键
	affectCount, err := redis.Int(c.Do("DEL", "class_name"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("3.class_name已经删除，受影响行数: %v \n", affectCount)
	}

	//4 存取json对象 :SETNX 等价于SET if not exist
	key := "jsonKey"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, _ := json.Marshal(imap)

	_, err = c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]string

	buf, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(buf, &result)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println("4.获取json对象成功:userName", result["username"])
	fmt.Println("                 phonenumber", result["phonenumber"])

	//5.设置过期时间 : EXPIRE
	_, err = c.Do("EXPIRE", key, 24*60*60)
	if err != nil {
		fmt.Println(err)
	}

	//6.管道 按照队列先进先出的原则进行send,receive操作
	c.Send("SET", "userId", "DF123", "userName", "123456")
	c.Send("GET", "userId", "userName")
	c.Flush()
	c.Receive()                   // reply from SET
	valueGet, errr := c.Receive() // reply from GET
	fmt.Println(redis.String(valueGet, errr))
}
