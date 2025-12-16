package service

type AuthService interface {
	GenerateToken(userID, email string) (string, error)
}
