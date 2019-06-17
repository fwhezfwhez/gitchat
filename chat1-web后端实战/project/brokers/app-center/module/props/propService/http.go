package propService

import (
	"fmt"
	"gitchat/chat1-web后端实战/project/brokers/app-center/common/dependent/db"
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/props/propModel"
	"github.com/fwhezfwhez/errorx"
	"github.com/gin-gonic/gin"
)

func HTTPGetProps(c *gin.Context) {
	var props []propModel.Prop
	propName := c.DefaultQuery("prop_name", "")
	propId := c.DefaultQuery("prop_id", "")

	engine := db.DB.Model(&propModel.Prop{})

	if propName != "" {
		engine = engine.Where("prop_name like ?", "%"+propName+"%")
	}
	if propId != "" {
		engine = engine.Where("prop_id = ?", propId)
	}

	if e := engine.Order("id desc").Find(&props).Error; e != nil {
		c.JSON(500, gin.H{"message": "服务错误", "debug_message": errorx.Wrap(e).Error()})
		return
	}
	if len(props) == 0 {
		c.JSON(200, gin.H{"message": "no record"})
		return
	}
	c.JSON(200, props)
}

func HTTPGetOneProp(c *gin.Context) {
	id := c.Param("id")
	var prop propModel.Prop
	var count int
	if e := db.DB.Model(&propModel.Prop{}).Where("id=?", id).Count(&count).First(&prop).Error; e != nil {
		c.JSON(500, gin.H{"message": "服务错误", "debug_message": errorx.Wrap(e).Error()})
		return
	}
	if count == 0 {
		c.JSON(200, gin.H{"message": fmt.Sprintf("id '%s' not found", id)})
		return
	}
	c.JSON(200, prop)
}
func HTTPAddProp(c *gin.Context) {
	var param propModel.Prop
	if e := c.Bind(&param); e != nil {
		c.JSON(400, gin.H{"message": e.Error()})
		return
	}
	if e := db.DB.Model(&propModel.Prop{}).Create(&param).Error; e != nil {
		c.JSON(500, gin.H{"message": "服务错误", "debug_message": errorx.Wrap(e).Error()})
		return
	}
	c.JSON(200, param)
}
func HTTPModifyProp(c *gin.Context) {
	var param propModel.Prop
	id := c.Param("id")
	if e := c.Bind(&param); e != nil {
		c.JSON(400, gin.H{"message": e.Error()})
		return
	}
	if e := db.DB.Model(&propModel.Prop{}).Where("id=?", id).Updates(&param).Error; e != nil {
		c.JSON(500, gin.H{"message": "服务错误", "debug_message": errorx.Wrap(e).Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
func HTTPDeleteProp(c *gin.Context) {
	id := c.Param("id")
	engine := db.DB.Model(&propModel.Prop{})
	if id != "" {
		engine = db.DB.Model(&propModel.Prop{}).Where("id=?", id)
	}
	if e := engine.Delete(&propModel.Prop{}).Error; e != nil {
		c.JSON(500, gin.H{"message": "服务错误", "debug_message": errorx.Wrap(e).Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func HTTPGetUserProp(c *gin.Context) {
	var userProps []propModel.UserProp

	userId := c.DefaultQuery("user_id", "")
	propId := c.DefaultQuery("prop_id", "")
	engine := db.DB.Model(&propModel.UserProp{})

	if userId != "" {
		engine = engine.Where("user_id=?", userId)
	}
	if propId != "" {
		engine = engine.Where("prop_id=?", propId)
	}

	var count int
	if e := engine.Count(&count).Error; e != nil {
		c.JSON(500, gin.H{"message": "服务错误", "debug_message": errorx.Wrap(e).Error()})
		return
	}
	if count == 0 {
		c.JSON(200, gin.H{"message": "no record"})
		return
	}
	if e := engine.Find(&userProps).Error; e != nil {
		c.JSON(500, gin.H{"message": "服务错误", "debug_message": errorx.Wrap(e).Error()})
		return
	}
	c.JSON(200, userProps)
}
func HTTPDeleteUserProp(c *gin.Context) {
	id := c.Param("id")
	if e := db.DB.Model(&propModel.UserProp{}).Where("id=?", id).Delete(&propModel.UserProp{}).Error; e != nil {
		c.JSON(500, gin.H{"message": "服务错误", "debug_message": errorx.Wrap(e).Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
