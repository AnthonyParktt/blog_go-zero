package post

import (
	"context"

	"go-zero_less/postcenter/cmd/api/internal/svc"
	"go-zero_less/postcenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostsLogic) GetPosts(req *types.GetPostsReq) (resp *types.GetPostsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
