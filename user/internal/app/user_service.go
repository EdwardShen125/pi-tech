package app

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/internal/domain/entity"
	"user/internal/domain/repository"
	"user/internal/domain/service"
)

type UserService struct {
	conn            sqlx.SqlConn
	userRepo        repository.UserRepository
	userDomainSvc   service.UserDomainService
	eventRepository repository.OutboxEventRepository
}

// NewUserService 构造函数，注入依赖
func NewUserService(
	conn sqlx.SqlConn,
	repo repository.UserRepository,
	domainSvc service.UserDomainService,
	eventRepository repository.OutboxEventRepository,
) *UserService {
	return &UserService{
		conn:            conn,
		userRepo:        repo,
		userDomainSvc:   domainSvc,
		eventRepository: eventRepository,
	}
}

// RegisterUser 处理用户注册
func (s *UserService) RegisterUser(ctx context.Context, firstName, lastName, email, password string) (*entity.User, error) {
	// 创建用户实体
	user, err := entity.NewUser(firstName, lastName, email, password)
	if err != nil {
		return nil, err
	}
	// 对密码进行哈希处理
	hashedPwd, err := s.userDomainSvc.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPwd

	// Transaction Outbox 模式
	err = s.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		// 持久化用户
		if _, err := s.userRepo.WithSession(session).Insert(ctx, user); err != nil {
			return err
		}

		for _, event := range user.DomainEvents() {
			_, err := s.eventRepository.WithSession(session).Insert(ctx, entity.ToOutboxEvent(event))
			if err != nil {
				return err
			}
		}

		return nil
	})

	return user, err
}

func (s *UserService) UpdateUserBalance(ctx context.Context, userID string, newBalance int64) error {
	// 加载用户数据
	user, err := s.userRepo.FindOne(ctx, userID)
	if err != nil {
		return err
	}
	// 修改余额，更新版本
	if err = user.UpdateBalance(newBalance); err != nil {
		return err
	}
	// 更新数据库时会检查版本号
	return s.userRepo.Update(ctx, user)
}

// FindOneByMobile 根据手机号查询用户
func (s *UserService) FindOneByMobile(ctx context.Context, mobile string) (*entity.User, error) {
	return s.userRepo.FindOneByMobile(ctx, mobile)
}
