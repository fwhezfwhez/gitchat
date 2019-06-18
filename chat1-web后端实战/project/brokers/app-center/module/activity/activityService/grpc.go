package activityService

import (
	"context"
	"encoding/json"
	"fmt"
	"gitchat/chat1-web后端实战/project/brokers/app-center/common/dependent/db"
	"gitchat/chat1-web后端实战/project/brokers/app-center/common/independent"
	"gitchat/chat1-web后端实战/project/brokers/app-center/common/independent/redistool"
	"gitchat/chat1-web后端实战/project/brokers/app-center/config"
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/activity/activityModel"
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/activity/activityPb"
	"github.com/fwhezfwhez/errorx"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

type ActivityService struct{}

func (as ActivityService) GetUserActivityProcess(ctx context.Context, request *activityPb.GetUserActivityProcessRequest) (*activityPb.GetUserActivityProcessResponse, error) {
	var uap = activityModel.UserActivityProcess{
		UserId:     int(request.UserId),
		ActivityId: int(request.ActivityId),
	}
	conn := redistool.RedisPool.Get()
	mode := config.Cfg.GetString("mode")
	defer conn.Close()
	e := independent.BindFromRedis(&uap, mode, conn)
	if e != nil {
		log.Println(errorx.Wrap(e).Error())
		var count int
		if e := db.DB.Model(&activityModel.UserActivityProcess{}).Where("user_id=? and activity_id=?", uap.UserId, uap.ActivityId).Count(&count).Error; e != nil {
			return nil, errorx.Wrap(e)
		}
		if count == 0 {
			return nil, errorx.NewFromStringf(fmt.Sprintf("activity_id '%d', user_id '%d' not found", uap.ActivityId, uap.UserId))
		}

		if e := db.DB.Model(&activityModel.UserActivityProcess{}).Where("user_id=? and activity_id=?", uap.UserId, uap.ActivityId).First(&uap).Error; e != nil {
			return nil, errorx.Wrap(e)
		}
	}
	d

	if uap.HasExpire() {
		tmp, e := uap.UpgradeToPgRedis(db.DB, conn, mode)
		return &activityPb.GetUserActivityProcessResponse{
			Id:          int32(tmp.Id),
			UserId:      int32(tmp.UserId),
			ActivityId:  int32(tmp.ActivityId),
			JointConfig: []byte(tmp.JointConfig),
		}, errorx.Wrap(e)
	}

	return &activityPb.GetUserActivityProcessResponse{
		Id:          int32(uap.Id),
		UserId:      int32(uap.UserId),
		ActivityId:  int32(uap.ActivityId),
		JointConfig: []byte(uap.JointConfig),
	}, nil
}

func (as ActivityService) UpdateUserActivityProcess(ctx context.Context, request *activityPb.UpdateUserActivityProcessRequest) (*activityPb.UpdateUserActivityProcessResponse, error) {
	e := UpdateUserActivityProcess(config.Mode, activityModel.UserActivityProcess{
		UserId:      int(request.UserId),
		ActivityId:  int(request.ActivityId),
		JointConfig: json.RawMessage(request.JointConfig),
	})
	return &activityPb.UpdateUserActivityProcessResponse{
		Message: "success",
	}, errorx.Wrap(e)
}
func UpdateUserActivityProcess(mode string, uap activityModel.UserActivityProcess) error {
	if uap.UserId == 0 || uap.ActivityId == 0 {
		return errorx.NewFromStringf("user_id '%d' and activity_id  '%d' is either empty", uap.UserId, uap.ActivityId)
	}

	var count int
	engine := db.DB.Model(&activityModel.UserActivityProcess{}).Where("user_id=? and activity_id=?", uap.UserId, uap.ActivityId)

	if e := engine.Count(&count).Error; e != nil {
		return errorx.Wrap(e)
	}

	if count == 0 {
		// 不存在即创建
		var tmp = activityModel.UserActivityProcess{
			UserId:      uap.UserId,
			ActivityId:  uap.ActivityId,
			DateTime:    time.Now().Format("2006-01-02 15:04:05"),
			JointConfig: uap.JointConfig,
		}
		if e := engine.Create(&tmp).Error; e != nil {
			return errorx.Wrap(e)
		}
		tmp.SyncRedis(config.Mode, nil)
		return nil
	}
	// 存在，即修改
	if e := engine.Updates(&uap).Error; e != nil {
		return errorx.Wrap(e)
	}

	var tmp activityModel.UserActivityProcess
	if e := engine.First(&tmp).Error; e != nil {
		return errorx.Wrap(e)
	}
	return tmp.SyncRedis(config.Mode, nil)
}

func (as ActivityService) GetActivityConfig(ctx context.Context, request *activityPb.GetActivityConfigRequest) (*activityPb.GetActivityConfigResponse, error) {
	ac, e := GetActivityConfig(config.Mode, int(request.Id))
	return &activityPb.GetActivityConfigResponse{
		Id:          int32(ac.Id),
		State:       int32(ac.State),
		AwardConfig: []byte(ac.RewardConfig),
		OpenConfig:  []byte(ac.OpenConfig),
	}, errorx.Wrap(e)
}

func GetActivityConfig(mode string, activityId int) (activityModel.ActivityConfig, error) {
	var activityConfig activityModel.ActivityConfig

	conn := redistool.RedisPool.Get()
	defer conn.Close()

	buf, e := redis.Bytes(conn.Do("GET", activityConfig.RedisKey(mode)))
	if e != nil {
		log.Println(errorx.Wrap(e).Error())
	}
	if e == nil && len(buf) != 0 {
		e2 := json.Unmarshal(buf, &activityConfig)
		if e2 != nil {
			return activityModel.ActivityConfig{}, errorx.Wrap(e2)
		}
		return activityConfig, nil
	}

	var count int
	engine := db.DB.Model(&activityModel.ActivityConfig{}).Where("id=?", activityId)
	if e := engine.Count(&count).Error; e != nil {
		return activityModel.ActivityConfig{}, errorx.Wrap(e)
	}
	if count == 0 {
		return activityModel.ActivityConfig{}, errorx.NewFromStringf("activity_id '%d' not found", activityId)
	}
	if e := engine.First(&activityConfig).Error; e != nil {
		return activityModel.ActivityConfig{}, errorx.Wrap(e)
	}

	return activityConfig, nil

}
func (as ActivityService) UpdateActivityConfig(ctx context.Context, request *activityPb.UpdateActivityConfigRequest) (*activityPb.UpdateActivityConfigResponse, error) {
	e := UpdateActivityConfig(config.Mode, activityModel.ActivityConfig{
		Id:           int(request.Id),
		State:        int(request.State),
		OpenConfig:   json.RawMessage(request.OpenConfig),
		RewardConfig: json.RawMessage(request.AwardConfig),
	})

	return &activityPb.UpdateActivityConfigResponse{
		Message: "success",
	}, errorx.Wrap(e)
}
func UpdateActivityConfig(mode string, ac activityModel.ActivityConfig) error {
	var count int
	engine := db.DB.Model(&activityModel.ActivityConfig{}).Where("id=?", ac.Id)
	if e := engine.Count(&count).Error; e != nil {
		return errorx.Wrap(e)
	}
	if count == 0 {
		return errorx.NewFromStringf("activity_id '%d' not found", ac.Id)
	}

	if e := engine.Updates(&ac).Error; e != nil {
		return errorx.Wrap(e)
	}

	var tmp activityModel.ActivityConfig

	if e := engine.Where("id=?", ac.Id).First(&tmp).Error; e != nil {
		return errorx.Wrap(e)
	}
	return tmp.SyncRedis(mode, nil)
}
