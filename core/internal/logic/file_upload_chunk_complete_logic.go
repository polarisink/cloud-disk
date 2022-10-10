package logic

import (
	"context"
	"github.com/zeromicro/zero-contrib/logx/zapx"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkCompleteLogic {
	return &FileUploadChunkCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkCompleteLogic) FileUploadChunkComplete(req *types.FileUploadChunkCompleteRequest) (resp *types.FileUploadChunkCompleteReply, err error) {
	// todo: add your logic here and delete this line
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)
	return
}
