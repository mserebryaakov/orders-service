package user

import "context"

// Интерфейс взаимодействия с репозиторием пользователей
type UserRepository interface {
	// Создание пользовтеля
	CreateUser(ctx context.Context, user User) (string, error)
	// Получение пользователя
	GetUser(ctx context.Context, user UserSignUpDTO) (u UserSignUpDTO, err error)
}
