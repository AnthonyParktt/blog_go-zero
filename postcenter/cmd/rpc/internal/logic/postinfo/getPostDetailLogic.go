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

type GetPostDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailLogic {
	return &GetPostDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostDetailLogic) GetPostDetail(in *pb.PostInfoId) (*pb.PostInfoDetailWithComments, error) {
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, in.Id)
	switch {
	case errors.Is(err, model.ErrNotFound):
		return nil, status.Errorf(codes.NotFound, "文章不存在%v", err)
	case err != nil:
		return nil, status.Errorf(codes.Internal, "内部错误：%v", err)
	default:
		comments, errComment := l.svcCtx.CommentsModel.FindAll(l.ctx, in.Id)
		if errComment != nil {
			return nil, status.Errorf(codes.Internal, "内部错误：%v", errComment)
		}
		postInfoDetail := &pb.PostInfoDetailWithComments{
			PostInfo: &pb.PostInfoBase{
				Id:      post.Id,
				Title:   post.Title,
				Content: post.Content,
				UserId:  post.UserId,
			},
		}
		if comments != nil {
			var commentsPb []*pb.CommentsInfoBase
			commentsPb = make([]*pb.CommentsInfoBase, len(comments))
			// 转换为pb格式
			for index, comment := range comments {
				commentsPb[index] = &pb.CommentsInfoBase{
					Id:      comment.Id,
					Content: comment.Content,
					UserId:  comment.UserId,
					PostId:  comment.PostId,
				}
			}
			postInfoDetail.CommentsInfoList = append(postInfoDetail.CommentsInfoList, &pb.CommentsInfoListPage{
				CommentsInfoList: commentsPb,
			})
		}
		return postInfoDetail, nil
	}
}
