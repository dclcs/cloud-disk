package logic

import (
	"context"
	"errors"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	user := new(models.UserBasic)
	has, err := models.Engine.Where("name = ? and password = ?", req.Name, helper.MD5(req.Password)).Get(user)

	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.New("用户名或密码错误")
	}

	token, err := helper.GenerateToken(uint64(user.Id), user.Identity, user.Name)

	if err != nil {
		return nil, err
	}

	resp = new(types.LoginReply)
	resp.Token = token
	return
}
