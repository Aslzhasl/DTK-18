package auth

// UserInfo — информация о пользователе, полученная после верификации
type UserInfo struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
	Email   string `json:"email"`
	Role    string `json:"role"`
}

// AuthClient — интерфейс внешнего сервиса аутентификации
type AuthClient interface {
	VerifyUser(token string) (UserInfo, error)
}
