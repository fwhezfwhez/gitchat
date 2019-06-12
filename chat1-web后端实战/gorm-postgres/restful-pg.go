package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
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
	c.JSON(200, users)
}
func post(c *gin.Context) {
	var user User
	if e:=c.Bind(&user);e!=nil{
		panic(e)
	}

	if e := db.Raw("insert into user_info(username) values(?) returning *", user.Username).Scan(&user).Error; e != nil {
		c.JSON(500, gin.H{"message": e.Error()})
		return
	}
	c.JSON(200, user)
}
func patch(c *gin.Context) {
	var user User
	c.Bind(&user)
	id := c.Param("id")
	db.Raw("update user_info set username=? where id=? returning *", user.Username, id).Scan(&user)
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
	db.Raw("select * from user_info where id=?", id).Scan(&user)
	c.JSON(200, user)
}
