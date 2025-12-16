package lib

import (
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password vo.Password) (vo.Password, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return vo.NewHashedPassword(string(hashedBytes)), nil
}
