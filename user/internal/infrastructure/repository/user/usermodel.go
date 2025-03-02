package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
	"user/internal/domain/entity"
	"user/internal/domain/repository"
)

var userCASRowsWithPlaceHolder = strings.Replace(
	strings.Join(
		stringx.Remove(
			userFieldNames,
			"`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`",
		),
		"=?,")+"=?",
	"version = ?", "version = version + 1", 1)

var _ repository.UserRepository = (*customUserModel)(nil)

type (
	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) repository.UserRepository {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

// POToEntity 将 PO 转换为领域实体
func POToEntity(po *User) *entity.User {
	e := &entity.User{}
	copier.Copy(e, po)

	// ...
	return e
}

// EntityToPO 反向转换
func EntityToPO(e *entity.User) *User {
	po := &User{}
	copier.Copy(po, e)

	// ...
	return po
}

func (m *customUserModel) WithSession(session sqlx.Session) repository.UserRepository {
	return NewUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUserModel) Insert(ctx context.Context, e *entity.User) (int64, error) {
	res, err := m.defaultUserModel.Insert(ctx, EntityToPO(e))
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (m *customUserModel) FindOne(ctx context.Context, id string) (*entity.User, error) {
	po, err := m.defaultUserModel.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return POToEntity(po), nil
}

func (m *customUserModel) FindOneByMobile(ctx context.Context, mobile string) (*entity.User, error) {
	po, err := m.defaultUserModel.FindOneByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}
	return POToEntity(po), nil
}

func (m *customUserModel) FindOneByName(ctx context.Context, name sql.NullString) (*entity.User, error) {
	po, err := m.defaultUserModel.FindOneByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return POToEntity(po), nil
}

func (m *customUserModel) Update(ctx context.Context, newData *entity.User) error {
	po := EntityToPO(newData)
	query := fmt.Sprintf("update %s set %s where `id` = ? and `version` = ? ", m.table, userCASRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, po.Name, po.Password, po.Mobile, po.Gender, po.Nickname, po.Type, po.Id, po.Version)
	return err
}

func (m *customUserModel) Delete(ctx context.Context, id string) error {
	return m.defaultUserModel.Delete(ctx, id)
}
