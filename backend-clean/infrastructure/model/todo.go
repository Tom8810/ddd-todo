package model

import "time"

type Todo struct {
	ID          string     `gorm:"column:id;type:char(36);primaryKey;not null"`
	Title       string     `gorm:"column:title;type:varchar(64);not null"`
	Description *string    `gorm:"column:description;type:text"`
	Status      string     `gorm:"column:status;type:varchar(10);not null"`
	Deadline    *time.Time `gorm:"column:deadline;type:datetime"`
	Priority    *string    `gorm:"column:priority;type:varchar(10)"`
	UserId      string     `gorm:"column:user_id;type:char(36);not null"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:datetime;not null"`
}
