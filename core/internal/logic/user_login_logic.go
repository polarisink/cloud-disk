package logic

import (
	"context"
	"core/helper"
	"core/models"
	"errors"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	//从数据库查询当前用户
	user:=new(models.UserBasic)
	has, err := models.Engine.Where("name = ? AND password = ?", req.Name, req.Password).Get(user)
	if err!=nil {
		return nil,err
	}
	if !has{
		return nil,errors.New("用户名或密码错误")
	}
	//生成token
	token,err:=helper.GenerateToken(user.Id,user.Identity,user.Name)
	if err!=nil {
		return nil,err
	}
	resp = new(types.LoginReply)
	resp.Token = token
	return
}
