package logic

import (
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"

	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailReply, err error) {
	// todo: add your logic here and delete this line
	// 从数据库查询当前用户
	user:=new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("identity = ? ", req.Identity).Get(user)
	if err!=nil {
		return nil,err
	}
	if !has{
		return nil,errors.New("用户不存在")
	}
	resp = new(types.UserDetailReply)
	resp.Email = user.Email
	resp.Name = user.Name
	return
}
