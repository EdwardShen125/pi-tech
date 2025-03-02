package service

import (
	"errors"
	"strings"
	"user/internal/domain/entity"
)

// UserDomainService 定义领域层中关于用户的业务逻辑接口
type UserDomainService interface {
	// HashPassword 对密码进行哈希（实际可调用第三方库）
	HashPassword(password string) (string, error)
	// VerifyPassword 验证密码（示例中简单对比）
	VerifyPassword(hashed, plain string) bool
	// ChangeUserPassword 修改用户密码，确保当前密码正确、新密码符合要求
	ChangeUserPassword(user *entity.User, currentPassword, newPassword string) error
	// UpdateEmail 更新用户邮箱，校验格式等业务规则
	UpdateEmail(user *entity.User, newEmail string) error
}

type userDomainServiceImpl struct{}

// NewUserDomainService 创建 UserDomainService 实现
func NewUserDomainService() UserDomainService {
	return &userDomainServiceImpl{}
}

func (s *userDomainServiceImpl) HashPassword(password string) (string, error) {
	// 此处仅做示例，实际应调用 bcrypt 等库
	if password == "" {
		return "", nil
	}
	// 模拟哈希结果：原密码前加"hashed-"
	return "hashed-" + password, nil
}

func (s *userDomainServiceImpl) VerifyPassword(hashed, plain string) bool {
	expected, _ := s.HashPassword(plain)
	return hashed == expected
}

// ChangeUserPassword 实现修改密码的业务规则：
// 1. 检查当前密码是否正确（此处简化为直接比较，实际项目应使用安全哈希校验）；
// 2. 新密码长度必须满足要求。
func (s *userDomainServiceImpl) ChangeUserPassword(user *entity.User, currentPassword, newPassword string) error {
	// 假设 user.Password 已存储为 "hashed-" + 密码（实际应使用 bcrypt 等哈希算法）
	if user.Password != "hashed-"+currentPassword {
		return errors.New("当前密码不正确")
	}
	if len(newPassword) < 6 {
		return errors.New("新密码长度至少6位")
	}
	// 模拟密码哈希（实际请调用加密库）
	user.Password = "hashed-" + newPassword
	return nil
}

// UpdateEmail 实现更新邮箱的业务规则：
// 校验邮箱不为空且去除前后空格（实际可扩展邮箱格式校验、禁止域校验等）
func (s *userDomainServiceImpl) UpdateEmail(user *entity.User, newEmail string) error {
	newEmail = strings.TrimSpace(newEmail)
	if newEmail == "" {
		return errors.New("邮箱不能为空")
	}
	// 此处可以添加更多业务规则，如邮箱格式校验等
	user.Email = newEmail
	return nil
}
