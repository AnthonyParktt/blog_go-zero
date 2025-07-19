package postinfologic

import (
	"context"
	"errors"
	"go-zero_less/postcenter/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-zero_less/postcenter/cmd/rpc/internal/svc"
	"go-zero_less/postcenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostListPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostListPageLogic {
	return &GetPostListPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostListPageLogic) GetPostListPage(in *pb.PostSearchRequest) (*pb.PostInfoListPage, error) {
	queryParam := model.PostQueryParam{
		Key:       in.Keyword,
		UserId:    in.UserId,
		Page:      in.Page,
		Size:      in.PageSize,
		DelFlag:   false,
		StartTime: in.StartDate,
		EndTime:   in.EndDate,
		Status:    "",
	}
	posts, err := l.svcCtx.PostModel.FindPosts(l.ctx, queryParam)
	switch {
	case errors.Is(err, model.ErrNotFound):
		return nil, nil
	case err != nil:
		return nil, status.Errorf(codes.Internal, "内部错误%v", err)
	default:
		var postInfoList []*pb.PostInfoBase
		for _, post := range posts {
			postInfoList = append(postInfoList, &pb.PostInfoBase{
				Id:        post.Id,
				Title:     post.Title,
				Content:   post.Content,
				UserId:    post.UserId,
				UpdatedAt: uint64(post.UpdatedAt.Time.Unix()),
				CreatedAt: uint64(post.CreatedAt.Time.Unix())})
		}
		return &pb.PostInfoListPage{
			PostInfoList: postInfoList,
		}, nil
	}
}
