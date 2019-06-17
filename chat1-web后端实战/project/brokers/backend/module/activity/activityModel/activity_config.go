package activityModel

import "encoding/json"

// 活动配置
type ActivityConfig struct{
	Id  int    `gorm:"column:id;default:" json:"id" form:"id"`
	State  int    `gorm:"column:state;default:" json:"state" form:"state"`
	OpenConfig  json.RawMessage    `gorm:"column:open_config;default:" json:"open_config" form:"open_config"`
	RewardConfig  json.RawMessage    `gorm:"column:reward_config;default:" json:"reward_config" form:"reward_config"`
}

func (o ActivityConfig) TableName() string {
	return "activity_config"
}

// 用户进度
type UserActivityProcess struct {
	Id          int             `gorm:"column:id;default:" json:"id" form:"id"`
	UserId      int             `gorm:"column:user_id;default:" json:"user_id" form:"user_id"`
	ActivityId  int             `gorm:"column:activity_id;default:" json:"activity_id" form:"activity_id"`
	JointConfig json.RawMessage `gorm:"column:joint_config;default:" json:"joint_config" form:"joint_config"`
}

func (o UserActivityProcess) TableName() string {
	return "user_activity_process"
}

