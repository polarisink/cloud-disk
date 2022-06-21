package logic

import (
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest) (resp *types.UserFileMoveReply, err error) {
	// todo: add your logic here and delete this line

	return
}
