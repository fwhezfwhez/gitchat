package main

import (
	"fmt"
	"gitchat/chat1-web后端实战/project/brokers/backend/common/dependent/db"
	"gitchat/chat1-web后端实战/project/brokers/backend/common/dependent/middleware"
	jwt_util "gitchat/chat1-web后端实战/project/brokers/backend/common/independent/jwt-util"
	"gitchat/chat1-web后端实战/project/brokers/backend/module/activity/activityRouter"
	"gitchat/chat1-web后端实战/project/brokers/backend/module/props/propRouter"
	"gitchat/chat1-web后端实战/project/brokers/backend/module/user/userModel"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.POST("/login/", genToken)
	r.Use(middleware.JWTValidate)
	// activity
	activityRouter.Router(r)
	// prop
	propRouter.HTTPRouter(r)
	s := &http.Server{
		Addr:           "8001",
		Handler:        cors.AllowAll().Handler(r),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 21,
	}
	fmt.Println("backend http srv  listens on 8002")
	s.ListenAndServe()
}

func genToken(c *gin.Context) {
	var param userModel.BackendUser
	if e := c.Bind(&param); e != nil {
		c.JSON(400, gin.H{"message": e.Error()})
		return
	}
	var count int
	engine := db.DB.Model(&userModel.BackendUser{}).
		Where("username = ?", param.Username)
	if e := engine.Count(&count).Error; e != nil {
		c.JSON(500, gin.H{"message": e.Error()})
		return
	}
	if count == 0 {
		c.JSON(400, gin.H{"message": "用户名不存在"})
		return
	}

	var user userModel.BackendUser
	if e := engine.First(&user).Error; e != nil {
		c.JSON(500, gin.H{"message": e.Error()})
		return
	}
	if user.Password != MD5(fmt.Sprintf("salt=%s&password=%s", cfg.GetString("backend.secret"), param.Password)) {
		c.JSON(400, gin.H{"message": "密码错误"})
		return
	}
	token, e := jwt_util.JwtTool.GenerateJWT(map[string]interface{}{
		"user_id":  user.Id,
		"username": user.Username,
		"role_id":  user.RoleId,
	})
	if e != nil {
		c.JSON(500, gin.H{"message": e.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
