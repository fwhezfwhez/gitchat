package activityRouter

import (
	"gitchat/chat1-web后端实战/project/brokers/backend/module/activity/activityService"
	"github.com/gin-gonic/gin"
)

func Router(r gin.IRoutes) {
	// 增加活动配置
    r.POST("/activity/", activityService.AddActivity)
    // 列表活动配置
    r.GET("/activity/", activityService.ListActivity)
    // 修改活动配置
    r.PATCH("/activity/:id/", activityService.ModifyActivity)

    // 显示用户进度列表
    r.GET("/user/activity/process/", activityService.GetUserActivityProcess)
    // 按照id获取某条用户进度
    r.GET("/user/activity/process/:id/", activityService.GetOneUserActivityProcessById)
    // 按照user_id,activity_id获取用户进度
    r.POST("/user-activity-process/", activityService.GetSpecificUserActivityProcess)
}
