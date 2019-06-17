package main

import (
	"context"
	"fmt"
	"gitchat/chat1-web后端实战/project/brokers/activities/inviteActivity/common"
	"gitchat/chat1-web后端实战/project/brokers/activities/inviteActivity/pb"
	"github.com/gin-gonic/gin"
)

func inviteeClick(c *gin.Context) {
	type Param struct {
		InviterId int `json:"inviter_id" binding:"required"`
	}
	var param Param
	if e:= c.Bind(&param);e!=nil{
		c.JSON(400, gin.H{"message": e.Error()})
		return
	}

	conn := pb.NewPropServiceClient(common.GrpcClient)
	rsp, e := conn.PresentProp(context.TODO(), &pb.PresentPropRequest{
         UserId: int32(param.InviterId),
         PropBlock: &pb.PropBlock{
         	PropId:1,
         	PropNum:100,
         	ExpireIn: -1,
         	PropTitle:"冒险家的证明",
		 },

	})
	if e != nil {
		c.JSON(500, gin.H{
			"message":       "app远程调用出错",
			"debug_message": fmt.Sprintf("receive status '%d', message '%s'", rsp.Status, rsp.Message),
		})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
