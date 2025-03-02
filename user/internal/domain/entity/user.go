package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// User 表示领域层中的用户实体
type User struct {
	ID           string // 聚合根ID
	FirstName    string
	LastName     string
	Email        string
	Password     string // 注意：实际保存为哈希
	Balance      int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	domainEvents []DomainEvent // 存放领域事件的集合
}

// NewUser 是一个工厂函数，用于创建一个新的 User 实体
func NewUser(firstName, lastName, email, password string) (*User, error) {
	// 简单校验
	if strings.TrimSpace(firstName) == "" || strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return nil, errors.New("first name, email, and password are required")
	}

	user := &User{
		ID:        uuid.New().String(),
		FirstName: strings.TrimSpace(firstName),
		LastName:  strings.TrimSpace(lastName),
		Email:     strings.TrimSpace(email),
		Password:  password, // 此处实际应进行密码哈希处理
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// 抛出领域事件：用户注册成功
	user.AddDomainEvent(&UserRegisteredEvent{
		User:       user,
		occurredAt: time.Now(),
	})
	return user, nil
}

// PublicUser 返回供外部展示的用户信息（不包含敏感数据）
func (u *User) PublicUser() *User {
	return &User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// UpdateEmail 更新用户邮箱，并执行必要的业务校验
func (u *User) UpdateEmail(newEmail string) error {
	newEmail = strings.TrimSpace(newEmail)
	if newEmail == "" {
		return errors.New("email cannot be empty")
	}
	// 此处可以增加更复杂的邮箱格式校验
	u.Email = newEmail
	u.UpdatedAt = time.Now()
	return nil
}

// UpdateBalance 更新账户余额，可以是加值或设置新余额，具体业务逻辑视情况而定
func (u *User) UpdateBalance(newBalance int64) error {
	// 例如：不允许负余额
	if newBalance < 0 {
		return errors.New("balance cannot be negative")
	}
	u.Balance = newBalance
	u.UpdatedAt = time.Now()
	return nil
}

// AddDomainEvent 添加一个领域事件到实体内
func (u *User) AddDomainEvent(event DomainEvent) {
	u.domainEvents = append(u.domainEvents, event)
}

// DomainEvents 返回实体中累积的所有领域事件
func (u *User) DomainEvents() []DomainEvent {
	return u.domainEvents
}
