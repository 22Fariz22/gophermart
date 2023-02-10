package delivery

import (
	"fmt"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase, l logger.Interface) {
	h := NewHandler(uc, l)

	authEndpoints := router.Group("/user")
	{
		authEndpoints.POST("/register/", h.SignUp)
		authEndpoints.POST("/login", h.SignIn)
	}
	fmt.Println("auth-register-RegisterHTTPEndpoints: ")
}
