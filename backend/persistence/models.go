package persistence

import (
	"time"
)

type User struct {
	ID        uint64 `gorm:"primarykey;type:BIGINT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	Chats     []Chat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Chat struct {
	ID        uint64 `gorm:"primarykey;type:BIGINT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string    `gorm:"not null"`
	UserID    uint64    `gorm:"not null"`
	Messages  []Message `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Message struct {
	ID        uint64 `gorm:"primarykey;type:BIGINT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      string `gorm:"not null"`
	Content   string `gorm:"not null"`
	ChatID    uint64
}
