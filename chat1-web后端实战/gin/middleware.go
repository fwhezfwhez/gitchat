package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func validateToken(c *gin.Context) {
	//if c.Request.URL.String() == "/login/" {
	//	c.Next()
	//	return
	//}
	token := c.Request.Header.Get("Authorization")
	if token != "a token example" {
		c.Abort()
		c.JSON(403, gin.H{"message": "token invalid"})
		return
	}
	c.Next()
}

func main() {
	r := gin.Default()
	r.POST("/login/", login)
	r.Use(validateToken)
	r.POST("/weather/", weather)
	r.Run(":8081")
}

func login(c *gin.Context) {
	c.JSON(200, gin.H{"token": "a token example"})
}
func weather(c *gin.Context) {
	c.JSON(200, gin.H{"message": fmt.Sprintf("%s 晴", time.Now().Format("2006-01-02"))})
}
