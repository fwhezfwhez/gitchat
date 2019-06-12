package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var users = []User{{1, "张三"}, {2, "李四"}, {3, "王五"}}

func main() {
	r := gin.Default()
	r.GET("/user/", get)
	r.GET("/user/:id/", getOne)
	r.POST("/user/", post)

	r.PATCH("/user/:id/", patch)
	r.DELETE("/user/:id/", deleteById)

	r.Run(":8080")
}

func get(c *gin.Context) {
	c.JSON(200, users)
}
func post(c *gin.Context) {
	var user User
	c.Bind(&user)

	users = append(users, User{Username: user.Username, Id: user.Id})
	c.JSON(200, users)
}
func patch(c *gin.Context) {
	var user User
	c.Bind(&user)
	id := c.Param("id")
	for i,v:=range users{
		if strconv.Itoa(v.Id) == id {
			users[i].Username = user.Username
			user.Id = v.Id
		}
	}
	c.JSON(200, user)
}
func deleteById(c *gin.Context) {
	id := c.Param("id")
	for i,v:=range users{
		if strconv.Itoa(v.Id) == id {
			users = append(users[:i], users[i+1:]...)
		}
	}
	c.JSON(200, users)
}
func getOne(c *gin.Context) {
	var user User
	id := c.Param("id")
	for _,v:=range users{
		if strconv.Itoa(v.Id) == id {
			user = v
			break
		}
	}
	c.JSON(200, user)
}
