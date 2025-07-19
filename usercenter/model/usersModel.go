package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		withSession(session sqlx.Session) UsersModel
		FindUserByName(ctx context.Context, username string) (*Users, error)
		FindOneUser(ctx context.Context, id uint64, username string) (*Users, error)
		FindUsers(ctx context.Context, params UserQueryParams) ([]*Users, error)
		AddUserPostNum(ctx context.Context, userId uint64, postNumIncr int64) error
	}

	customUsersModel struct {
		*defaultUsersModel
	}
	UserQueryParams struct {
		Id        uint64
		Username  string
		Email     string
		Status    string
		StartTime int64
		EndTime   int64
		PageSize  int32
		PageNum   int32
		OrderBy   string
		OrderType string
		Keyword   string // 支持模糊查询的关键字
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}

func (m *customUsersModel) withSession(session sqlx.Session) UsersModel {
	return NewUsersModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUsersModel) FindUserByName(ctx context.Context, username string) (*Users, error) {
	query := fmt.Sprintf("select %s from %s where username = ? limit 1", usersRows, m.table)
	var resp Users
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUsersModel) FindOneUser(ctx context.Context, id uint64, username string) (*Users, error) {
	var conditions []string
	var args []interface{}

	if id != 0 {
		conditions = append(conditions, "id = ?")
		args = append(args, id)
	}
	if username != "" {
		conditions = append(conditions, "username = ?")
		args = append(args, username)
	}
	if len(conditions) == 0 {
		return nil, ErrMissConditions
	}

	query := fmt.Sprintf("select %s from %s where %s limit 1", usersRows, m.table, strings.Join(conditions, " and "))
	var resp Users
	err := m.conn.QueryRowCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUsersModel) FindUsers(ctx context.Context, params UserQueryParams) ([]*Users, error) {
	var conditions []string
	var args []interface{}

	//添加一个固定条件，deleted_at is null
	conditions = append(conditions, "deleted_at is null")

	if params.Id != 0 {
		conditions = append(conditions, "id = ?")
		args = append(args, params.Id)
	}
	if params.Username != "" {
		conditions = append(conditions, "username = ?")
		args = append(args, params.Username)
	}
	if params.Email != "" {
		conditions = append(conditions, "email = ?")
		args = append(args, params.Email)
	}
	if params.Status != "" {
		conditions = append(conditions, "status = ?")
		args = append(args, params.Status)
	}
	if params.StartTime != 0 {
		conditions = append(conditions, "created_at >= ?")
		args = append(args, params.StartTime)
	}
	if params.EndTime != 0 {
		conditions = append(conditions, "created_at <= ?")
		args = append(args, params.EndTime)
	}

	//var resp []Users
	query := fmt.Sprintf("select %s from %s where %s", usersRows, m.table, strings.Join(conditions, " and "))
	var resp []*Users
	err := m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUsersModel) AddUserPostNum(ctx context.Context, userId uint64, postNumIncr int64) error {
	query := fmt.Sprintf("update %s set post_num = post_num + ? where id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, postNumIncr, userId)
	return err
}
