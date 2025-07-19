package comments

import (
	"context"
	"github.com/spf13/cast"
	"go-zero_less/postcenter/cmd/rpc/pb"

	"go-zero_less/postcenter/cmd/api/internal/svc"
	"go-zero_less/postcenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (resp *types.CreateCommentResp, err error) {
	userId := l.ctx.Value("userId")
	userIdInt := cast.ToUint64(userId)
	var comments = &pb.CommentsInfoCreateRequest{
		CommentsInfo: &pb.CommentsInfoBase{
			PostId:  req.PostId,
			UserId:  userIdInt,
			Content: req.Content,
		},
	}
	commentId, err := l.svcCtx.PostRpcClient.CreateCommentsInfo(l.ctx, comments)
	if err != nil {
		return nil, err
	}

	return &types.CreateCommentResp{
		Id: commentId.Id,
	}, nil
}
