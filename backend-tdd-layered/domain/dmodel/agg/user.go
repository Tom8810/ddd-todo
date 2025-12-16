package agg

import (
	"time"

	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

type User struct {
	ID        vo.UserID
	Name      vo.Name
	Email     vo.Email
	Password  vo.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(
	userID vo.UserID,
	name vo.Name,
	email vo.Email,
	password vo.Password,
) *User {
	now := time.Now()
	return &User{
		ID:        userID,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func ReconstructUser(
	userID vo.UserID,
	name vo.Name,
	email vo.Email,
	password vo.Password,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	return &User{
		ID:        userID,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func GetID(u *User) string {
	return u.ID.Value()
}

func GetName(u *User) string {
	return u.Name.Value()
}

func GetEmail(u *User) string {
	return u.Email.Value()
}

func GetPassword(u *User) string {
	return u.Password.Value()
}

func GetCreatedAt(u *User) time.Time {
	return u.CreatedAt
}

func GetUpdatedAt(u *User) time.Time {
	return u.UpdatedAt
}

func (u *User) UpdateName(name vo.Name) {
	u.Name = name
	u.updateTimestamp()
}

func (u *User) UpdateEmail(email vo.Email) {
	u.Email = email
	u.updateTimestamp()
}

func (u *User) UpdatePassword(password vo.Password) {
	u.Password = password
	u.updateTimestamp()
}

func (u *User) updateTimestamp() {
	u.UpdatedAt = time.Now()
}

// // VerifyPassword checks if the provided password matches the user's password
// func (u *User) VerifyPassword(plainPassword string) bool {
// 	return u.Password.Verify(plainPassword)
// }
