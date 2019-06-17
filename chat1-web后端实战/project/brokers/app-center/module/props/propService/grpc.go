package propService

import (
	"context"
	"github.com/fwhezfwhez/errorx"
	"gitchat/chat1-web后端实战/project/brokers/app-center/common/dependent/db"
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/props/propModel"
	"gitchat/chat1-web后端实战/project/brokers/app-center/module/props/propPb"
)

type PropService struct {
}

func (ps *PropService) PresentProp(ctx context.Context, req *propPb.PresentPropRequest) (*propPb.PresentPropResponse, error) {
	if e := db.DB.Model(&propModel.UserProp{}).Create(&propModel.UserProp{
		UserId:    int(req.UserId),
		PropId:    int(req.PropBlock.PropId),
		ExpireIn:  int(req.PropBlock.ExpireIn),
		PropNum:   int(req.PropBlock.PropNum),
		PropTitle: req.PropBlock.PropTitle,
	}).Error; e != nil {
		return &propPb.PresentPropResponse{
			Message: errorx.Wrap(e).Error(),
			Status:  500,
		}, errorx.Wrap(e)
	}
	return &propPb.PresentPropResponse{
		Message: "success",
		Status:  200,
	}, nil
}
