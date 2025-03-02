package event

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/internal/domain/entity"
	"user/internal/domain/repository"
)

var _ repository.OutboxEventRepository = (*customEventModel)(nil)

type (
	customEventModel struct {
		*defaultEventModel
	}
)

// NewEventModel returns a model for the database table.
func NewEventModel(conn sqlx.SqlConn) repository.OutboxEventRepository {
	return &customEventModel{
		defaultEventModel: newEventModel(conn),
	}
}

// POToEntity 将 PO 转换为领域实体
func POToEntity(po *Event) *entity.OutboxEvent {
	e := &entity.OutboxEvent{}
	copier.Copy(e, po)

	// ...
	return e
}

// EntityToPO 反向转换
func EntityToPO(e *entity.OutboxEvent) *Event {
	po := &Event{}
	copier.Copy(po, e)

	// ...
	return po
}

func (m *customEventModel) WithSession(session sqlx.Session) repository.OutboxEventRepository {
	return NewEventModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customEventModel) Insert(ctx context.Context, e *entity.OutboxEvent) (int64, error) {
	res, err := m.defaultEventModel.Insert(ctx, EntityToPO(e))
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
