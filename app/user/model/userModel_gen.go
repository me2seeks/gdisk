// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`update_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheCloudDiskUserIdPrefix    = "cache:cloudDisk:user:id:"
	cacheCloudDiskUserEmailPrefix = "cache:cloudDisk:user:email:"
)

type (
	userModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByEmail(ctx context.Context, email string) (*User, error)
		FindOneByIdentity(ctx context.Context, Identity string) (*User, error)
		Update(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *User) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id       int64     `db:"id"`
		Identity string    `db:"identity"`
		Email    string    `db:"email"`
		Password string    `db:"password"`
		Name     string    `db:"name"`
		Sex      int64     `db:"sex"` // 性别 0:男 1:女
		Avatar   string    `db:"avatar"`
		Info     string    `db:"info"`
		DeleteAt time.Time `db:"delete_at"`
		DelState int64     `db:"del_state"`
		CreateAt time.Time `db:"create_at"`
		UpdateAt time.Time `db:"update_at"`
		Version  int64     `db:"version"` // 版本号
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	data.DeleteAt = time.Unix(0, 0)
	cloudDiskUserEmailKey := fmt.Sprintf("%s%v", cacheCloudDiskUserEmailPrefix, data.Email)
	cloudDiskUserIdKey := fmt.Sprintf("%s%v", cacheCloudDiskUserIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Identity, data.Email, data.Password, data.Name, data.Sex, data.Avatar, data.Info, data.DeleteAt, data.DelState, data.Version)
		}
		return conn.ExecCtx(ctx, query, data.Identity, data.Email, data.Password, data.Name, data.Sex, data.Avatar, data.Info, data.DeleteAt, data.DelState, data.Version)
	}, cloudDiskUserEmailKey, cloudDiskUserIdKey)
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	cloudDiskUserIdKey := fmt.Sprintf("%s%v", cacheCloudDiskUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, cloudDiskUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	cloudDiskUserEmailKey := fmt.Sprintf("%s%v", cacheCloudDiskUserEmailPrefix, email)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, cloudDiskUserEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? and del_state = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByIdentity(ctx context.Context, Identity string) (*User, error) {
	cloudDiskUserEmailKey := fmt.Sprintf("%s%v", cacheCloudDiskUserEmailPrefix, Identity)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, cloudDiskUserEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `identity` = ? and del_state = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, Identity, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	cloudDiskUserEmailKey := fmt.Sprintf("%s%v", cacheCloudDiskUserEmailPrefix, data.Email)
	cloudDiskUserIdKey := fmt.Sprintf("%s%v", cacheCloudDiskUserIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Identity, data.Email, data.Password, data.Name, data.Sex, data.Avatar, data.Info, data.DeleteAt, data.DelState, data.Version, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.Identity, data.Email, data.Password, data.Name, data.Sex, data.Avatar, data.Info, data.DeleteAt, data.DelState, data.Version, data.Id)
	}, cloudDiskUserEmailKey, cloudDiskUserIdKey)
}

func (m *defaultUserModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *User) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	cloudDiskUserEmailKey := fmt.Sprintf("%s%v", cacheCloudDiskUserEmailPrefix, data.Email)
	cloudDiskUserIdKey := fmt.Sprintf("%s%v", cacheCloudDiskUserIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, userRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Identity, data.Email, data.Password, data.Name, data.Sex, data.Avatar, data.Info, data.DeleteAt, data.DelState, data.Version, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.Identity, data.Email, data.Password, data.Name, data.Sex, data.Avatar, data.Info, data.DeleteAt, data.DelState, data.Version, data.Id, oldVersion)
	}, cloudDiskUserEmailKey, cloudDiskUserIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return nil
}

func (m *defaultUserModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	cloudDiskUserEmailKey := fmt.Sprintf("%s%v", cacheCloudDiskUserEmailPrefix, data.Email)
	cloudDiskUserIdKey := fmt.Sprintf("%s%v", cacheCloudDiskUserIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, cloudDiskUserEmailKey, cloudDiskUserIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCloudDiskUserIdPrefix, primary)
}
func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
