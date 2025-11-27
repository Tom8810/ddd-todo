package dto

import "time"

type UserOutput struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUpdateInput struct {
	UserID      string
	UpdateField string
}
