package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"
	"log"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	// todo: add your logic here and delete this line
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("该邮箱验证码不存在或已过期")
	}
	if code != req.Code {
		err = errors.New("验证码错误")
		return
	}
	count, err := l.svcCtx.Engine.Where("name=?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		err = errors.New("用户名已存在")
		return
	}
	user := &models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
	insert, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println("insert user row: ", insert)
	return

}
