package model

import "time"

type User struct {
	ID        string    `gorm:"column:id;type:char(36);primaryKey;not null"`
	Name      string    `gorm:"column:name;type:varchar(64);not null"`
	Email     string    `gorm:"column:email;type:varchar(256);unique;not null"`
	Password  string    `gorm:"column:password;type:varchar(256);not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null"`
}
