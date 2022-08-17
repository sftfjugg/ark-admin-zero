// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysJobFieldNames          = builder.RawFieldNames(&SysJob{})
	sysJobRows                = strings.Join(sysJobFieldNames, ",")
	sysJobRowsExpectAutoSet   = strings.Join(stringx.Remove(sysJobFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	sysJobRowsWithPlaceHolder = strings.Join(stringx.Remove(sysJobFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheArkAdminZeroSysJobIdPrefix   = "cache:arkAdminZero:sysJob:id:"
	cacheArkAdminZeroSysJobNamePrefix = "cache:arkAdminZero:sysJob:name:"
)

type (
	sysJobModel interface {
		Insert(ctx context.Context, data *SysJob) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysJob, error)
		FindOneByName(ctx context.Context, name string) (*SysJob, error)
		Update(ctx context.Context, data *SysJob) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysJobModel struct {
		sqlc.CachedConn
		table string
	}

	SysJob struct {
		Id         int64     `db:"id"`          // 编号
		Name       string    `db:"name"`        // 岗位名称
		Status     int64     `db:"status"`      // 0=禁用 1=开启
		OrderNum   int64     `db:"order_num"`   // 排序值
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 开启时间
	}
)

func newSysJobModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSysJobModel {
	return &defaultSysJobModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_job`",
	}
}

func (m *defaultSysJobModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	arkAdminZeroSysJobIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysJobIdPrefix, id)
	arkAdminZeroSysJobNameKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysJobNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, arkAdminZeroSysJobIdKey, arkAdminZeroSysJobNameKey)
	return err
}

func (m *defaultSysJobModel) FindOne(ctx context.Context, id int64) (*SysJob, error) {
	arkAdminZeroSysJobIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysJobIdPrefix, id)
	var resp SysJob
	err := m.QueryRowCtx(ctx, &resp, arkAdminZeroSysJobIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysJobRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultSysJobModel) FindOneByName(ctx context.Context, name string) (*SysJob, error) {
	arkAdminZeroSysJobNameKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysJobNamePrefix, name)
	var resp SysJob
	err := m.QueryRowIndexCtx(ctx, &resp, arkAdminZeroSysJobNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", sysJobRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
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

func (m *defaultSysJobModel) Insert(ctx context.Context, data *SysJob) (sql.Result, error) {
	arkAdminZeroSysJobIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysJobIdPrefix, data.Id)
	arkAdminZeroSysJobNameKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysJobNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, sysJobRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Status, data.OrderNum)
	}, arkAdminZeroSysJobIdKey, arkAdminZeroSysJobNameKey)
	return ret, err
}

func (m *defaultSysJobModel) Update(ctx context.Context, newData *SysJob) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	arkAdminZeroSysJobIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysJobIdPrefix, data.Id)
	arkAdminZeroSysJobNameKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysJobNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysJobRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.Status, newData.OrderNum, newData.Id)
	}, arkAdminZeroSysJobIdKey, arkAdminZeroSysJobNameKey)
	return err
}

func (m *defaultSysJobModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheArkAdminZeroSysJobIdPrefix, primary)
}

func (m *defaultSysJobModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysJobRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysJobModel) tableName() string {
	return m.table
}
