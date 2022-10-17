package user

// Модель пользователя
type User struct {
	ID        string `json:"-" bson:"_id,omitempty"`
	Firstname string `json:"firstname" bson:"firstname" binding:"required"`
	Lastname  string `json:"lastname" bson:"lastname" binding:"required"`
	Username  string `json:"username" bson:"username" binding:"required"`
	Password  string `json:"password" bson:"password" binding:"required"`
}

// Модель пользователя при аутентификации
type UserSignUpDTO struct {
	ID       string `json:"-" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}
