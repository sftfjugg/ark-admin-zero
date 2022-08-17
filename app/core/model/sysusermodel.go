package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type SysUserDetail struct {
	Id         int64     `db:"id"`          // 编号
	Account    string    `db:"account"`     // 账号
	Username   string    `db:"username"`    // 姓名
	Nickname   string    `db:"nickname"`    // 昵称
	Avatar     string    `db:"avatar"`      // 头像
	Gender     int64     `db:"gender"`      // 0=保密 1=女 2=男
	Profession string    `db:"profession"`  // 职称
	Job        string    `db:"job"`         // 岗位
	Dept       string    `db:"dept"`        // 部门
	Roles      string    `db:"roles"`       // 角色集
	Birthday   time.Time `db:"birthday"`    // 生日
	Email      string    `db:"email"`       // 邮件
	Mobile     string    `db:"mobile"`      // 手机号
	Remark     string    `db:"remark"`      // 备注
	OrderNum   int64     `db:"order_num"`   // 排序值
	Status     int64     `db:"status"`      // 0=禁用 1=开启
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
}

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		FindByCondition(ctx context.Context, condition string, value int64) ([]*SysUser, error)
		FindByPage(ctx context.Context) ([]*SysUserDetail, error)
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn, c cache.CacheConf) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn, c),
	}
}

func (m *customSysUserModel) FindByCondition(ctx context.Context, condition string, value int64) ([]*SysUser, error) {
	query := fmt.Sprintf("select %s from %s where %s=?", sysUserRows, m.table, condition)
	var resp []*SysUser
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, value)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysUserModel) FindByPage(ctx context.Context) ([]*SysUserDetail, error) {
	query := fmt.Sprintf("SELECT u.id,u.account,u.username,u.nickname,u.avatar,u.gender,p.name as profession,j.name as job,d.name as dept,GROUP_CONCAT(r.name) as roles,u.birthday,u.email,u.mobile,u.remark,u.order_num,u.status,u.create_time,u.update_time FROM `sys_user` u LEFT JOIN sys_profession p ON u.profession_id=p.id LEFT JOIN sys_dept d ON u.dept_id=d.id LEFT JOIN sys_job j ON u.job_id=j.id LEFT JOIN sys_role r ON JSON_CONTAINS(u.role_ids,JSON_ARRAY(r.id)) GROUP BY u.id")
	var resp []*SysUserDetail
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}