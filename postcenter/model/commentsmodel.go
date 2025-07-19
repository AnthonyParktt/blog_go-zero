package model

import (
	"context"
	"errors"
	"fmt"
	"go-zero_less/usercenter/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentsModel = (*customCommentsModel)(nil)

type (
	// CommentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentsModel.
	CommentsModel interface {
		commentsModel
		withSession(session sqlx.Session) CommentsModel
		FindAll(ctx context.Context, postId uint64) ([]Comments, error)
	}

	customCommentsModel struct {
		*defaultCommentsModel
	}
)

// NewCommentsModel returns a model for the database table.
func NewCommentsModel(conn sqlx.SqlConn) CommentsModel {
	return &customCommentsModel{
		defaultCommentsModel: newCommentsModel(conn),
	}
}

func (m *customCommentsModel) withSession(session sqlx.Session) CommentsModel {
	return NewCommentsModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customCommentsModel) FindAll(ctx context.Context, postId uint64) ([]Comments, error) {
	query := fmt.Sprintf("select %s from %s where post_id = ? and deleted_at is null", commentsRows, m.table)
	var resp []Comments
	err := m.conn.QueryRowsCtx(ctx, &resp, query, postId)
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, model.ErrNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
