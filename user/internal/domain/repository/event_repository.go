package repository

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/internal/domain/entity"
)

// OutboxEventRepository 定义了事件持久化所需的方法
type OutboxEventRepository interface {
	WithSession(session sqlx.Session) OutboxEventRepository
	Insert(ctx context.Context, data *entity.OutboxEvent) (int64, error)
}
