package v1

import "github.com/gin-gonic/gin"

// Response при order create
type IdResponse struct {
	Id string `json:"id"`
}

// Token пользователя
type TokenResponse struct {
	Token string `json:"token"`
}

// Cтруктура ошибки
type errorResponse struct {
	Message string `json:"message"`
}

// Cтруктура статуса ответа
type statusResponse struct {
	Status string `json:"status"`
}

// Функция обработки ошибок
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	// метод блокирует выполнение следующих обработчик и записывает в ответ статус и сообщение
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
