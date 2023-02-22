package http

import (
	"fmt"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	usecase auth.UseCase
	l       logger.Interface
}

func NewAuthMiddleware(usecase auth.UseCase, l logger.Interface) gin.HandlerFunc {
	return (&AuthMiddleware{
		usecase: usecase,
		l:       l,
	}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		fmt.Println("Status Unauthorized.")
		m.l.Info("Status Unauthorized")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		fmt.Println("Status Unauthorized.")
		m.l.Info("Status Unauthorized.")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		fmt.Println("Not bearer.Status Unauthorized.")
		m.l.Info("Not bearer.Status Unauthorized.")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.usecase.ParseToken(c.Request.Context(), m.l, headerParts[1])
	if err != nil {
		status := http.StatusUnauthorized
		if err == auth.ErrInvalidAccessToken {
			m.l.Info("Invalid Access Token.")
			status = http.StatusUnauthorized
		}
		m.l.Info("Status Internal Server Error")
		c.AbortWithStatus(status)
		return
	}
	c.Set(auth.CtxUserKey, user)
}
