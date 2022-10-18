package v1

import (
	"context"
	"net/http"
	"orders-service/internal/domain/user"
	service "orders-service/internal/services/user"
	"orders-service/pkg/logger"
	"strings"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	log         *logger.Logger
	userUseCase *service.UserUseCase
}

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func NewUserHandler(log *logger.Logger, userUseCase *service.UserUseCase) *userHandler {
	return &userHandler{
		log:         log,
		userUseCase: userUseCase,
	}
}

// Регистрация эндпоинтов для работы с заказами
func (h *userHandler) Register(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
}

// Middleware на авторизацию
func (h *userHandler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.userUseCase.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

// Регистрация пользователя
// @Summary SignUp
// @Tags auth
// @Description create account
// @Accept  json
// @Produce  json
// @Param input body user.User true "account info"
// @Success 200 {string} IdResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *userHandler) signUp(c *gin.Context) {
	var user user.User

	// Валидация body
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()

	// Создание пользователя
	id, err := h.userUseCase.CreateUser(ctx, user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Успешный ответ
	c.JSON(http.StatusOK, &IdResponse{
		Id: id,
	})
}

// Аутентификация пользователя
// @Summary SignIn
// @Tags auth
// @Description authentification account
// @Accept  json
// @Produce  json
// @Param input body user.UserSignUpDTO true "account info"
// @Success 200 {string} TokenResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *userHandler) signIn(c *gin.Context) {
	var user user.UserSignUpDTO

	// Валидация body
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()

	// Генерация токена пользователя
	token, err := h.userUseCase.GenerateToken(ctx, user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Успешный ответ
	c.JSON(http.StatusOK, &TokenResponse{
		Token: token,
	})
}
