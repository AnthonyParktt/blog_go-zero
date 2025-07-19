package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"go-zero_less/usercenter/model"
	"strings"
)

var _ PostsModel = (*customPostsModel)(nil)

type (
	// PostsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostsModel.
	PostsModel interface {
		postsModel
		withSession(session sqlx.Session) PostsModel
		DeleteSoft(ctx context.Context, id uint64) (int64, error)
		FindPosts(ctx context.Context, param PostQueryParam) ([]*Posts, error)
	}

	customPostsModel struct {
		*defaultPostsModel
	}

	PostQueryParam struct {
		Key       string
		UserId    uint64
		Page      uint64
		Size      uint64
		DelFlag   bool
		StartTime uint64
		EndTime   uint64
		Status    string
	}
)

// NewPostsModel returns a model for the database table.
func NewPostsModel(conn sqlx.SqlConn) PostsModel {
	return &customPostsModel{
		defaultPostsModel: newPostsModel(conn),
	}
}

func (m *customPostsModel) withSession(session sqlx.Session) PostsModel {
	return NewPostsModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customPostsModel) DeleteSoft(ctx context.Context, id uint64) (int64, error) {
	query := fmt.Sprintf("update %s set deleted_at = now() where id = ?", m.table)
	execCtx, err := m.conn.ExecCtx(ctx, query, id)
	if err != nil {
		return 0, err
	}

	rows, err := execCtx.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (m *customPostsModel) FindPosts(ctx context.Context, params PostQueryParam) ([]*Posts, error) {
	var conditions []string
	var args []interface{}

	if params.UserId != 0 {
		conditions = append(conditions, "user_id = ?")
		args = append(args, params.UserId)
	}
	if params.StartTime != 0 {
		conditions = append(conditions, "created_at >= ?")
		args = append(args, params.StartTime)
	}
	if params.EndTime != 0 {
		conditions = append(conditions, "created_at <= ?")
		args = append(args, params.EndTime)
	}
	if params.Status != "" {
		conditions = append(conditions, "status = ?")
		args = append(args, params.Status)
	}
	if len(conditions) == 0 {
		return nil, ErrMissConditions
	}

	columns := strings.Join(stringx.Remove(postsFieldNames, "content"), ",")
	query := fmt.Sprintf("select %s from %s where deleted_at is null %s", columns, m.table, strings.Join(conditions, " and "))
	posts := []*Posts{}
	err := m.conn.QueryRowsCtx(ctx, posts, query, args...)
	switch {
	case errors.Is(err, model.ErrNotFound):
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return posts, nil
	}
}
