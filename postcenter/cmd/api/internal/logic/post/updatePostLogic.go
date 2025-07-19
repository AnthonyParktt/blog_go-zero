package post

import (
	"context"

	"go-zero_less/postcenter/cmd/api/internal/svc"
	"go-zero_less/postcenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePostLogic) UpdatePost(req *types.CreatePostReq) (resp *types.CreatePostResp, err error) {
	// todo: add your logic here and delete this line

	return
}
