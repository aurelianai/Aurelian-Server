package persistence

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName *string
	LastName  *string
	Email     *string
	Password  *string
	Chats     []Chat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Chat struct {
	gorm.Model
	Title    *string
	UserID   uint
	Messages []Message `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Message struct {
	gorm.Model
	role   *string
	conent *string
	ChatID uint
}
