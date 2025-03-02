package entity

import (
	jsoniter "github.com/json-iterator/go"
	"time"
)

// DomainEvent 定义了所有领域事件应实现的接口
type DomainEvent interface {
	EventName() string
	OccurredAt() time.Time
	AggregateID() string
}

// ToOutboxEvent 转换成收件箱模式下领域事件的实体
func ToOutboxEvent(e DomainEvent) *OutboxEvent {
	bytes, err := jsoniter.Marshal(e)
	if err != nil {
		return nil
	}

	return &OutboxEvent{
		AggregateID: e.AggregateID(),
		EventType:   e.EventName(),
		Payload:     string(bytes),
		OccurredAt:  e.OccurredAt(),
	}
}

// OutboxEvent 是收件箱模式的领域事件
type OutboxEvent struct {
	AggregateID string
	EventType   string
	Payload     string // 序列化后的事件数据
	OccurredAt  time.Time
	Published   bool // 是否已发布
}

// UserRegisteredEvent 是用户注册成功时产生的领域事件
type UserRegisteredEvent struct {
	User       *User     // 事件中携带了注册的用户
	occurredAt time.Time // 事件发生时间
}

// EventName 返回事件名称
func (e *UserRegisteredEvent) EventName() string {
	return "UserRegisteredEvent"
}

// OccurredAt 返回事件发生时间
func (e *UserRegisteredEvent) OccurredAt() time.Time {
	return e.occurredAt
}

// AggregateID 聚合id
func (e *UserRegisteredEvent) AggregateID() string {
	return e.User.ID
}
