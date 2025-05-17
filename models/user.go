package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PersonID  uint      `gorm:"not null" json:"person_id"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Status    string    `gorm:"type:char(1);default:'A'" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Person    *Person   `gorm:"foreignKey:PersonID" json:"person,omitempty"`
}

func (User) TableName() string {
	return "users"
}
