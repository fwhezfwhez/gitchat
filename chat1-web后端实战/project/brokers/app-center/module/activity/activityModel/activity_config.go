package activityModel

import (
	"encoding/json"
	"fmt"
	"gitchat/chat1-web后端实战/project/brokers/app-center/common/independent/redistool"
	"github.com/fwhezfwhez/errorx"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

// 活动配置
type ActivityConfig struct {
	Id           int             `gorm:"column:id;default:" json:"id" form:"id"`
	State        int             `gorm:"column:state;default:" json:"state" form:"state"`
	OpenConfig   json.RawMessage `gorm:"column:open_config;default:" json:"open_config" form:"open_config"`
	RewardConfig json.RawMessage `gorm:"column:reward_config;default:" json:"reward_config" form:"reward_config"`
}

func (o ActivityConfig) TableName() string {
	return "activity_config"
}

func (o ActivityConfig) RedisKey(mode string) string {
	if mode == "" {
		mode = "dev"
	}
	return fmt.Sprintf("game:activity:%s:%d", mode, o.Id)
}
func (o ActivityConfig) SyncRedis(mode string, conn redis.Conn) error {
	if mode == "" {
		mode = "dev"
	}
	if conn == nil {
		conn = redistool.RedisPool.Get()
		defer conn.Close()
	}

	buf, e := json.Marshal(o)
	if e != nil {
		log.Println(errorx.Wrap(e).Error())
		return e
	}
	var key = fmt.Sprintf("game:activity:%s:%d", mode, o.Id)
	_, e = conn.Do("SETEX", key, 60*60*24*7, buf)
	if e != nil {
		log.Println(errorx.Wrap(e).Error())
		return e
	}
	return nil
}

// 用户进度
type UserActivityProcess struct {
	Id          int             `gorm:"column:id;default:" json:"id" form:"id"`
	UserId      int             `gorm:"column:user_id;default:" json:"user_id" form:"user_id"`
	DateTime    string          `gorm:"column:date_time;default:" json:"date_time" form:"date_time"`
	ActivityId  int             `gorm:"column:activity_id;default:" json:"activity_id" form:"activity_id"`
	JointConfig json.RawMessage `gorm:"column:joint_config;default:" json:"joint_config" form:"joint_config"`
}

func (o UserActivityProcess) TableName() string {
	return "user_activity_process"
}

func (o UserActivityProcess) RedisKey(mode string) string {
	if mode == "" {
		mode = "dev"
	}
	return fmt.Sprintf("game:user_activity_process:%s:%d:%d", mode, o.UserId, o.ActivityId)
}
func (o UserActivityProcess) SyncRedis(mode string, conn redis.Conn) error {
	if mode == "" {
		mode = "dev"
	}
	if conn == nil {
		conn = redistool.RedisPool.Get()
		defer conn.Close()
	}

	buf, e := json.Marshal(o)
	if e != nil {
		log.Println(errorx.Wrap(e).Error())
		return errorx.Wrap(e)
	}
	_, e = conn.Do("SETEX", o.RedisKey(mode), 60*60*24, buf)
	if e != nil {
		log.Println(errorx.Wrap(e).Error())
		return errorx.Wrap(e)
	}
	return nil
}

// o 必须具备user_id与activity_id
func (o UserActivityProcess) UpgradeToPgRedis(db *gorm.DB, conn redis.Conn, mode string) (UserActivityProcess, error) {

	if e := db.Model(&o).Where("user_id=? and activity_id =?", o.UserId, o.ActivityId).Updates(&o).Error; e != nil {
		log.Println(e.Error())
		return UserActivityProcess{}, errorx.Wrap(e)
	}
	var uap UserActivityProcess
	if e := db.Model(&UserActivityProcess{}).First(&uap).Error; e != nil {
		log.Println(e.Error())
		return UserActivityProcess{}, errorx.Wrap(e)
	}

	return uap, errorx.Wrap(uap.SyncRedis(mode, conn))
}

func (o UserActivityProcess) HasExpire(latestFreshTime time.Time)bool {
    dt,e :=time.ParseInLocation("2006-01-02 15:04:05", o.DateTime, time.Local);
    if e!=nil {
    	log.Println(errorx.Wrap(e).Error())
    	return true
	}
    if dt.Unix() < latestFreshTime.Unix() {
    	// 过期了
    	return true
	}
    return false
}
