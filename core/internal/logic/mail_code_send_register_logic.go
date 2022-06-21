package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	// todo: add your logic here and delete this line
	count, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("该邮箱已被注册")
		return
	}
	code := helper.RandCode()
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	err = helper.MailCodeSend(req.Email, code)
	/*if err != nil {
		return nil, err
	}*/
	return
}
