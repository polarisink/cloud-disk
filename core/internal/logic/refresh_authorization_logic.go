package logic

import (
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthorizationRequest) (resp *types.RefreshAuthorizationReply, err error) {
	// todo: add your logic here and delete this line

	return
}
