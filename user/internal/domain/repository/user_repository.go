package repository

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/internal/domain/entity"
)

// UserRepository 定义了用户持久化所需的方法
type UserRepository interface {
	WithSession(session sqlx.Session) UserRepository
	Insert(ctx context.Context, data *entity.User) (int64, error)
	FindOne(ctx context.Context, id int64) (*entity.User, error)
	FindOneByMobile(ctx context.Context, mobile string) (*entity.User, error)
	FindOneByName(ctx context.Context, name sql.NullString) (*entity.User, error)
	Update(ctx context.Context, data *entity.User) error
	Delete(ctx context.Context, id int64) error
}
