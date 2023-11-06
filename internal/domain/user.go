package domain

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time    `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time    `gorm:"default:now()" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"deleted_at"`

	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Phone    string `gorm:"unique;not null"`
}

type UserRepository interface {
	Save(User) error
	FindByID(string) (*User, error)
	Update(User) error
	Delete(string) error
}
