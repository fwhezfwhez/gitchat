package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var pool = getRedis("redis://localhost:6379")

func (u *User) SyncRedis(conn redis.Conn) {
	if conn == nil {
		conn = pool.Get()
		defer conn.Close()
	}
	buf, _ := json.Marshal(u)
	key := fmt.Sprintf("gitchat:user_info:%d", u.Id)
	_, e := conn.Do("SETEX", key, 60*60*24, buf)
	if e != nil {
		panic(e)
	}
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
			"localhost",
			"5433",
			"gitchat",
			"test",
			"disable",
			"123456",
		),
	)

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetConnMaxLifetime(10 * time.Second)
	db.DB().SetMaxIdleConns(30)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/user/", get)
	r.GET("/user/:id/", getOne)
	r.POST("/user/", post)

	r.PATCH("/user/:id/", patch)
	r.DELETE("/user/:id/", deleteById)

	r.Run(":8082")
}

func get(c *gin.Context) {
	var users []User
	if e := db.Raw("select * from user_info").Scan(&users).Error; e != nil {
		c.JSON(500, gin.H{"message": e.Error()})
		return
	}
	conn := pool.Get()
	defer conn.Close()
	for i, _ := range users {
		users[i].SyncRedis(conn)
	}
	c.JSON(200, users)
}
func post(c *gin.Context) {
	var user User
	if e := c.Bind(&user); e != nil {
		panic(e)
	}

	if e := db.Raw("insert into user_info(username) values(?) returning *", user.Username).Scan(&user).Error; e != nil {
		c.JSON(500, gin.H{"message": e.Error()})
		return
	}
	user.SyncRedis(nil)
	c.JSON(200, user)
}
func patch(c *gin.Context) {
	var user User
	c.Bind(&user)
	id := c.Param("id")
	db.Raw("update user_info set username=? where id=? returning *", user.Username, id).Scan(&user)
	user.SyncRedis(nil)
	c.JSON(200, user)
}
func deleteById(c *gin.Context) {
	id := c.Param("id")
	db.Exec("delete from user_info where id=?", id)
	c.JSON(200, gin.H{"message": "success"})
}
func getOne(c *gin.Context) {
	var user User
	id := c.Param("id")

	conn :=pool.Get()
	defer conn.Close()

	buf,e :=redis.Bytes(conn.Do("GET", fmt.Sprintf("gitchat:user_info:%s", id)))
	if e!=nil {
		panic(e)
	}

	if len(buf) !=0 {
		e= json.Unmarshal(buf, &user)
		if e!=nil {
			panic(e)
		}
	} else {
		db.Raw("select * from user_info where id=?", id).Scan(&user)
	}

	c.JSON(200, user)
}
func getRedis(url string) *redis.Pool {
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
