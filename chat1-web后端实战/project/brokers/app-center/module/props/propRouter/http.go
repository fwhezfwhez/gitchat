package propRouter

import (
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/props/propService"
	"github.com/gin-gonic/gin"
)

func HTTPRouter(r gin.IRoutes) {
	// 道具表 crud
	r.GET("/prop/", propService.HTTPGetProps)
	r.GET("/prop/:id/", propService.HTTPGetOneProp)
    r.POST("/prop/", propService.HTTPAddProp)
	r.PATCH("/prop/:id/", propService.HTTPModifyProp)
	r.DELETE("prop/:id/", propService.HTTPDeleteProp)

	// 用户背包表
	r.GET("/user/prop/", propService.HTTPGetUserProp)
	r.DELETE("/user/prop/:id/", propService.HTTPDeleteUserProp)
}
