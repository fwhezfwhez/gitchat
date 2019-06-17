package main

import (
	"fmt"
	"gitchat/chat1-web后端实战/project/brokers/activities/inviteActivity/config"
	"gitchat/chat1-web后端实战/project/brokers/activities/inviteActivity/util/independent"
	"github.com/gin-gonic/gin"
	"time"
)

func AppFilter(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	fmt.Println(config.Cfg.GetString("app.secret"))

	if auth == "" {
		c.JSON(403, gin.H{"message": "未授权来源", "debug_message": "request.Header['Authorization'] is empty", "tip_id": 1})
		c.Abort()
		return
	}
	// 连续两日的hash有效, 并每日变动
	if independent.MD5(fmt.Sprintf("date=%s&secret=%s", time.Now().Local().Format("2006-01-02"), config.Cfg.GetString("app.secret"))) != auth &&
		independent.MD5(fmt.Sprintf("date=%s&secret=%s", time.Now().AddDate(0, 0, -1).Format("2006-01-02"), config.Cfg.GetString("app.secret"))) != auth {
		c.JSON(403, gin.H{
			"message":       "未授权来源",
			"debug_message": fmt.Sprintf("request.Header['Authorization'] invalid, got '%s'", auth),
			"tip_id":        1,
		})
		c.Abort()
		return
	}
	c.Next()
}
