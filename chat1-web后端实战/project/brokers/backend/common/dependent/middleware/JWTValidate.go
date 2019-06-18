package middleware

import (
	jwt_util "gitchat/chat1-web后端实战/project/brokers/backend/common/independent/jwt-util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTValidate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(402, gin.H{"message": "valid fail"})
		c.Abort()
		return
	}

	tk, info := jwt_util.JwtTool.ValidateJWT(token)
	if !tk.Valid {
		c.JSON(402, gin.H{"message": info})
		c.Abort()
		return
	}
	r := tk.Claims.(jwt.MapClaims)
	c.Set("user_id", r["user_id"])
}
