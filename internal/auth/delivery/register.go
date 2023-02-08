package delivery

import (
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	router.POST("/api/user/register", h.SignUp)
	router.POST("/api/user/login", h.SignIn)
}
