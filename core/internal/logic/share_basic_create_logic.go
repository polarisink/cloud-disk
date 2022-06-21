package logic

import (
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest) (resp *types.ShareBasicCreateReply, err error) {
	// todo: add your logic here and delete this line

	return
}
