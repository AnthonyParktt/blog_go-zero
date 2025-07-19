package post

import (
	"context"
	"go-zero_less/postcenter/cmd/api/internal/svc"
	"go-zero_less/postcenter/cmd/api/internal/types"
	"go-zero_less/postcenter/cmd/rpc/pb"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostReq) (resp *types.CreatePostResp, err error) {
	userId := l.ctx.Value("userId")
	userIdInt := cast.ToUint64(userId)
	postDTO := pb.PostInfoCreateRequest{
		PostInfo: &pb.PostInfoBase{
			UserId:  userIdInt,
			Title:   req.Title,
			Content: req.Content,
		},
	}
	postInfoId, err := l.svcCtx.PostRpcClient.CreatePost(l.ctx, &postDTO)
	if err != nil {
		return nil, err
	}
	return &types.CreatePostResp{
		Id: postInfoId.Id,
	}, nil
}
